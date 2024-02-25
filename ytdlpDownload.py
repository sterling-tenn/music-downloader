import subprocess
import yt_dlp
import sys

def get_video_info(video_url):
    try:
        ydl = yt_dlp.YoutubeDL()
        info = ydl.extract_info(video_url, download=False)

        video_title = info.get('title', None)
        channel_name = info.get('uploader', None)
        
        return video_title, channel_name
    except yt_dlp.DownloadError as e:
        print("Error:", e)
        return None, None

def download_audio(video_url, filename):
    command = [
        "yt-dlp",
        "-x",  # Extract audio only
        "--audio-format", "mp3",
        "--audio-quality", "0",
        "--embed-thumbnail",
        "--add-metadata",
        "-o", filename,
        video_url
    ]

    try:
        subprocess.run(command, check=True)
    except subprocess.CalledProcessError as e:
        print("Error downloading audio:", e)

video_url = sys.argv[1]
video_title, channel_name = get_video_info(video_url)
filename = channel_name + " - " + video_title + ".mp3"
download_audio(video_url, filename)