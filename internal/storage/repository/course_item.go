package repository

import (
	"database/sql"
	"encoding/json"
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

// Create Course Item
func (r *CourseItemRepository) Create(req *pb.CreateCourseItem) (*pb.Void, error) {
	query := `INSERT INTO course_items
                (course_id, description, price, days_per_week, lesson_hours, week_days, picture_url) 
              VALUES ($1, $2::jsonb, $3, $4, $5, $6::jsonb, $7)`

	descriptionJson, err := json.Marshal(req.Descritption)
	if err != nil {
		return nil, fmt.Errorf("marshal description failed: %w", err)
	}

	weekDaysJson, err := json.Marshal(req.WeekDays)
	if err != nil {
		return nil, fmt.Errorf("marshal week_days failed: %w", err)
	}

	_, err = r.db.Exec(query, req.CourseId, string(descriptionJson), req.Price, req.DaysPerWeek, req.LessonHours, string(weekDaysJson), req.PictureUrl)
	return nil, err
}

// Update Course Item
func (r *CourseItemRepository) Update(req *pb.UpdateCourseItem) (*pb.Void, error) {
	query := `UPDATE course_items SET `
	var updates []string
	var args []interface{}

	if req.CourseId != "" {
		args = append(args, req.CourseId)
		updates = append(updates, "course_id=$"+strconv.Itoa(len(args)))
	}
	if req.Descritption != nil {
		descriptionJson, err := json.Marshal(req.Descritption)
		if err != nil {
			return nil, fmt.Errorf("marshal description failed: %w", err)
		}
		args = append(args, string(descriptionJson))
		updates = append(updates, "description=$"+strconv.Itoa(len(args)))
	}
	if req.Price > 0 {
		args = append(args, req.Price)
		updates = append(updates, "price=$"+strconv.Itoa(len(args)))
	}
	if req.DaysPerWeek > 0 {
		args = append(args, req.DaysPerWeek)
		updates = append(updates, "days_per_week=$"+strconv.Itoa(len(args)))
	}
	if req.LessonHours > 0 {
		args = append(args, req.LessonHours)
		updates = append(updates, "lesson_hours=$"+strconv.Itoa(len(args)))
	}
	if req.WeekDays != nil {
		weekDaysJson, err := json.Marshal(req.WeekDays)
		if err != nil {
			return nil, fmt.Errorf("marshal week_days failed: %w", err)
		}
		args = append(args, string(weekDaysJson))
		updates = append(updates, "week_days=$"+strconv.Itoa(len(args)))
	}
	if req.PictureUrl != "" {
		args = append(args, req.PictureUrl)
		updates = append(updates, "picture_url=$"+strconv.Itoa(len(args)))
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(updates, ", ") + ", updated_at=now() WHERE id=$" + strconv.Itoa(len(args)+1) + " AND deleted_at=0"
	args = append(args, req.Id)

	_, err := r.db.Exec(query, args...)
	return nil, err
}

// Delete Course Item
func (r *CourseItemRepository) Delete(req *pb.ById) (*pb.Void, error) {
	query := `UPDATE course_items SET deleted_at=EXTRACT(EPOCH FROM NOW()) WHERE id=$1 AND deleted_at=0`
	_, err := r.db.Exec(query, req.Id)
	return nil, err
}

// Get Course Item by ID
func (r *CourseItemRepository) GetById(req *pb.ById) (*pb.CourseItemRes, error) {
	query := `SELECT id, course_id, description::jsonb, price, days_per_week, lesson_hours, 
			  week_days::jsonb, picture_url, to_char(created_at, 'YYYY-MM-DD HH24:MI') 
			  FROM course_items WHERE id=$1 AND deleted_at=0`

	res := &pb.CourseItemRes{}
	var descriptionJson, weekDaysJson string

	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id, &res.CourseId, &descriptionJson, &res.Price, &res.DaysPerWeek,
		&res.LessonHours, &weekDaysJson, &res.PictureUrl, &res.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("course item not found")
	}
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(descriptionJson), &res.Descritption); err != nil {
		return nil, fmt.Errorf("unmarshal description failed: %w", err)
	}
	if err = json.Unmarshal([]byte(weekDaysJson), &res.WeekDays); err != nil {
		return nil, fmt.Errorf("unmarshal week_days failed: %w", err)
	}

	return res, nil
}

// Get List of Course Items
func (r *CourseItemRepository) GetList(req *pb.GetListCourseItemReq) (*pb.GetListCourseItemRes, error) {
	query := `SELECT COUNT(*) OVER (), id, course_id, description::jsonb, price, days_per_week, 
			  lesson_hours, week_days::jsonb, picture_url, to_char(created_at, 'YYYY-MM-DD HH24:MI') 
			  FROM course_items WHERE deleted_at=0`
	var args []interface{}
	var filters []string

	// Filter by language if specified
	if req.Language != "" {
		filters = append(filters, "description->>$1 IS NOT NULL")
		args = append(args, req.Language)
	}

	if req.Language != "" {
		filters = append(filters, "week_days->>$1 IS NOT NULL")
		args = append(args, req.Language)
	}

	if req.CourseId != "" {
		args = append(args, req.CourseId)
		filters = append(filters, "course_id=$"+strconv.Itoa(len(args)))
	}
	if req.MaxPrice > 0 {
		args = append(args, req.MaxPrice)
		filters = append(filters, "price<=$"+strconv.Itoa(len(args)))
	}
	if req.MinPrice > 0 {
		args = append(args, req.MinPrice)
		filters = append(filters, "price>=$"+strconv.Itoa(len(args)))
	}
	if req.DaysPerWeek > 0 {
		args = append(args, req.DaysPerWeek)
		filters = append(filters, "days_per_week=$"+strconv.Itoa(len(args)))
	}
	if req.LessonHours > 0 {
		args = append(args, req.LessonHours)
		filters = append(filters, "lesson_hours=$"+strconv.Itoa(len(args)))
	}

	if len(filters) > 0 {
		query += " AND " + strings.Join(filters, " AND ")
	}

	query += " LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, req.Filter.Limit, req.Filter.Offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := &pb.GetListCourseItemRes{}
	for rows.Next() {
		var item pb.CourseItemRes
		var descriptionJson, weekDaysJson string
		err := rows.Scan(
			&res.TotalCount, &item.Id, &item.CourseId, &descriptionJson, &item.Price,
			&item.DaysPerWeek, &item.LessonHours, &weekDaysJson, &item.PictureUrl, &item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal([]byte(descriptionJson), &item.Descritption); err != nil {
			return nil, fmt.Errorf("unmarshal description failed: %w", err)
		}
		if err = json.Unmarshal([]byte(weekDaysJson), &item.WeekDays); err != nil {
			return nil, fmt.Errorf("unmarshal week_days failed: %w", err)
		}

		res.CourseItems = append(res.CourseItems, &item)
	}

	return res, nil
}
