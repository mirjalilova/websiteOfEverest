package repository

import (
	"database/sql"
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
	query := `INSERT INTO teachers (name, experience_years, ielts_score, profile_picture_url, contact, graduated_students, bio)
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, req.Name, req.ExperienceYears, req.IeltsScore, req.ProfilePictureUrl, req.Contact, req.GraduatedStudents, req.Bio)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *TeacherRepository) Update(req *pb.UpdateTeacher) (*pb.Void, error) {
	query := `UPDATE teachers SET`
	var args []interface{}
	var conditions []string

	if req.Name != "" && req.Name != "string" {
		args = append(args, req.Name)
		conditions = append(conditions, "name=$"+strconv.Itoa(len(args)))
	}
	if req.ExperienceYears != 0 {
		args = append(args, req.ExperienceYears)
		conditions = append(conditions, "experience_years=$"+strconv.Itoa(len(args)))
	}
	if req.IeltsScore != 0 {
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
	if req.GraduatedStudents != 0 {
		args = append(args, req.GraduatedStudents)
		conditions = append(conditions, "graduated_students=$"+strconv.Itoa(len(args)))
	}
	if req.Bio != "" && req.Bio != "string" {
		args = append(args, req.Bio)
		conditions = append(conditions, "bio=$"+strconv.Itoa(len(args)))
	}

	conditions = append(conditions, "updated_at=NOW()")
	query += strings.Join(conditions, ", ")
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
                name, 
                experience_years, 
                ielts_score, 
                profile_picture_url, 
                contact, 
                graduated_students, 
                bio, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as created_at,
            FROM 
                teachers 
            WHERE id = $1 AND deleted_at = 0`

	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.Name,
		&res.ExperienceYears,
		&res.IeltsScore,
		&res.ProfilePictureUrl,
		&res.Contact,
		&res.GraduatedStudents,
		&res.Bio,
		&res.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("teacher not found")
	} else if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *TeacherRepository) GetList(req *pb.GetListTeacherReq) (*pb.GetListTeacherRes, error) {
	res := &pb.GetListTeacherRes{}
	query := `SELECT 
                COUNT(id) OVER() AS total_count, 
                id, 
                name, 
                experience_years, 
                ielts_score, 
                profile_picture_url, 
                contact, 
                graduated_students, 
                bio, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as created_at 
            FROM 
                teachers 
            WHERE deleted_at = 0`

	var args []interface{}
	filters := []string{}

	if req.Name != "" && req.Name != "strinh" {
		filters = append(filters, "name ILIKE '%' || $"+strconv.Itoa(len(args)+1)+" || '%'")
		args = append(args, req.Name)
	}
	if req.ExperienceYearsMin != 0 {
		filters = append(filters, "experience_years >= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.ExperienceYearsMin)
	}
	if req.ExperienceYearsMax != 0 {
		filters = append(filters, "experience_years <= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.ExperienceYearsMax)
	}
	if req.IeltsScoreMin != 0 {
		filters = append(filters, "ielts_score >= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.IeltsScoreMin)
	}
	if req.IeltsScoreMax != 0 {
		filters = append(filters, "ielts_score <= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.IeltsScoreMax)
	}

	if len(filters) > 0 {
		query += " AND " + strings.Join(filters, " AND ")
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var count int32
	for rows.Next() {
		var teacher pb.TeacherRes
		err := rows.Scan(
			&count,
			&teacher.Id,
			&teacher.Name,
			&teacher.ExperienceYears,
			&teacher.IeltsScore,
			&teacher.ProfilePictureUrl,
			&teacher.Contact,
			&teacher.GraduatedStudents,
			&teacher.Bio,
			&teacher.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Teacher = append(res.Teacher, &teacher)
		res.TotalCount = count
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}