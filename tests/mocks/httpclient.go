package mocks

import (
	httpclient "github.com/boliev/fnl-news/pkg/httpClient"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	httpclient.Resty
	mock.Mock
}

func (c *ClientMock) Get(url string) (string, error) {
	args := c.Called(url)
	return args.String(0), args.Error(1)
}
