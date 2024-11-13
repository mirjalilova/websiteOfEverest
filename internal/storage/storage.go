package storage

import (
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
)

type BranchesI interface {
	CreateBranches(*pb.CreateBranches)(*pb.Void, error)
	UpdateBranches(*pb.UpdateBranches)(*pb.Void, error)
	DeleteBranches(*pb.ById) (*pb.Void, error)
	GetBranchesById(*pb.ById) (*pb.BranchesRes, error)
	GetBranchesList(*pb.GetListBranchesReq) (*pb.GetListBranchesRes, error)
}

type CertificateI interface {
	CreateCertificate(*pb.CreateCertificate)(*pb.Void, error)
	UpdateCertificate(*pb.UpdateCertificate)(*pb.Void, error)
	DeleteCertificate(*pb.ById) (*pb.Void, error)
	GetCertificateById(*pb.ById) (*pb.CertificateRes, error)
	GetCertificateList(*pb.GetListCertificateReq) (*pb.GetListCertificateRes, error)
}

type CourseItemI interface {
	CreateCourseItem(*pb.CreateCourseItem)(*pb.Void, error)
	UpdateCourseItem(*pb.UpdateCourseItem)(*pb.Void, error)
	DeleteCourseItem(*pb.ById) (*pb.Void, error)
	GetCourseItemById(*pb.ById) (*pb.CourseItemRes, error)
	GetCourseItemList(*pb.GetListCourseItemReq) (*pb.GetListCourseItemRes, error)
}

type CourseI interface {
	CreateCourse(*pb.CreateCourse)(*pb.Void, error)
	UpdateCourse(*pb.UpdateCourse)(*pb.Void, error)
	DeleteCourse(*pb.ById) (*pb.Void, error)
	GetCourseById(*pb.ById) (*pb.CourseRes, error)
	GetCourseList(*pb.GetListCourseReq) (*pb.GetListCourseRes, error)
}

type DashboardI interface{

}

type GalleryI interface {
	CreateGallery(*pb.CreateGallery)(*pb.Void, error)
	UpdateGallery(*pb.UpdateGallery)(*pb.Void, error)
	DeleteGallery(*pb.ById) (*pb.Void, error)
	GetGalleryById(*pb.ById) (*pb.GalleryRes, error)
	GetGalleryList(*pb.GetListGalleryReq) (*pb.GetListGalleryRes, error)
}

type TeacherI interface {
	CreateTeacher(*pb.CreateTeacher)(*pb.Void, error)
	UpdateTeacher(*pb.UpdateTeacher)(*pb.Void, error)
	DeleteTeacher(*pb.ById) (*pb.Void, error)
	GetTeacherById(*pb.ById) (*pb.TeacherRes, error)
	GetTeacherList(*pb.GetListTeacherReq) (*pb.GetListTeacherRes, error)
}
