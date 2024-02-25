from spotdl import Spotdl
import sys

DEFAULT_CONFIG = {
    "client_id": "5f573c9620494bae87890c0f08a60293",
    "client_secret": "212476d9b0f3472eaa762d90b19b0ba8"
}

spotdl = Spotdl(client_id=DEFAULT_CONFIG["client_id"],
                client_secret=DEFAULT_CONFIG["client_secret"],
                )

youtubeURL = sys.argv[1]
spotifyURL = sys.argv[2]
search  = youtubeURL + "|" + spotifyURL

song = spotdl.search([search])
spotdl.download(song[0])