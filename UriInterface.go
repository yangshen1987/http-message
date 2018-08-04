package http_message

type UriInterface interface {
	GetScheme()
	GetAuthority()
	GetUserInfo()
	GetHost()
	GetPort()
	GetPath() string
	GetQuery()
	GetFragment()
	WithScheme()
	WithUserInfo()
	WithHost()
	WithPort()
	WithPath()
	WithQuery()
	WithFragment()
	__ToString()
}

