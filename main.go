package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const DEST_DIR = "C:/Users/sterl/OneDrive/Music" // change this to your music directory

func main() {
	var youtubeURL string
	var spotifyURL string

	for {
		fmt.Print("Enter YouTube URL: ")
		fmt.Scanln(&youtubeURL)
		if youtubeURL == "" {
			fmt.Println("YouTube URL is required")
			continue
		}

		fmt.Print("Enter Spotify URL: ")
		fmt.Scanln(&spotifyURL)

		if spotifyURL == "" {
			fmt.Println("[yt-dlp download]")
			ytdlpDownload(youtubeURL)
		} else {
			fmt.Println("[spotdl download]")
			spotdlDownload(youtubeURL, spotifyURL)
		}

		songFileName, err := getLastModifiedMP3FileName()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Downloaded: " + songFileName)

		normalizeVolume(songFileName)
		fmt.Println("[mp3gain] done")

		moveFile(songFileName, DEST_DIR)
		fmt.Println("Moved to " + DEST_DIR + "/" + songFileName)
		fmt.Println()

		youtubeURL = ""
		spotifyURL = ""
	}
}

func getLastModifiedMP3FileName() (string, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return "", err
	}

	var lastModifiedTime time.Time
	var lastModifiedFileName string

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".mp3") {
			info, err := file.Info()
			if err != nil {
				return "", err
			}

			if info.ModTime().After(lastModifiedTime) {
				lastModifiedTime = info.ModTime()
				lastModifiedFileName = file.Name()
			}
		}
	}

	return lastModifiedFileName, nil
}

func ytdlpDownload(yt string) {
	cmd := exec.Command("python", "ytdlpDownload.py", yt)

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func spotdlDownload(yt string, spot string) {
	cmd := exec.Command("python", "spotdlDownload.py", yt, spot)

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func normalizeVolume(songFileName string) {
	// since mp3gain doesn't support unicode characters in file names, we rename the file to a temporary name - https://github.com/cfgnunes/wxmp3gain/issues/2
	os.Rename(songFileName, "songFileNameTemp.mp3")

	args := []string{"./mp3gain.exe", "/r", "/c", "songFileNameTemp.mp3"}
	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("[mp3gain]", string(output))
	os.Rename("songFileNameTemp.mp3", songFileName)
}

func moveFile(file string, dir string) {
	os.Rename(file, dir+"/"+file)
}
