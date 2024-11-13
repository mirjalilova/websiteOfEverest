package service

import (
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	r "github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type CertificateService struct {
	storage r.Storage
	pb.UnimplementedCertificateServiceServer
}

func NewCertificateService(repo *r.Storage) *CertificateService {
	return &CertificateService{storage: *repo}
}
