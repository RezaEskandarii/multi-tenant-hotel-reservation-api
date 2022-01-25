package common_services

import "os"

type FileDto struct {
	FileName   string
	Bytes      []byte
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
}
