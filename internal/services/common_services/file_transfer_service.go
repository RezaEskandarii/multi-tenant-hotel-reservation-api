package common_services

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"os"
	"path/filepath"
	"reservation-api/internal/dto"
	"reservation-api/internal/utils/hash_utils"
	"sync"
	"time"
)

// FileTransformer interface is related to file management,
// which includes three upload and delete upload methods
type FileTransformer interface {
	Upload(bucketName, serverName string, file *os.File, wg *sync.WaitGroup) (*dto.FileTransferResponse, error)
	Remove(fileName, bucketName, versionID string) error
	Download(fileName, bucketName string) error
}

// FileTransferService implements FileTransformer interface
// this struct implements io functions with minio object manager.
type FileTransferService struct {
	Client *minio.Client
	Ctx    context.Context
}

// NewFileTransferService returns new instance of FileTransferService struct and gives minio client's config.
func NewFileTransferService(endpoint, accessKeyID, secretAccessKey string, useSSL bool, ctx context.Context) *FileTransferService {
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

// Upload uploads files via minion with FileDto input.
func (s *FileTransferService) Upload(bucketName, serverName string, file *os.File, wg *sync.WaitGroup) (*dto.FileTransferResponse, error) {

	if file == nil {
		wg.Done()
		return nil, errors.New("file is empty")
	}

	defer file.Close()

	// get file stat.
	fileStat, err := file.Stat()
	if err != nil {
		wg.Done()
		return nil, err
	}

	// check bucket exists.
	bucketExists, err := s.Client.BucketExists(s.Ctx, bucketName)
	if err != nil {
		return nil, err
	}
	// create bucket if not exists.
	if !bucketExists {
		s.Client.MakeBucket(s.Ctx, bucketName, minio.MakeBucketOptions{
			Region:        "",
			ObjectLocking: false,
		})
	}

	// generate random fileName
	fileName := s.generateRandomFileName(fileStat.Name())
	result, err := s.Client.PutObject(s.Ctx, bucketName, fileName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

	if err != nil {
		wg.Done()
		return nil, err
	}
	wg.Done()
	return &dto.FileTransferResponse{
		Message:    "Successfully uploaded.",
		BucketName: result.Bucket,
		FileName:   fileName,
		FileSize:   result.Size,
		VersionID:  result.VersionID,
	}, nil
}

// Remove removes file from bucket.
func (s *FileTransferService) Remove(bucketName, fileName, versionID string) error {
	// opts represents options specified by user for RemoveObject call
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

// Download finds object by given bucketName and fileName
// and streams founded object
func (s *FileTransferService) Download(bucketName, fileName string) error {

	obj, err := s.Client.GetObject(s.Ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	return s.stream(obj)
}

// stream streams given minio object.
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

// generates randomFileName
func (s *FileTransferService) generateRandomFileName(filename string) string {

	// get file extension
	fileExtension := filepath.Ext(filename)
	// generate random string
	randomStr := fmt.Sprintf("%s%s%d", filename, time.Now().String(), time.Now().UnixNano())
	// convert generated random string to SHA256 hash.
	randomStr = hash_utils.GenerateSHA256(randomStr)

	return fmt.Sprintf("%s%s%s", randomStr, uuid.New().String(), fileExtension)
}
