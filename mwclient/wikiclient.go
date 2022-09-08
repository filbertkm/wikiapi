package mwclient

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	BaseURL string
	httpClient *http.Client
}

func NewClient(lang string) *Client {
	return &Client {
		BaseURL: fmt.Sprintf("https://%s.wikipedia.org/w/api.php", lang),
		httpClient: &http.Client {
			Timeout: time.Minute,
		},
	}
}

func (c *Client) GetPages(title string) (map[string]Page, error) {
	return c.getPageData(title)
}