package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
)

func New(baseURL string, c *http.Client) *Client {
	return &Client{
		baseURL: baseURL,
		client:  c,
	}
}

type Client struct {
	baseURL string
	client  *http.Client
}

func (c *Client) Get(ctx context.Context, g game.ID, endpoint string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v/%v/%v", c.baseURL, g, endpoint), nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) RawApiGet(ctx context.Context, endpoint string, out interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v/%v", c.baseURL, endpoint), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(string(body))
	}
	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ApiGet(ctx context.Context, g game.ID, endpoint string, out interface{}) error {
	return c.RawApiGet(ctx, fmt.Sprintf("%v/%v", g, endpoint), out)
}
