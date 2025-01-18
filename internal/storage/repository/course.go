package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) Create(req *pb.CreateCourse) (*pb.Void, error) {
	query := `INSERT INTO courses
                (name, duration, picture_url)
            VALUES($1::jsonb, $2::jsonb, $3)`

	nameJson, err := json.Marshal(req.Name)
	if err != nil {
		return nil, fmt.Errorf("marshal name failed: %w", err)
	}

	jsonDuration, err := json.Marshal(req.Duration)
	if err != nil {
		return nil, fmt.Errorf("marshal duration failed: %w", err)
	}

	_, err = r.db.Exec(query, string(nameJson), string(jsonDuration))
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *CourseRepository) Update(req *pb.UpdateCourse) (*pb.Void, error) {
	query := `UPDATE courses SET`

	var args []interface{}
	var conditions []string

	if req.Name != nil {
		nameJson, err := json.Marshal(req.Name)
		if err != nil {
			return nil, fmt.Errorf("marshal name failed: %w", err)
		}
		args = append(args, string(nameJson))
		conditions = append(conditions, "name=$"+strconv.Itoa(len(args)))
	}

	if req.Duration != nil {
		durationJson, err := json.Marshal(req.Duration)
		if err != nil {
			return nil, fmt.Errorf("marshal duration failed: %w", err)
		}
		args = append(args, string(durationJson))
		conditions = append(conditions, "duration=$"+strconv.Itoa(len(args)))
	}

	if req.PictureUrl != "" {
		args = append(args, req.PictureUrl)
		conditions = append(conditions, "picture_url=$"+strconv.Itoa(len(args)))
	}

	if len(conditions) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	conditions = append(conditions, "updated_at=now()")
	query += " " + strings.Join(conditions, ", ")
	query += " WHERE id=$" + strconv.Itoa(len(args)+1) + " AND deleted_at=0"

	args = append(args, req.Id)
	fmt.Printf("Generated Query: %s\nArgs: %v\n", query, args)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update course: %v", err)
	}

	return &pb.Void{}, nil
}

func (r *CourseRepository) Delete(req *pb.ById) (*pb.Void, error) {
	query := `UPDATE courses SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0`

	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *CourseRepository) GetById(req *pb.ById) (*pb.CourseRes, error) {
	res := &pb.CourseRes{}
	var nameJson string
	var durationJson string

	query := `SELECT
                id, 
                name::jsonb, 
                duration::jsonb,
				picture_url, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at
            FROM 
                courses
            WHERE 
                id = $1 AND deleted_at = 0`
	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&nameJson,
		&durationJson,
		&res.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("course not found")
	} else if err != nil {
		return nil, err
	}

	// Unmarshal JSON into MultilingualField
	err = json.Unmarshal([]byte(nameJson), &res.Name)
	if err != nil {
		return nil, fmt.Errorf("unmarshal name failed: %w", err)
	}

	err = json.Unmarshal([]byte(durationJson), &res.Duration)
	if err != nil {
		return nil, fmt.Errorf("unmarshal duration failed: %w", err)
	}

	return res, nil
}

func (r *CourseRepository) GetList(req *pb.GetListCourseReq) (*pb.GetListCourseRes, error) {
	query := `SELECT
                COUNT(id) OVER () AS total_count,
                id, 
                name::jsonb, 
                duration::jsonb,
				picture_url, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at
            FROM
                courses
            WHERE deleted_at = 0`

	var args []interface{}
	var filters []string

	// Filter by language if specified
	if req.Language != "" {
		filters = append(filters, "name->>$1 IS NOT NULL")
		args = append(args, req.Language)
	}

	if req.Language != "" {
		filters = append(filters, "duration->>$1 IS NOT NULL")
		args = append(args, req.Language)
	}

	if len(filters) > 0 {
		query += " AND " + strings.Join(filters, " AND ")
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, req.Filter.Limit, req.Filter.Offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var count int32
	res := &pb.GetListCourseRes{}

	for rows.Next() {
		var nameJson string
		var durationJson string
		var course pb.CourseRes
		err = rows.Scan(
			&count,
			&course.Id,
			&nameJson,
			&durationJson,
			&course.PictureUrl,
			&course.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Unmarshal JSON into MultilingualField
		err = json.Unmarshal([]byte(nameJson), &course.Name)
		if err != nil {
			return nil, fmt.Errorf("unmarshal name failed: %w", err)
		}

		err = json.Unmarshal([]byte(durationJson), &course.Duration)
		if err != nil {
			return nil, fmt.Errorf("unmarshal duration failed: %w", err)
		}

		res.Courses = append(res.Courses, &course)
	}
	

	return res, nil
}
