package http_message

type MessageInterFace interface {
	GetProtocolVersion() string
	WithProtocolVersion() string
	GetHeaders() [][]string
	HasHeader() bool
	GetHeader() []string
	GetHeaderLine() string
	WithHeader() MessageInterFace
	WithAddedHeader() MessageInterFace
	WithoutHeader() MessageInterFace
	GetBody() StreamInterface
	WithBody() MessageInterFace
}
