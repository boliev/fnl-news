package httpclient

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/text/encoding/charmap"
)

// Resty1251 struct
type Resty1251 struct {
}

// NewResty1251 Resty with 1251 encoding support constructor
func NewResty1251() *Resty1251 {
	return &Resty1251{}
}

// Get sends Get Request
func (c Resty1251) Get(url string) (string, error) {
	client := resty.New()

	resp, err := client.R().
		Get(url)

	if err != nil {
		return "", err
	}
	res := resp.String()

	decoder := charmap.Windows1251.NewDecoder()
	res, err = decoder.String(res)
	if err != nil {
		return "", err
	}

	return res, nil
}

// Post sends Post Request
func (c Resty1251) Post(url string, body interface{}, headers map[string]string) error {
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
