package http_message

type RequestInterface interface {
	GetRequestTarget() string
	WithRequestTarget() RequestInterface
	GetMethod() string
	WithMethod() RequestInterface
	GetUri() UriInterface
	WithUri() RequestInterface
}
