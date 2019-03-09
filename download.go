package main

import (
	"fmt"
	"os/exec"

	"github.com/sergivb01/spotidown/spotify"
)

func DownloadSpotifyResource(r spotify.SpotifyResource, c chan string) {
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

func downloadVideo(url string) error {
	fmt.Printf("Starting to download %s...\n", url)
	_, err := exec.Command("youtube-dl", "-f", "bestaudio",
		"-o", "music/%(title)s.%(ext)s",
		"--extract-audio",
		"--audio-format", "mp3",
		"--embed-thumbnail", "--add-metadata", url).CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
