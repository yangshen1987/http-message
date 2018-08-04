package http_message

type StreamInterface interface {
	__toString()string
	Close()
	Detach()
	GetSize()
	Tell()
	Eof()
	IsSeekable()
	Seek(offset int, whence int)
	Rewind()
	IsWritable()
	Write(data string)
	IsReadable()
	Read(length int)
	GetContents()
	GetMetadata(key string)
}

