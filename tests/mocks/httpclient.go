package mocks

import (
	httpclient "github.com/boliev/fnl-news/pkg/httpClient"
	"github.com/stretchr/testify/mock"
)

// ClientMock mock for http client
type ClientMock struct {
	httpclient.Resty
	mock.Mock
}

// Get mock function for get
func (c *ClientMock) Get(url string) (string, error) {
	args := c.Called(url)
	return args.String(0), args.Error(1)
}
