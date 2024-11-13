package minio

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/mirjalilova/websiteOfEverest/config"
)

type MinioClient struct {
	Client     *minio.Client
	BucketName string
}

func NewMinioClient(cnf *config.Config) (*MinioClient, error) {
	minioClient, err := minio.New(cnf.MINIO_ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(cnf.MINIO_ACCESS_KEY, cnf.MINIO_SECRET_KEY, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %v", err)
	}

	client := &MinioClient{
		Client:     minioClient,
		BucketName: cnf.MINIO_BUCKET_NAME,
	}

	// Create bucket if it doesn't exist
	err = client.createBucket()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (mc *MinioClient) createBucket() error {
	exists, err := mc.Client.BucketExists(context.Background(), mc.BucketName)
	if err != nil {
		return fmt.Errorf("failed to check if bucket exists: %v", err)
	}
	if !exists {
		err = mc.Client.MakeBucket(context.Background(), mc.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}
		log.Printf("Bucket %s created successfully", mc.BucketName)
	}
	return nil
}

func (mc *MinioClient) UploadFile(ctx context.Context, objectName string, filePath string) (string, error) {
	_, err := mc.Client.FPutObject(ctx, mc.BucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}
	return mc.Client.EndpointURL().String() + "/" + mc.BucketName + "/" + objectName, nil
}
