package main
import (
	"fmt"
	"os/exec"

	"github.com/sergivb01/spotidown/spotify"
)

// DownloadSpotifyResource downloads a track from spotify
// by finding the track in youtube and downloading the first result
func DownloadSpotifyResource(r spotify.Resource, c chan string) {
	str := fmt.Sprintf("%s %s", r.Name, r.Artist)
	fmt.Printf("Now working on: %s...\n", str)

	sr, err := ytClient.FindVideos(str)
	if err != nil {
		fmt.Printf("Error found while trying to find videos for %s: %v", str, err)
	}
	url, err := sr.GetURL()
	if err != nil {
		fmt.Printf("%v\n", err)
		c <- str
		return
	}

	if err := downloadVideo(url); err != nil {
		fmt.Printf("Error while trying to download %s: %v\n", r.Name, err)
	}
	c <- str
}

// downloadVideo downloads a youtube video based on the url using youtube-dl
func downloadVideo(url string) error {
	fmt.Printf("Starting to download %s...\n", url)
	err := exec.Command("youtube-dl", "-f", "bestaudio",
		"-o", "music/%(title)s.%(ext)s",
		"--extract-audio",
		"--audio-format", "mp3",
		"--embed-thumbnail", "--add-metadata", url).Run()
	if err != nil {
		return err
	}
	return nil
}
