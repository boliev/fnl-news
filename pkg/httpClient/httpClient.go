package httpclient

// Client interface
type Client interface {
	Get(url string) (string, error)
}
