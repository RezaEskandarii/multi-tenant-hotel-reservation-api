package common_services

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"path/filepath"
	"reservation-api/internal/utils"
	"time"
)

type IFileTransferService interface {
	Upload(fileInput *FileDto) (*FileTransferResponse, error)
	Remove(fileName string, bucketName string) error
	Download(fileName string, bucketName string) (error, []byte)
}

type FileTransferService struct {
	Client *minio.Client
	Ctx    context.Context
}

func (s *FileTransferService) New(endpoint, accessKeyID, secretAccessKey string, useSSL bool, ctx context.Context) *FileTransferService {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		panic(err.Error())
	}

	return &FileTransferService{
		Client: minioClient,
		Ctx:    ctx,
	}
}

func (s *FileTransferService) Upload(fileInput *FileDto) (*FileTransferResponse, error) {

	if fileInput.File == nil {
		return nil, errors.New("file is empty")
	}

	file := fileInput.File
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileName := s.generateRandomFileName(fileStat.Name())
	result, err := s.Client.PutObject(s.Ctx, fileInput.BucketName, fileName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

	if err != nil {
		return nil, err
	}

	return &FileTransferResponse{
		Message:    "",
		BucketName: result.Bucket,
		FileName:   fileName,
		FileSize:   result.Size,
	}, nil
}

func (s *FileTransferService) Remove(fileName string, bucketName string) error {
	panic("not implemented")
}

func (s *FileTransferService) Download(fileName string, bucketName string) (error, []byte) {
	panic("not implemented")
}

func (s *FileTransferService) generateRandomFileName(filename string) string {

	fileExtension := filepath.Ext(filename)
	randomStr := fmt.Sprintf("%s%s%s", filename, time.Now().String(), time.Now().UnixNano())
	randomStr = utils.GenerateSHA256(randomStr)

	return fmt.Sprintf("%s%s", randomStr, fileExtension)
}
