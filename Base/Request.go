package Base

import "http-message"

type Request struct {
	
}

var server []string

var uri  http_message.UriInterface

var method string

var requestTarget string

func GetRequestTarget() string  {
	if requestTarget != "" {
		return requestTarget
	}
	target := uri.GetPath()
}
