package http_message

type ResponseInterface interface {
	GetStatusCode()
	WithStatus()
	GetReasonPhrase()
}

