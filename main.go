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
		fmt.Println()
	}
}

func downloadSong(yt string, spot string) string {
	cmd := exec.Command("python", "downloadSong.py", yt, spot)

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}

	content, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Printf("failed to read file: %s", err)
	}
	songFileName := string(content)

	os.Remove("output.txt")

	return songFileName
}

func normalizeVolume(songFileName string) {
	// since mp3gain doesn't support unicode characters in file names, we rename the file to a temporary name - https://github.com/cfgnunes/wxmp3gain/issues/2
	os.Rename(songFileName, "songFileNameTemp.mp3")

	args := []string{"./mp3gain.exe", "/r", "/c", "songFileNameTemp.mp3"}
	cmd := exec.Command(args[0], args[1:]...)

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
	// fmt.Println("[mp3gain]", string(output))
	os.Rename("songFileNameTemp.mp3", songFileName)
}

func moveFile(file string, dir string) {
	os.Rename(file, dir+"/"+file)
}
