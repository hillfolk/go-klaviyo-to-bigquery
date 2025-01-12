package client

import (
	"github.com/go-resty/resty/v2"
)

const BASE_API_URL = "https://a.klaviyo.com/api"

// TQcTdC
// pk_1803af864063f7e05c2cf037e9a332ca6d

/*
	curl --request GET \
	     --url https://a.klaviyo.com/api/{endpoint}/ \
	     --header 'Authorization: Klaviyo-API-Key your-private-api-key' \
	     --header 'accept: application/json' \
	     --header 'revision: 2024-10-15'
*/
type Client struct {
	*resty.Client
	key      string
	revision string
	query    *query
}

// 1. Klaviyo API 를 호출하기 위한 Client 를 생성합니다.
func NewClient(key string) *Client {
	c := resty.New()
	c.SetAuthScheme("Klaviyo-API-Key")
	c.SetAuthToken(key)
	c.SetBaseURL(BASE_API_URL)
	c.SetHeader("accept", "application/json")
	c.SetHeader("revision", "2024-10-15")

	return &Client{
		Client: c,
		key:    key,
		query:  newQuery(),
	}
}

func (c *Client) ClearQuery() {
	c.query = newQuery()
}

func (c *Client) GetQuery() *query {
	return c.query
}

func (c *Client) SetRawQuery(q string) {
	c.SetRawQuery(q)
}
