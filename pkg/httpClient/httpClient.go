package httpclient

// Client interface
type Client interface {
	Get(url string) (string, error)
	Post(url string, body interface{}, headers map[string]string) error
}
