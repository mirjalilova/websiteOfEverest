package service

import (
	"context"
	"log/slog"

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

func (s *CertificateService) Create(ctx context.Context, req *pb.CreateCertificate) (*pb.Void, error) {
	res, err := s.storage.CertificateS.Create(req)
	if err != nil {
		slog.Error("Error creating certificate", "err", err)
		return nil, err
	}

	slog.Info("Created certificate")
	return res, nil
}

func (s *CertificateService) Update(ctx context.Context, req *pb.UpdateCertificate) (*pb.Void, error) {
	res, err := s.storage.CertificateS.Update(req)
	if err != nil {
		slog.Error("Error updating certificate", "err", err)
		return nil, err
	}

	slog.Info("Updated certificate")
	return res, nil
}

func (s *CertificateService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	res, err := s.storage.CertificateS.Delete(req)
	if err != nil {
		slog.Error("Error deleting certificate", "err", err)
		return nil, err
	}

	slog.Info("Deleted certificate")
	return res, nil
}

func (s *CertificateService) GetById(ctx context.Context, req *pb.ById) (*pb.CertificateRes, error) {
	res, err := s.storage.CertificateS.GetById(req)
	if err != nil {
		slog.Error("Error getting certificate by id", "err", err)
		return nil, err
	}

	slog.Info("Certificate retrieved")
	return res, nil
}

func (s *CertificateService) GetList(ctx context.Context, req *pb.GetListCertificateReq) (*pb.GetListCertificateRes, error) {
	res, err := s.storage.CertificateS.GetList(req)
	if err != nil {
		slog.Error("Error getting certificates list", "err", err)
		return nil, err
	}

	slog.Info("Listing certificates")
	return res, nil
}
