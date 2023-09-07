package request

type RequestInterface interface {
	KeepAlive(url string) bool
}
