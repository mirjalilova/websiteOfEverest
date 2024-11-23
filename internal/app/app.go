package app

import (
	"github.com/mirjalilova/websiteOfEverest/api"
	"github.com/mirjalilova/websiteOfEverest/api/handlers"
	"github.com/mirjalilova/websiteOfEverest/config"
	"github.com/mirjalilova/websiteOfEverest/internal/clients"
	"github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
	"github.com/mirjalilova/websiteOfEverest/pkg/db"
	"github.com/mirjalilova/websiteOfEverest/pkg/logger"
	"github.com/mirjalilova/websiteOfEverest/pkg/minio"
)

func Run(cfg *config.Config) {
	infoLog, errLog := logger.InitLogger()

	// Postgres Connection
	db, err := db.Connect(cfg)
	if err != nil {
		errLog.Println("Can't connect to database, details:", err)
		return
	}
	defer db.Close() // Ensures the DB connection is closed when the function exits
	infoLog.Println("Connected to Postgres")

	conn, err := repository.NewPostgresStorage(db)
	if err != nil {
		errLog.Println("Failed to connect to Postgres:", err)
		return
	}
	infoLog.Println("Connected to Postgres")

	// MinIO Connection
	minioClient, err := minio.NewMinioClient(cfg)
	if err != nil {
		errLog.Println("Failed to connect to MinIO:", err)
		return
	}
	infoLog.Println("Connected to MinIO")

	// gRPC Clients
	clients, err := clients.NewGrpcClients(cfg, conn)
	if err != nil {
		errLog.Printf("Error while connecting to gRPC clients: %s", err)
		return
	}
	infoLog.Println("Connected to gRPC clients")

	// Handlers and API setup
	handlers := handlers.NewHandler(*clients, minioClient)
	roter := api.NewApi(handlers)

	// Run server
	if err := roter.Run(cfg.ServerPort); err != nil {
		errLog.Fatal("Failed to start server:", err)
		return
	}
	infoLog.Println("Server started on port", cfg.ServerPort)
}
