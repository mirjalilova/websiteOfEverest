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
				(name, ielts_score, cefr_level, description, certificate_url)
			VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, req.Name, req.IeltsScore, req.CefrLevel, req.Description, req.CertificateUrl)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *CertificateRepository) Update(req *pb.UpdateCertificate) (*pb.Void, error) {
	query := `UPDATE certificates SET`

	var args []interface{}
	var conditions []string

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
	if req.Description != "string" && req.Description != "" {
		args = append(args, req.Description)
		conditions = append(conditions, fmt.Sprintf("description = $%d", len(args)))
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

	args = append(args, req.Id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
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
				CAST(ielts_score AS FLOAT) as ielts_score, 
				cefr_level, 
				description, 
				certificate_url, 
				to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at,
			FROM 
				certificates 
			WHERE 
				id = $1 AND deleted_at = 0`

	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.Name,
		&res.IeltsScore,
		&res.CefrLevel,
		&res.Description,
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
				CAST(ielts_score AS FLOAT) as ielts_score,
				cefr_level, 
				description, 
				certificate_url, 
				to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at
			FROM
				certificates
			WHERE
				deleted_at = 0;
`

	rows, err := r.db.Query(query)
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
			&certificate.Description,
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
