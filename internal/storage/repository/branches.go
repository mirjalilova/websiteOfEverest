package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
)

type BranchesRepository struct {
	db *sql.DB
}

func NewBranchesRepository(db *sql.DB) *BranchesRepository {
	return &BranchesRepository{db: db}
}

func (r *BranchesRepository) Create(req *pb.CreateBranches) (*pb.Void, error) {
	query := `INSET INTO branches
                (name, description, google_url, yandex_url, contact)
                VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query, req.Name, req.Description, req.GoogleUrl, req.YandexUrl, req.Contact)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *BranchesRepository) Update(req *pb.UpdateBranches) (*pb.Void, error) {
	query := `UPDATE branches SET`

	var args []interface{}
	var conditions []string

	if req.Name != "" && req.Name != "string" {
		args = append(args, req.Name)
		conditions = append(conditions, "name=$"+strconv.Itoa(len(args)))
	}
	if req.Description != "" && req.Description != "string" {
		args = append(args, req.Description)
		conditions = append(conditions, "description=$"+strconv.Itoa(len(args)))
	}
	if req.GoogleUrl != "" && req.GoogleUrl != "string" {
		args = append(args, req.GoogleUrl)
		conditions = append(conditions, "google_url=$"+strconv.Itoa(len(args)))
	}
	if req.YandexUrl != "" && req.YandexUrl != "string" {
		args = append(args, req.YandexUrl)
		conditions = append(conditions, "yandex_url=$"+strconv.Itoa(len(args)))
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

func (r *BranchesRepository) Delete(req *pb.ById) (*pb.Void, error) {
	query := `UPDATE branches SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0`

	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *BranchesRepository) GetById(req *pb.ById) (*pb.BranchesRes, error) {
	res := &pb.BranchesRes{}

	query := `SELECT 
                id, 
                name, 
                description, 
                google_url, 
                yandex_url, 
                contact, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at, 
            FROM 
				branches 
			WHERE 
				id = $1 AND deleted_at = 0`

	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.Name,
		&res.Description,
		&res.GoogleUrl,
		&res.YandexUrl,
		&res.Contact,
		&res.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("branch not found")
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *BranchesRepository) GetList(req *pb.GetListBranchesReq) (*pb.GetListBranchesRes, error) {
	res := &pb.GetListBranchesRes{}

    query := `SELECT
                COUNT(id) OVER () AS total_count,
                id, 
                name, 
                description, 
                google_url, 
                yandex_url,
                contact, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at, 
            FROM
                branches
            WHERE
                deleted_at = 0`
    
    rows, err := r.db.Query(query)
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var count int32 
    for rows.Next() {
        var branch pb.BranchesRes
        err := rows.Scan(
            &count,
            &branch.Id,
            &branch.Name,
            &branch.Description,
            &branch.GoogleUrl,
            &branch.YandexUrl,
            &branch.Contact,
            &branch.CreatedAt,
        )
        if err!= nil {
            return nil, err
        }
        res.Branches = append(res.Branches, &branch)
        res.TotalCount = count
    }

    if rows.Err() != nil {
        return nil, err
    }

	return res, nil
}
