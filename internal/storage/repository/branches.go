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
	query := `INSERT INTO branches
                (name, description, google_url, yandex_url, contact, img_url)
                VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.Exec(query, req.Name, req.Description, req.GoogleUrl, req.YandexUrl, req.Contact, req.ImgUrl)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *BranchesRepository) Update(req *pb.UpdateBranches) (*pb.Void, error) {
    query := `UPDATE branches SET `
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
    if req.ImgUrl != "" && req.ImgUrl != "string" {
        args = append(args, req.ImgUrl)
        conditions = append(conditions, "img_url=$"+strconv.Itoa(len(args)))
    }
	if req.Contact != "" && req.Contact != "string" {
		args = append(args, req.Contact)
		conditions = append(conditions, "contact=$"+strconv.Itoa(len(args)))
	}

    if len(conditions) == 0 {
        return nil, fmt.Errorf("no fields to update")
    }

    // Append `updated_at` to conditions
    conditions = append(conditions, "updated_at = now()")

    // Join conditions and construct the full query
    query += strings.Join(conditions, ", ")
    query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

    // Add ID to arguments
    args = append(args, req.Id)

    // Log the query for debugging
    fmt.Printf("Generated Query: %s\nArgs: %v\n", query, args)

    // Execute the query
    _, err := r.db.Exec(query, args...)
    if err != nil {
        return nil, fmt.Errorf("failed to update branch: %w", err)
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
				img_url, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at
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
		&res.ImgUrl,
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
				img_url,
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at
            FROM
                branches
            WHERE
                deleted_at = 0`

	var args []interface{}
	var conditions []string

	if req.Name != "" && req.Name != "string" {
		args = append(args, "%"+req.Name+"%")
		conditions = append(conditions, "name ILIKE $"+strconv.Itoa(len(args)))
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
		var branch pb.BranchesRes
		err := rows.Scan(
			&count,
			&branch.Id,
			&branch.Name,
			&branch.Description,
			&branch.GoogleUrl,
			&branch.YandexUrl,
			&branch.Contact,
			&branch.ImgUrl,
			&branch.CreatedAt,
		)
		if err != nil {
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
