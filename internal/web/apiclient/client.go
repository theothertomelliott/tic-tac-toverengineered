package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
)

func New(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
	}
}

type Client struct {
	baseURL string
}

func (c *Client) Get(g game.ID, endpoint string) (*http.Response, error) {
	return http.Get(fmt.Sprintf("%v/%v/%v", c.baseURL, g, endpoint))
}

func (c *Client) RawApiGet(endpoint string, out interface{}) error {
	resp, err := http.Get(fmt.Sprintf("%v/%v", c.baseURL, endpoint))
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

func (c *Client) ApiGet(g game.ID, endpoint string, out interface{}) error {
	return c.RawApiGet(fmt.Sprintf("%v/%v", g, endpoint), out)
}
