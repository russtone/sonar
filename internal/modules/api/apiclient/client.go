package apiclient

import (
	"crypto/tls"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/bi-zone/sonar/internal/actions"
)

type Client struct {
	client *resty.Client
}

var _ actions.Actions = &Client{}

func New(url string, token string, insecure bool) *Client {
	c := resty.New().
		SetHostURL(url).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetHeader("Content-Type", "application/json")

	if insecure {
		c.SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	}

	return &Client{c}
}
