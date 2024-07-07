# music-downloader
Work in progress script to help me automate my music collection process<br/>

Requirements:
* Python + [spotdl](https://github.com/spotDL/spotify-downloader), only tested with spotdl v4.1.11
	* `pip install -r requirements.txt`
* [ffmpeg](https://ffmpeg.org/)
	* `choco install ffmpeg` in elevated command prompt

Planned:
* ~~Pass through mp3 gain normalizer [MP3gain](http://mp3gain.sourceforge.net) (I currently use [wxmp3gain](https://github.com/cfgnunes/wxmp3gain) manually per song)~~ [Done]
	* Uses the CLI version of mp3gain: [mp3gain-dos-1_5_2.zip](https://sourceforge.net/projects/mp3gain/files/mp3gain/1.5.2/mp3gain-dos-1_5_2.zip/download)
* Add search functionality for easier link grabbing from both Youtube and Spotify
* General UI improvements
