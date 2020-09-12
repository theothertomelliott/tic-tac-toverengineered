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

func (c *Client) ApiGet(g game.ID, endpoint string, out interface{}) error {
	resp, err := http.Get(fmt.Sprintf("%v/%v/%v", c.baseURL, g, endpoint))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}
	return nil
}
