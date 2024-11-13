package clients

import (
	"github.com/mirjalilova/websiteOfEverest/config"
	"github.com/mirjalilova/websiteOfEverest/internal/service"
	"github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type GrpcClients struct {
	Branches    *service.BranchesService
	Certificate *service.CertificateService
	CourseItem  *service.CourseItemService
	Course      *service.CourseService
	// Dashboard      pb.DashboardServiceClient
	Gallery *service.GalleryService
	Teacher *service.TeacherService
}

func NewGrpcClients(cfg *config.Config, conn *repository.Storage) (*GrpcClients, error) {

	return &GrpcClients{
		Branches:    service.NewdBranchesService(conn),
		Certificate: service.NewCertificateService(conn),
		CourseItem:  service.NewCourseItemService(conn),
		Course:      service.NewdCourseService(conn),
		// Dashboard:      service.NewDashboardService(conn),
		Gallery: service.NewGalleryService(conn),
		Teacher: service.NewTeacherService(conn),
	}, nil
}
