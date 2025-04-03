package services

import (
	"context"
	"errors"
	"io"
	"mime"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

type MinIOUploadService struct {
	client        *minio.Client
	defaultBucket string
}

func NewMinIOUploadService(client *minio.Client, defaultBucket string) *MinIOUploadService {
	if client == nil {
		panic("client MinIO não pode ser nil")
	}
	return &MinIOUploadService{
		client:        client,
		defaultBucket: defaultBucket,
	}
}

type UploadOptions struct {
	BucketName  string
	ObjectName  string
	ContentType string
}

func (s *MinIOUploadService) UploadFile(
	ctx context.Context,
	file io.Reader,
	fileSize int64,
	options UploadOptions,
) (minio.UploadInfo, error) {
	// Proteção contra nil pointer
	if s == nil {
		return minio.UploadInfo{}, errors.New("serviço MinIO não inicializado")
	}

	// Validações básicas
	if file == nil {
		return minio.UploadInfo{}, errors.New("arquivo não pode ser nil")
	}
	if options.ObjectName == "" {
		return minio.UploadInfo{}, errors.New("nome do objeto é obrigatório")
	}

	// Usar bucket padrão se não especificado
	bucketName := options.BucketName
	if bucketName == "" {
		if s.defaultBucket == "" {
			return minio.UploadInfo{}, errors.New("nome do bucket é obrigatório")
		}
		bucketName = s.defaultBucket
	}

	// Determinar content-type se não especificado
	contentType := options.ContentType
	if contentType == "" {
		ext := filepath.Ext(options.ObjectName)
		contentType = mime.TypeByExtension(ext)
		if contentType == "" {
			contentType = "application/octet-stream"
		}
	}

	// Fazer upload
	return s.client.PutObject(
		ctx,
		bucketName,
		options.ObjectName,
		file,
		fileSize,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
}
