package httpclient

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

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

// Post sends Post Request
func (c Resty) Post(url string, body interface{}, headers map[string]string) error {
	client := resty.New()

	res, err := client.R().
		SetBody(body).
		SetHeaders(headers).
		Post(url)
	if res != nil && res.StatusCode() > 299 {
		return fmt.Errorf("code: %d, response: %s. request: %s", res.StatusCode(), res.String(), res.Request.Body)
	}

	return err
}
