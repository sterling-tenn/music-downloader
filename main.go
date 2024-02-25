package main

import (
	"fmt"
	"os"
	"os/exec"
)

const DEST_DIR = "C:/Users/sterl/OneDrive/Music" // change this to your music directory

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
		moveFile(songFileName, DEST_DIR)
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

func moveFile(file string, dir string) {
	os.Rename(file, dir+"/"+file)
}
