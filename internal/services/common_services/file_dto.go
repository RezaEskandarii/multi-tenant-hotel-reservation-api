package common_services

import "os"

type FileDto struct {
	BucketName string
	ServerName string
	FileSize   float64
	File       *os.File
}

type FileTransferResponse struct {
	Message    string `json:"message"`
	BucketName string `json:"bucket_name"`
	FileName   string `json:"file_name"`
	FileSize   int64  `json:"file_size"`
	VersionID  string `json:"version_id"`
}
