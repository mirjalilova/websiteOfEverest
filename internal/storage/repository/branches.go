package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
)

type BranchesRepo struct {
	db *sql.DB
}

func NewBranchesRepository(db *sql.DB) *BranchesRepo {
	return &BranchesRepo{db: db}
}

// Create Branch
func (r *BranchesRepo) Create(req *pb.CreateBranches) (*pb.Void, error) {
	query := `INSERT INTO branches (name, google_url, yandex_url, contact, img_url) 
			  VALUES ($1::jsonb, $2, $3, $4, $5)`

	nameJson, err := json.Marshal(req.Name)
	if err != nil {
		return nil, fmt.Errorf("marshal name failed: %w", err)
	}

	_, err = r.db.Exec(query, string(nameJson), req.GoogleUrl, req.YandexUrl, req.Contact, req.ImgUrl)
	return nil, err
}

// Update Branch
func (r *BranchesRepo) Update(req *pb.UpdateBranches) (*pb.Void, error) {
	query := `UPDATE branches SET `
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
	if req.GoogleUrl != "" && req.GoogleUrl != "string" {
		args = append(args, req.GoogleUrl)
		updates = append(updates, "google_url=$"+strconv.Itoa(len(args)))
	}
	if req.YandexUrl != "" && req.YandexUrl != "string" {
		args = append(args, req.YandexUrl)
		updates = append(updates, "yandex_url=$"+strconv.Itoa(len(args)))
	}
	if req.Contact != "" && req.Contact != "string" {
		args = append(args, req.Contact)
		updates = append(updates, "contact=$"+strconv.Itoa(len(args)))
	}
	if req.ImgUrl != "" && req.ImgUrl != "string" {
		args = append(args, req.ImgUrl)
		updates = append(updates, "img_url=$"+strconv.Itoa(len(args)))
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(updates, ", ") + ", updated_at=now() WHERE id=$" + strconv.Itoa(len(args)+1)
	args = append(args, req.Id)

	_, err :=r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

// Get Branch By ID
func (r *BranchesRepo) GetById(req *pb.ById) (*pb.BranchesRes, error) {
	query := `SELECT id, name::jsonb, google_url, yandex_url, contact, img_url, 
			  to_char(created_at, 'YYYY-MM-DD HH24:MI') 
			  FROM branches WHERE id=$1 AND deleted_at=0`

	res := &pb.BranchesRes{}
	var nameJson string

	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id, &nameJson, &res.GoogleUrl, &res.YandexUrl, &res.Contact, &res.ImgUrl, &res.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("branch not found")
	}
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into MultilingualField
	err = json.Unmarshal([]byte(nameJson), &res.Name)
	if err != nil {
		return nil, fmt.Errorf("unmarshal name failed: %w", err)
	}

	return res, nil
}

// Get Branch List
func (r *BranchesRepo) GetList(req *pb.GetListBranchesReq) (*pb.GetListBranchesRes, error) {
	query := `SELECT COUNT(*) OVER (), id, name::jsonb, google_url, yandex_url, 
			  contact, img_url, to_char(created_at, 'YYYY-MM-DD HH24:MI') 
			  FROM branches WHERE deleted_at=0`
	var args []interface{}
	var filters []string

	// Filter by language if specified
	// if req.Language != "" {
	// 	filters = append(filters, "name->>$1 IS NOT NULL")
	// 	args = append(args, req.Language)
	// }

	// Filter by branch name if specified
	if req.Language != "" {
		filters = append(filters, "name->>$1 ILIKE $2")
		args = append(args, req.Language, "%"+req.Language+"%")
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

	res := &pb.GetListBranchesRes{}
	for rows.Next() {
		var branch pb.BranchesRes
		var nameJson string
		err := rows.Scan(
			&res.TotalCount, &branch.Id, &nameJson, &branch.GoogleUrl,
			&branch.YandexUrl, &branch.Contact, &branch.ImgUrl, &branch.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Unmarshal JSON into MultilingualField
		err = json.Unmarshal([]byte(nameJson), &branch.Name)
		if err != nil {
			return nil, fmt.Errorf("unmarshal name failed: %w", err)
		}

		res.Branches = append(res.Branches, &branch)
	}

	return res, nil
}


// Delete Branch
func (r *BranchesRepo) Delete(req *pb.ById) (*pb.Void, error) {
	query := `UPDATE branches SET deleted_at=EXTRACT(EPOCH FROM NOW()) WHERE id=$1 AND deleted_at=0`
	_, err := r.db.Exec(query, req.Id)
	return &pb.Void{}, err
}
