package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var youtubeURL string
	var spotifyURL string

	for {
		fmt.Print("Enter YouTube URL: ")
		fmt.Scanln(&youtubeURL)

		fmt.Print("Enter Spotify URL: ")
		fmt.Scanln(&spotifyURL)

		songFileName := downloadSong(youtubeURL, spotifyURL)
		normalizeVolume(songFileName)
	}
}

func downloadSong(yt string, spot string) string {
	cmd := exec.Command("python", "downloadSong.py", yt, spot)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}

	songFileName := string(output)[:len(output)-2] // remove the newline character
	return songFileName
}

func normalizeVolume(songFileName string) {
	args := []string{"./mp3gain.exe", "/r", "/c", songFileName}
	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("[mp3gain]", string(output))
}