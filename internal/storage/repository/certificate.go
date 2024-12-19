package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
)

type CertificateRepository struct {
	db *sql.DB
}

func NewCertificateRepository(db *sql.DB) *CertificateRepository {
	return &CertificateRepository{db: db}
}

// Create Certificate
func (r *CertificateRepository) Create(req *pb.CreateCertificate) (*pb.Void, error) {
	query := `INSERT INTO certificates (name, ielts_score, cefr_level, certificate_url) 
			  VALUES ($1::jsonb, $2, $3, $4)`

	nameJson, err := json.Marshal(req.Name)
	if err != nil {
		return nil, fmt.Errorf("marshal name failed: %w", err)
	}

	_, err = r.db.Exec(query, string(nameJson), req.IeltsScore, req.CefrLevel, req.CertificateUrl)
	return nil, err
}

// Update Certificate
func (r *CertificateRepository) Update(req *pb.UpdateCertificate) (*pb.Void, error) {
	query := `UPDATE certificates SET `
	var updates []string
	var args []interface{}

	if req.Name != nil {
		nameJson, err := json.Marshal(req.Name)
		if err != nil {
			return nil, fmt.Errorf("marshal name failed: %w", err)
		}
		args = append(args, string(nameJson))
		updates = append(updates, "name=$"+strconv.Itoa(len(args)))
	}
	if req.IeltsScore != 0 {
		args = append(args, req.IeltsScore)
		updates = append(updates, "ielts_score=$"+strconv.Itoa(len(args)))
	}
	if req.CefrLevel != "" && req.CefrLevel != "string" {
		args = append(args, req.CefrLevel)
		updates = append(updates, "cefr_level=$"+strconv.Itoa(len(args)))
	}
	if req.CertificateUrl != "" && req.CertificateUrl != "string" {
		args = append(args, req.CertificateUrl)
		updates = append(updates, "certificate_url=$"+strconv.Itoa(len(args)))
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(updates, ", ") + ", updated_at=now() WHERE id=$" + strconv.Itoa(len(args)+1) + " AND deleted_at=0"
	args = append(args, req.Id)

	_, err := r.db.Exec(query, args...)
	return nil, err
}

// Get Certificate By ID
func (r *CertificateRepository) GetById(req *pb.ById) (*pb.CertificateRes, error) {
	query := `SELECT id, name::jsonb, ielts_score::TEXT, cefr_level, certificate_url, 
			  to_char(created_at, 'YYYY-MM-DD HH24:MI')
			  FROM certificates WHERE id=$1 AND deleted_at=0`

	res := &pb.CertificateRes{}
	var nameJson string

	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id, &nameJson, &res.IeltsScore, &res.CefrLevel, &res.CertificateUrl, &res.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("certificate not found")
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(nameJson), &res.Name)
	if err != nil {
		return nil, fmt.Errorf("unmarshal name failed: %w", err)
	}

	return res, nil
}

// Get Certificate List
func (r *CertificateRepository) GetList(req *pb.GetListCertificateReq) (*pb.GetListCertificateRes, error) {
	query := `SELECT COUNT(*) OVER (), id, name::jsonb, ielts_score::TEXT, cefr_level, 
			  certificate_url, to_char(created_at, 'YYYY-MM-DD HH24:MI') 
			  FROM certificates WHERE deleted_at=0`

	var args []interface{}
	var filters []string

	// Filter by language if specified
	if req.Language != "" {
		filters = append(filters, "name->>$1 IS NOT NULL")
		args = append(args, req.Language)
	}

	if len(filters) > 0 {
		query += " AND " + strings.Join(filters, " AND ")
	}

	// Pagination
	query += " LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, req.Filter.Limit, req.Filter.Offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := &pb.GetListCertificateRes{}
	for rows.Next() {
		var certificate pb.CertificateRes
		var nameJson string
		err := rows.Scan(
			&res.TotalCount, &certificate.Id, &nameJson, &certificate.IeltsScore,
			&certificate.CefrLevel, &certificate.CertificateUrl, &certificate.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(nameJson), &certificate.Name)
		if err != nil {
			return nil, fmt.Errorf("unmarshal name failed: %w", err)
		}

		res.Certificates = append(res.Certificates, &certificate)
	}

	return res, nil
}

// Delete Certificate
func (r *CertificateRepository) Delete(req *pb.ById) (*pb.Void, error) {
	query := `UPDATE certificates SET deleted_at=EXTRACT(EPOCH FROM NOW()) WHERE id=$1 AND deleted_at=0`
	_, err := r.db.Exec(query, req.Id)
	return nil, err
}
