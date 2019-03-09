package spotify
// Playlist defines the basic structure of a Spotify playlist
type Playlist struct {
	Items []struct {
		// AddedAt string `json:"added_at"`
		Track *Track `json:"track,omitempty"`
	} `json:"items"`
	Limit int64       `json:"limit"`
	Next  interface{} `json:"next"`
	Total int64       `json:"total"`
}

// Track defines basic structure of Spotify track
type Track struct {
	Name    string `json:"name"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
}

// Resource contains the most basic information about a track
type Resource struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

// Clear transforms the playlist into a SpotifyResource slice
func (p *Playlist) Clear() []Resource {
	var resources []Resource
	for _, item := range p.Items {
		r := Resource{
			Name:   item.Track.Name,
			Artist: item.Track.Artists[0].Name,
		}
		resources = append(resources, r)
	}

	return resources
}
