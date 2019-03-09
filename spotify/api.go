package spotify

type Playlist struct {
	Items []struct {
		// AddedAt string `json:"added_at"`
		Track *Track `json:"track,omitempty"`
	} `json:"items"`
	Limit int64       `json:"limit"`
	Next  interface{} `json:"next"`
	Total int64       `json:"total"`
}

type Track struct {
	Name    string `json:"name"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
}

type SpotifyResource struct {
	Name   string `json:"name"`
	Artist string `json:"name"`
}

func (p *Playlist) Clear() []SpotifyResource {
	var resources []SpotifyResource
	for _, item := range p.Items {
		r := SpotifyResource{
			Name:   item.Track.Name,
			Artist: item.Track.Artists[0].Name,
		}
		resources = append(resources, r)
	}

	return resources
}
