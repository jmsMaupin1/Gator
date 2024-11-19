package rss

import (
	"context"
	"encoding/xml"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type Client struct {
	httpClient *http.Client	
}

func NewClient() Client {
	return Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "gator")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var feed RSSFeed
	decoder := xml.NewDecoder(res.Body)
	if err := decoder.Decode(&feed); err != nil {
		return nil, err
	}

	return &feed, nil
}
