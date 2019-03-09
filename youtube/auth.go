package youtube

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type CLI struct {
	Token string
}

func NewClient(token string) *CLI {
	return &CLI{
		Token: token,
	}
}

func (c *CLI) request(method, url string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)

	return http.DefaultClient.Do(req)
}

func (c *CLI) FindVideos(query string) (*SearchResult, error) {
	vals := &url.Values{}
	vals.Add("q", query)
	vals.Add("key", c.Token)

	res, err := c.request("GET",
		"https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=1&type=video&"+vals.Encode(),
		nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// b, _ := ioutil.ReadAll(res.Body)
	// ioutil.WriteFile("cc.txt", b, 0644)

	video := &SearchResult{}
	if err := json.NewDecoder(res.Body).Decode(video); err != nil {
		return nil, err
	}

	return video, nil
}
