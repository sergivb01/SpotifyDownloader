package spotify

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type CLI struct {
	auth         *auth
	ClientID     string
	ClientSecret string
}

// Auth contains the basic authorization data
type auth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// NewClient creates an spotify client
func NewClient(id, secret string) *CLI {
	return &CLI{
		auth:         &auth{},
		ClientID:     id,
		ClientSecret: secret,
	}
}

func (c *CLI) authorize() error {
	httpc := &http.Client{}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpc.Transport = tr
	body := strings.NewReader(`grant_type=client_credentials`)
	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", body)
	req.Header.Add("cache-control", "no-cache")
	req.SetBasicAuth(c.ClientID, c.ClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := httpc.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(c.auth); err != nil {
		return err
	}

	return nil
}

func (c *CLI) request(method, url string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)

	err := c.authorize()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.auth.AccessToken)
	return http.DefaultClient.Do(req)
}

func (c *CLI) GetPlaylist(url string) (*Playlist, error) {
	res, err := c.request("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	playlist := &Playlist{}
	if err := json.NewDecoder(res.Body).Decode(playlist); err != nil {
		return nil, err
	}

	return playlist, nil
}
