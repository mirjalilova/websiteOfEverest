package repository

import (
	"database/sql"
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
                (name, duration)
            VALUES($1, $2)`

	_, err := r.db.Exec(query, req.Name, req.Duration)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *CourseRepository) Update(req *pb.UpdateCourse) (*pb.Void, error) {
	query := `UPDATE courses SET`

	var args []interface{}
	var conditions []string

	if req.Name != "string" && req.Name != "" {
		args = append(args, req.Name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
	}

	if req.Duration != 0 {
		args = append(args, req.Duration)
		conditions = append(conditions, fmt.Sprintf("duration = $%d", len(args)))
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

func (r *CourseRepository) Delete(req *pb.ById) (*pb.Void, error) {
    query := `UPDATE courses SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0`

	_, err := r.db.Exec(query, req.Id)
	if err!= nil {
        return nil, err
    }

	return &pb.Void{}, nil
}

func (r *CourseRepository) GetById(req *pb.ById) (*pb.CourseRes, error) {
	res := &pb.CourseRes{}

    query := `SELECT
                id, 
                name, 
                duration, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at,
            FROM 
                courses
            WHERE 
                id = $1 AND deleted_at = 0`
    err := r.db.QueryRow(query, req.Id).Scan(
        &res.Id,
        &res.Name,
        &res.Duration,
        &res.CreatedAt,
    )
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("course not found")
    } else if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *CourseRepository) GetList(req *pb.GetListCourseReq) (*pb.GetListCourseRes, error) {
	res := &pb.GetListCourseRes{}

    query := `SELECT
                COUNT(id) OVER () AS total_count,
                id, 
                name, 
                duration, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as formatted_created_at,
            FROM
                courses
            WHERE deleted_at = 0`
    
    rows, err := r.db.Query(query)
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var count int32
    for rows.Next() {
        var course pb.CourseRes
        err := rows.Scan(
            &count,
            &course.Id,
            &course.Name,
            &course.Duration,
            &course.CreatedAt,
        )
        if err!= nil {
            return nil, err
        }

        res.Courses = append(res.Courses, &course)
        res.TotalCount = count
    }

    if rows.Err()!= nil {
        return nil, rows.Err()
    }

	return res, nil
}
