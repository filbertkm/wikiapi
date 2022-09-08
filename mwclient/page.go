package mwclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Page struct {
	PageID int `json:"pageid"`
	Ns int `json:"ns"`
	Title string `json:"title"`
	Description string `json:"description"`
	Source string `json:"descriptionsource"`
	Missing bool `json:"missing,omitempty"`
}

type pageResponse struct {
	Batchcomplete string `json:"batchcomplete"`
	Query struct {
        Pages map[string]Page `json:"pages"`
    } `json:"query"`
}

func (c *Client) getPageData(title string) (map[string]Page, error) {
	url := fmt.Sprintf("%s?action=query&prop=description&descprefersource=local&titles=%s&format=json", c.BaseURL, title)

	pages := make(map[string]Page)
	r := new(pageResponse)

	res, err := c.httpClient.Get(url)
	if err != nil {
		return pages, err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return pages, err
	}

	if err := json.Unmarshal(resBody, r); err != nil {
		return pages, err
	}

	return r.Query.Pages, nil
}