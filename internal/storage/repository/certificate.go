package repository

import (
	"database/sql"
	"fmt"
	"log"
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

func (r *CertificateRepository) Create(req *pb.CreateCertificate) (*pb.Void, error) {

	query := `INSERT INTO certificates
				(name, ielts_score, cefr_level,certificate_url)
			VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, req.Name, req.IeltsScore, req.CefrLevel, req.CertificateUrl)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *CertificateRepository) Update(req *pb.UpdateCertificate) (*pb.Void, error) {
    query := `UPDATE certificates SET`
    var args []interface{}
    var conditions []string

    // Add conditions for each field
    if req.Name != "string" && req.Name != "" {
        args = append(args, req.Name)
        conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
    }
    if req.IeltsScore != 0 {
        args = append(args, req.IeltsScore)
        conditions = append(conditions, fmt.Sprintf("ielts_score = $%d", len(args)))
    }
    if req.CertificateUrl != "string" && req.CertificateUrl != "" {
        args = append(args, req.CertificateUrl)
        conditions = append(conditions, fmt.Sprintf("certificate_url = $%d", len(args)))
    }
    if req.CefrLevel != "string" && req.CefrLevel != "" {
        args = append(args, req.CefrLevel)
        conditions = append(conditions, fmt.Sprintf("cefr_level = $%d", len(args)))
    }

    // Check if there are fields to update
    if len(conditions) == 0 {
        return nil, fmt.Errorf("no fields to update")
    }

    // Add updated_at condition
    conditions = append(conditions, "updated_at = now()")

    // Construct the full query
    query += " " + strings.Join(conditions, ", ")
    query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

    // Add ID to arguments
    args = append(args, req.Id)

    // Debugging: Log query and args
    fmt.Printf("Generated Query: %s\nArgs: %v\n", query, args)

    // Execute the query
    _, err := r.db.Exec(query, args...)
    if err != nil {
        return nil, fmt.Errorf("failed to update certificate: %w", err)
    }

    return &pb.Void{}, nil
}

func (r *CertificateRepository) Delete(req *pb.ById) (*pb.Void, error) {
	query := `UPDATE certificates SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0`

	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *CertificateRepository) GetById(req *pb.ById) (*pb.CertificateRes, error) {
	res := &pb.CertificateRes{}

	query := `SELECT 
				id, 
				name, 
 				ielts_score::TEXT AS ielts_score,			
				cefr_level, 
				certificate_url, 
				to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at
			FROM 
				certificates 
			WHERE 
				id = $1 AND deleted_at = 0`

	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.Name,
		&res.IeltsScore,
		&res.CefrLevel,
		&res.CertificateUrl,
		&res.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("certificate not found")
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *CertificateRepository) GetList(req *pb.GetListCertificateReq) (*pb.GetListCertificateRes, error) {
	res := &pb.GetListCertificateRes{}

	query := `SELECT
				COUNT(id) OVER () as total_count,
				id, 
				name, 
				ielts_score::TEXT AS ielts_score,
				cefr_level, 
				certificate_url, 
				to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at
			FROM
				certificates
			WHERE
				deleted_at = 0
`

	var args []interface{}
	var conditions []string

	if req.Name != "" && req.Name != "string" {
		args = append(args, "%"+req.Name+"%")
		conditions = append(conditions, "name ILIKE $"+strconv.Itoa(len(args)))
	}
	if req.IeltsScore > 0 {
		args = append(args, req.IeltsScore)
		conditions = append(conditions, "ielts_score <= $"+strconv.Itoa(len(args)))
	}
	if req.CefrLevel > "" && req.CefrLevel != "string" {
		args = append(args, req.CefrLevel)
		conditions = append(conditions, "cefr_level >= $"+strconv.Itoa(len(args)))
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, req.Filter.Limit, req.Filter.Offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var count int32
	for rows.Next() {
		var certificate pb.CertificateRes
		err := rows.Scan(
			&count,
			&certificate.Id,
			&certificate.Name,
			&certificate.IeltsScore,
			&certificate.CefrLevel,
			&certificate.CertificateUrl,
			&certificate.CreatedAt,
		)
		if err != nil {
			log.Printf("Error in row scan: %v", err)
			return nil, err
		}

		res.Certificates = append(res.Certificates, &certificate)
		res.TotalCount = count
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return res, nil
}
