package storage

import (
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
)

type BranchesI interface {
	Create(*pb.CreateBranches)(*pb.Void, error)
	Update(*pb.UpdateBranches)(*pb.Void, error)
	Delete(*pb.ById) (*pb.Void, error)
	GetById(*pb.ById) (*pb.BranchesRes, error)
	GetList(*pb.GetListBranchesReq) (*pb.GetListBranchesRes, error)
}

type CertificateI interface {
	Create(*pb.CreateCertificate)(*pb.Void, error)
	Update(*pb.UpdateCertificate)(*pb.Void, error)
	Delete(*pb.ById) (*pb.Void, error)
	GetById(*pb.ById) (*pb.CertificateRes, error)
	GetList(*pb.GetListCertificateReq) (*pb.GetListCertificateRes, error)
}

type CourseItemI interface {
	Create(*pb.CreateCourseItem)(*pb.Void, error)
	Update(*pb.UpdateCourseItem)(*pb.Void, error)
	Delete(*pb.ById) (*pb.Void, error)
	GetById(*pb.ById) (*pb.CourseItemRes, error)
	GetList(*pb.GetListCourseItemReq) (*pb.GetListCourseItemRes, error)
}

type CourseI interface {
	Create(*pb.CreateCourse)(*pb.Void, error)
	Update(*pb.UpdateCourse)(*pb.Void, error)
	Delete(*pb.ById) (*pb.Void, error)
	GetById(*pb.ById) (*pb.CourseRes, error)
	GetList(*pb.GetListCourseReq) (*pb.GetListCourseRes, error)
}

type DashboardI interface{

}

type GalleryI interface {
	Create(*pb.CreateGallery)(*pb.Void, error)
	Update(*pb.UpdateGallery)(*pb.Void, error)
	Delete(*pb.ById) (*pb.Void, error)
	GetById(*pb.ById) (*pb.GalleryRes, error)
	GetList(*pb.GetListGalleryReq) (*pb.GetListGalleryRes, error)
}

type TeacherI interface {
	Create(*pb.CreateTeacher)(*pb.Void, error)
	Update(*pb.UpdateTeacher)(*pb.Void, error)
	Delete(*pb.ById) (*pb.Void, error)
	GetById(*pb.ById) (*pb.TeacherRes, error)
	GetList(*pb.GetListTeacherReq) (*pb.GetListTeacherRes, error)
}
