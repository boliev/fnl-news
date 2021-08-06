package httpclient

import "github.com/go-resty/resty/v2"

// Resty struct
type Resty struct {
}

// NewResty Resty constructor
func NewResty() *Resty {
	return &Resty{}
}

// Get sends Get Request
func (c Resty) Get(url string) (string, error) {
	client := resty.New()

	resp, err := client.R().
		Get(url)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
