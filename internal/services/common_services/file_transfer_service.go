package common_services

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"path/filepath"
	"reservation-api/internal/utils"
	"time"
)

type IFileTransferService interface {
	Upload(fileInput *FileDto) (*FileTransferResponse, error)
	Remove(fileName, bucketName, versionID string) error
	Download(fileName, bucketName string) error
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

	bucketExists, err := s.Client.BucketExists(s.Ctx, fileInput.BucketName)
	if err != nil {
		return nil, err
	}

	if !bucketExists {
		s.Client.MakeBucket(s.Ctx, fileInput.BucketName, minio.MakeBucketOptions{
			Region:        "",
			ObjectLocking: false,
		})
	}

	fileName := s.generateRandomFileName(fileStat.Name())
	result, err := s.Client.PutObject(s.Ctx, fileInput.BucketName, fileName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

	if err != nil {
		return nil, err
	}

	return &FileTransferResponse{
		Message:    "Successfully uploaded.",
		BucketName: result.Bucket,
		FileName:   fileName,
		FileSize:   result.Size,
		VersionID:  result.VersionID,
	}, nil
}

func (s *FileTransferService) Remove(bucketName, fileName, versionID string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
		VersionID:        versionID,
	}
	err := s.Client.RemoveObject(s.Ctx, bucketName, fileName, opts)
	if err != nil {
		return err
	}
	return nil
}

func (s *FileTransferService) Download(bucketName, fileName string) error {

	obj, err := s.Client.GetObject(s.Ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	return s.stream(obj)
}

func (s *FileTransferService) stream(r io.Reader) error {
	br := bufio.NewReader(r)
	b := make([]byte, 10000, 10000)
	for {
		_, err := br.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	return nil
}

func (s *FileTransferService) generateRandomFileName(filename string) string {
	fileExtension := filepath.Ext(filename)
	randomStr := fmt.Sprintf("%s%s%s", filename, time.Now().String(), time.Now().UnixNano())
	randomStr = utils.GenerateSHA256(randomStr)
	return fmt.Sprintf("%s%s", randomStr, fileExtension)
}
