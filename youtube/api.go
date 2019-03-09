package youtube
import "fmt"

// SearchResult defines a result from a video search
type SearchResult struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
	} `json:"items"`
}

// GetURL returns the video URL of a search result
func (sr *SearchResult) GetURL() (string, error) {
	if sr == nil {
		return "", fmt.Errorf("Error while trying to get a URL - there is no result: %v", sr)
	}
	return "http://youtube.com/watch?v=" + sr.Items[0].ID.VideoID, nil
}
