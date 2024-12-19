package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
)

type TeacherRepository struct {
	db *sql.DB
}

func NewTeacherRepository(db *sql.DB) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (r *TeacherRepository) Create(req *pb.CreateTeacher) (*pb.Void, error) {
	query := `INSERT INTO teachers (name, experience_years, ielts_score, profile_picture_url, contact, graduated_students)
              VALUES ($1::jsonb, $2, $3, $4, $5, $6)`

	nameJson, err := json.Marshal(req.Name)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(query, string(nameJson), req.ExperienceYears, req.IeltsScore, req.ProfilePictureUrl, req.Contact, req.GraduatedStudents)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *TeacherRepository) Update(req *pb.UpdateTeacher) (*pb.Void, error) {
	query := `UPDATE teachers SET`
	var args []interface{}
	var conditions []string

	if req.Name != nil{
		nameJson, err := json.Marshal(req.Name)
		if err != nil {
			return nil, fmt.Errorf("marshal name failed: %w", err)
		}
		args = append(args, string(nameJson))
		conditions = append(conditions, "name=$"+strconv.Itoa(len(args)))
	}
	if req.ExperienceYears != "" && req.ExperienceYears != "string" {
		args = append(args, req.ExperienceYears)
		conditions = append(conditions, "experience_years=$"+strconv.Itoa(len(args)))
	}
	if req.IeltsScore != "" && req.IeltsScore != "string" {
		args = append(args, req.IeltsScore)
		conditions = append(conditions, "ielts_score=$"+strconv.Itoa(len(args)))
	}
	if req.ProfilePictureUrl != "" {
		args = append(args, req.ProfilePictureUrl)
		conditions = append(conditions, "profile_picture_url=$"+strconv.Itoa(len(args)))
	}
	if req.Contact != "" && req.Contact != "string" {
		args = append(args, req.Contact)
		conditions = append(conditions, "contact=$"+strconv.Itoa(len(args)))
	}
	if req.GraduatedStudents != "" && req.GraduatedStudents != "string" {
		args = append(args, req.GraduatedStudents)
		conditions = append(conditions, "graduated_students=$"+strconv.Itoa(len(args)))
	}

	conditions = append(conditions, "updated_at=NOW()")
	query += " " + strings.Join(conditions, ", ")
	query += " WHERE id=$" + strconv.Itoa(len(args)+1) + " AND deleted_at=0"

	args = append(args, req.Id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *TeacherRepository) Delete(req *pb.ById) (*pb.Void, error) {
	query := `UPDATE teachers SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0`
	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *TeacherRepository) GetById(req *pb.ById) (*pb.TeacherRes, error) {
	res := &pb.TeacherRes{}
	query := `SELECT 
                id, 
                name::jsonb,
                experience_years, 
                ielts_score::TEXT AS ielts_score, 
                profile_picture_url, 
                contact, 
                graduated_students, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as created_at
            FROM 
                teachers 
            WHERE id = $1 AND deleted_at = 0`

    var nameJson string
	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&nameJson,
		&res.ExperienceYears,
		&res.IeltsScore,
		&res.ProfilePictureUrl,
		&res.Contact,
		&res.GraduatedStudents,
		&res.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("teacher not found")
	} else if err != nil {
		return nil, err
	}

	// Unmarshal JSON into MultilingualField
	err = json.Unmarshal([]byte(nameJson), &res.Name)
	if err != nil {
		return nil, fmt.Errorf("unmarshal name failed: %w", err)
	}

	return res, nil
}

func (r *TeacherRepository) GetList(req *pb.GetListTeacherReq) (*pb.GetListTeacherRes, error) {
	res := &pb.GetListTeacherRes{}
	query := `SELECT 
                COUNT(id) OVER() AS total_count, 
                id, 
                name::jsonb,
                experience_years, 
                ielts_score::TEXT AS ielts_score, 
                profile_picture_url, 
                contact, 
                graduated_students, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as created_at 
            FROM 
                teachers 
            WHERE deleted_at = 0`

	var args []interface{}
	filters := []string{}

	// Filter by language if specified
	if req.Language != "" {
		filters = append(filters, "name->>$1 IS NOT NULL")
		args = append(args, req.Language)
	}

	if len(filters) > 0 {
		query += " AND " + strings.Join(filters, " AND ")
	}

	// Pagination
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, req.Filter.Limit, req.Filter.Offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nameJson string
		var teacher pb.TeacherRes
		err := rows.Scan(
			&res.TotalCount,
			&teacher.Id,
			&nameJson,
			&teacher.ExperienceYears,
			&teacher.IeltsScore,
			&teacher.ProfilePictureUrl,
			&teacher.Contact,
			&teacher.GraduatedStudents,
			&teacher.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Unmarshal JSON into MultilingualField
		err = json.Unmarshal([]byte(nameJson), &teacher.Name)
		if err != nil {
			return nil, fmt.Errorf("unmarshal name failed: %w", err)
		}

		res.Teacher = append(res.Teacher, &teacher)
	}

	return res, nil
}
