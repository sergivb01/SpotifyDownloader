package main
import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/sergivb01/spotidown/spotify"
	"github.com/sergivb01/spotidown/youtube"
)

var (
	ytClient youtube.CLI
	spClient spotify.CLI

	playlistID = flag.String("playlist", "3RnyAFc44FNIMe2ACUJTxF", "The ID of the playlist")
)

func main() {
	spClient = *spotify.NewClient(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET"))
	ytClient = *youtube.NewClient(os.Getenv("YOUTUBE_KEY"))

	spotifyPlaylist, err := spClient.GetPlaylist("https://api.spotify.com/v1/playlists/" + *playlistID + "/tracks")
	if err != nil {
		fmt.Printf("%v", err)
	}

	playlist := spotifyPlaylist.Clear()

	c := make(chan string, len(playlist))
	defer close(c)

	for _, res := range playlist {
		go DownloadSpotifyResource(res, c)
		time.Sleep(time.Second)
	}

	for i := 1; i <= len(playlist); i++ {
		msg := <-c
		prc := i * 100 / len(playlist)
		fmt.Printf("[%d/%d - %d%%] Finished downloading %s\n", i, len(playlist), prc, msg)
	}

}
