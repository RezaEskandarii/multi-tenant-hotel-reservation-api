package common_services

type FileDto struct {
	FileName   string
	Bytes      []byte
	BucketName string
	ServerName string
	FileSize   float64
}

type FileTransferResponse struct {
	File         FileDto
	Message      string
	ErrorMessage string
	Error        error
}
