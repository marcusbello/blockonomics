package blockonomics

import (
	"net/http"
	"net/url"
)

type NewAddress struct {
	Address string `json:"address"`
	Crypto string `json:"crypto,omitempty"`
	Reset   int    `json:"reset,omitempty"`
	Account string `json:"account,omitempty"`
}

func (c *APIClient) NewAddress(account, crypto string, reset bool) (*NewAddress, error) {
	req, err := c.newRequest(http.MethodPost, "/api/new_address", nil)
	if err != nil {
		return nil, err
	}
	c.auth(req)

	params := url.Values{}
	if reset {
		params.Add("reset", "1")
	}
	if account != "" {
		params.Add("match_account", account)
	}
	if crypto != "" {
		params.Add("crypto", crypto)
	}
	req.URL.RawQuery = params.Encode()

	var newAddress NewAddress
	err = c.send(req, &newAddress)
	if err != nil {
		return nil, err
	}
	return &newAddress, nil
}
