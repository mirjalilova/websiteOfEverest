package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
)

type CourseItemRepository struct {
	db *sql.DB
}

func NewCourseItemRepository(db *sql.DB) *CourseItemRepository {
	return &CourseItemRepository{db: db}
}

func (r *CourseItemRepository) Create(req *pb.CreateCourseItem) (*pb.Void, error) {
	query := `INSERT INTO course_items
                (course_id, description, price, days_per_week, lesson_hours, week_days)
              VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.Exec(query, req.CourseId, req.Description, req.Price, req.DaysPerWeek, req.LessonHours, req.WeekDays)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *CourseItemRepository) Update(req *pb.UpdateCourseItem) (*pb.Void, error) {
    query := `UPDATE course_items SET`

    var args []interface{}
    var updates []string

    // Check each field and append to updates if necessary
    if req.CourseId != "" && req.CourseId != "string" {
        args = append(args, req.CourseId)
        updates = append(updates, fmt.Sprintf("course_id=$%d", len(args)))
    }
    if req.Description != "" && req.Description != "string" {
        args = append(args, req.Description)
        updates = append(updates, fmt.Sprintf("description=$%d", len(args)))
    }
    if req.Price != 0 {
        args = append(args, req.Price)
        updates = append(updates, fmt.Sprintf("price=$%d", len(args)))
    }
    if req.DaysPerWeek != 0 {
        args = append(args, req.DaysPerWeek)
        updates = append(updates, fmt.Sprintf("days_per_week=$%d", len(args)))
    }
    if req.LessonHours != 0 {
        args = append(args, req.LessonHours)
        updates = append(updates, fmt.Sprintf("lesson_hours=$%d", len(args)))
    }
    if req.WeekDays != "string" && req.WeekDays != "" {
        args = append(args, req.WeekDays)
        updates = append(updates, fmt.Sprintf("week_days=$%d", len(args)))
    }

    // Check if we have any fields to update
    if len(updates) == 0 {
        return nil, fmt.Errorf("no fields to update")
    }

    // Always add updated_at condition
    updates = append(updates, "updated_at=now()")

    // Join the updates and construct the query
    query += " " + strings.Join(updates, ", ")
    query += " WHERE id=$" + strconv.Itoa(len(args)+1) + " AND deleted_at=0"

    // Add ID to arguments
    args = append(args, req.Id)

    // Debugging: Print the generated query and arguments
    fmt.Printf("Generated Query: %s\nArgs: %v\n", query, args)

    // Execute the query
    _, err := r.db.Exec(query, args...)
    if err != nil {
        return nil, fmt.Errorf("failed to update course item: %w", err)
    }

    return &pb.Void{}, nil
}


func (r *CourseItemRepository) Delete(req *pb.ById) (*pb.Void, error) {
	query := `UPDATE course_items SET deleted_at=EXTRACT(EPOCH FROM NOW()) WHERE id=$1 AND deleted_at=0`

	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *CourseItemRepository) GetById(req *pb.ById) (*pb.CourseItemRes, error) {
	res := &pb.CourseItemRes{}

	query := `SELECT 
                id, 
                course_id, 
                description, 
                price, 
                days_per_week, 
                lesson_hours, 
                week_days, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as created_at 
              FROM 
                course_items 
              WHERE 
                id=$1 AND deleted_at=0`

	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.CourseId,
		&res.Description,
		&res.Price,
		&res.DaysPerWeek,
		&res.LessonHours,
		&res.WeekDays,
		&res.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("course item not found")
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *CourseItemRepository) GetList(req *pb.GetListCourseItemReq) (*pb.GetListCourseItemRes, error) {
	res := &pb.GetListCourseItemRes{}
	var args []interface{}
	var conditions []string

	query := `SELECT 
                COUNT(id) OVER() as total_count,
                id, 
                course_id, 
                description, 
                price, 
                days_per_week, 
                lesson_hours, 
                week_days, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as created_at 
              FROM 
                course_items 
              WHERE 
                deleted_at=0`

	if req.CourseId != "" && req.CourseId != "string" {
		args = append(args, req.CourseId)
		conditions = append(conditions, "course_id=$"+strconv.Itoa(len(args)))
	}
	if req.MaxPrice > 0 {
		args = append(args, req.MaxPrice)
		conditions = append(conditions, "price <= $"+strconv.Itoa(len(args)))
	}
	if req.MinPrice > 0 {
		args = append(args, req.MinPrice)
		conditions = append(conditions, "price >= $"+strconv.Itoa(len(args)))
	}
	if req.DaysPerWeek > 0 {
		args = append(args, req.DaysPerWeek)
		conditions = append(conditions, "days_per_week=$"+strconv.Itoa(len(args)))
	}
	if req.LessonHours > 0 {
		args = append(args, req.LessonHours)
		conditions = append(conditions, "lesson_hours = $"+strconv.Itoa(len(args)))
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args) +1, len(args)+2)
	args = append(args, req.Filter.Limit, req.Filter.Offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var count int32
	for rows.Next() {
		var item pb.CourseItemRes
		err := rows.Scan(
			&count,
			&item.Id,
			&item.CourseId,
			&item.Description,
			&item.Price,
			&item.DaysPerWeek,
			&item.LessonHours,
			&item.WeekDays,
			&item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.CourseItems = append(res.CourseItems, &item)
		res.TotalCount = count
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}
