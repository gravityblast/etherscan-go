package etherscan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type NetworkID = int

const (
	Mainnet = 1
	Rinkeby = 3

	baseUrlMainnet = "https://api.etherscan.io/api"
	baseUrlRinkeby = "https://api-rinkeby.etherscan.io/api"
)

type Client struct {
	c       httpClient
	baseURL string
	apiKey  string
}

func NewClient(networkID NetworkID, apiKey string) (*Client, error) {
	url, err := urlByNetworkID(networkID)
	if err != nil {
		return nil, err
	}

	return &Client{
		c:       http.DefaultClient,
		baseURL: url,
		apiKey:  apiKey,
	}, nil
}

func (c *Client) Account(address string) (*AccountResponse, error) {
	params := map[string]string{
		"module":  "account",
		"action":  "balance",
		"address": address,
		"tag":     "latest",
	}

	resp, err := c.get(params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var accountResp AccountResponse

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&accountResp)

	return &accountResp, err
}

func (c *Client) get(params map[string]string) (*http.Response, error) {
	query := c.buildQuery(params)
	url, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}

	url.RawQuery = query
	resp, err := c.c.Get(url.String())
	if err != nil {
		return resp, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		resp.Body.Close()
		return resp, fmt.Errorf("bad status code %d", resp.StatusCode)
	}

	return resp, nil
}

func (c *Client) buildQuery(params map[string]string) string {
	v := url.Values{}
	for key, value := range params {
		v.Set(key, value)
	}

	v.Set("apikey", c.apiKey)

	return v.Encode()
}
