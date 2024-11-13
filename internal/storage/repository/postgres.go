package repository

import (
	"database/sql"

	"github.com/mirjalilova/websiteOfEverest/internal/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db           *sql.DB
	BranchesS    storage.BranchesI
	CertificateS storage.CertificateI
	CourseItemS  storage.CourseItemI
	CourseS      storage.CourseI
	DashboardS   storage.DashboardI
	GalleryS     storage.GalleryI
	TeacherS     storage.TeacherI
}

func NewPostgresStorage(db *sql.DB) (*Storage, error) {

	branches := NewBranchesRepository(db)
	certificate := NewCertificateRepository(db)
	courseItem := NewCourseItemRepository(db)
	course := NewCourseRepository(db)
	dashboard := NewDashboardRepository(db)
	gallery := NewGalleryRepository(db)
	teacher := NewTeacherRepository(db)

	return &Storage{
		Db:           db,
		BranchesS:    branches,
		CertificateS: certificate,
		CourseItemS:  courseItem,
		CourseS:      course,
		DashboardS:   dashboard,
		GalleryS:     gallery,
		TeacherS:     teacher,
	}, nil
}
