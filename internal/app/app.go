package app

import (
	"log"

	"github.com/mirjalilova/websiteOfEverest/api/handlers"
	"github.com/mirjalilova/websiteOfEverest/config"
	"github.com/mirjalilova/websiteOfEverest/internal/clients"
	"github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
	"github.com/mirjalilova/websiteOfEverest/pkg/db"
	"github.com/mirjalilova/websiteOfEverest/pkg/logger"
	"github.com/mirjalilova/websiteOfEverest/pkg/minio"
	"github.com/mirjalilova/websiteOfEverest/api"
)

func Run(cfg *config.Config) {
	infoLog, errLog := logger.InitLogger()

	// Postgres Connection
	db, err := db.Connect(cfg)
	if err != nil {
		errLog.Println("Can't connect to database, details:", err, log.Ldate|log.Ltime|log.Lshortfile)
	}
	defer db.Close()
	infoLog.Println("Connected to Postgres")

	conn, err := repository.NewPostgresStorage(db)

	//MinIO
	minioClient, err := minio.NewMinioClient(cfg)
	if err != nil {
		errLog.Println("Failed to connect to MinIO", err)
		return
	}
	infoLog.Println("Connected to MinIO")

	// clients
	clients, err := clients.NewGrpcClients(cfg, conn)
	if err != nil {
		errLog.Printf("Error while connecting clients. err: %s", err)
	}
	infoLog.Println("Connected to gRPC clients")

	handlers := handlers.NewHandler(*clients, minioClient) 

	roter := api.NewApi(handlers)

	if err := roter.Run(cfg.ServerPort); err != nil {
		errLog.Fatal("Failed to start server", err)
        return
	}
	infoLog.Println("Server started on port", cfg.ServerPort)
}
