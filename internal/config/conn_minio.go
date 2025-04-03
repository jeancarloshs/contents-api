package config

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	DefaultBucket   string
}

func ConnectMinio() (*minio.Client, error) {
	endpoint := "localhost:9000"
	accessKeyID := "dmUOyomwRl7Q8HfcCrHp"
	secretAccessKey := "ycHEjA9rtBNCDixUsd308AWh7ctQBtEN2byo3y71"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Erro ao conectar ao MinIO: %v", err)
		return nil, err
	}

	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		log.Fatalf("Erro ao listar buckets: %v", err)
		return nil, err
	}
	fmt.Println("Conexão estabelecida. Total de buckets:", len(buckets))
	return minioClient, nil
}
