package handlers

import (
	"github.com/mirjalilova/websiteOfEverest/internal/clients"
	"github.com/mirjalilova/websiteOfEverest/pkg/minio"
)

type Handler struct {
	Clients *clients.GrpcClients
	MinIO   *minio.MinioClient
}

func NewHandler(client clients.GrpcClients, minio *minio.MinioClient) *Handler {
	return &Handler{
		Clients: &client,
		MinIO:   minio,
	}
}
