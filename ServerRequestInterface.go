package http_message

type ServerRequestInterface interface {
	GetServerParams()[]string
	GetCookieParams() []string
	WithCookieParams() ServerRequestInterface
	GetQueryParams() []string
	WithQueryParams(array []string)
	GetUploadedFiles()[]UriInterface
	WithUploadedFiles(uploadedFiles []interface{}) ServerRequestInterface
	GetParsedBody() map[string]string
	WithParsedBody() ServerRequestInterface
	GetAttributes()map[string]string
	GetAttribute(name string, defaultValue string) string
	WithAttribute()ServerRequestInterface
	WithoutAttribute() ServerRequestInterface
}