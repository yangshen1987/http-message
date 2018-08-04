package http_message

type UploadedFileInterface interface {
	GetStream()
	MoveTo(targetPath string)
	GetSize()
	GetError()
	GetClientFilename()
	GetClientMediaType()
}
