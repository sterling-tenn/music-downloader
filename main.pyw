import PySimpleGUI as sg
from spotdl import Spotdl

DEFAULT_CONFIG = {
    "client_id": "5f573c9620494bae87890c0f08a60293",
    "client_secret": "212476d9b0f3472eaa762d90b19b0ba8",
}

spotdl = Spotdl(client_id=DEFAULT_CONFIG["client_id"],
                client_secret=DEFAULT_CONFIG["client_secret"],
                )

layout =    [[sg.Text('Enter Youtube URL:')],
            [sg.InputText(key='YT_URL')],
            [sg.Text('Enter Spotify URL:')],
            [sg.InputText(key='SPOT_URL')],
            [sg.Submit(), sg.Cancel()]]

window = sg.Window('Music Downloader', layout)

while True:
    event, values = window.read()
    if(event == 'Cancel' or event == sg.WIN_CLOSED):
        break
    
    youtubeURL = values['YT_URL']
    spotifyURL = values['SPOT_URL']

    if(event == 'Submit' and (not values['YT_URL'] or not values['SPOT_URL'])):
        sg.popup('Please specify a Youtube and Spotify URL')
    else:
        search  = youtubeURL + "|" + spotifyURL

        song = spotdl.search([search])
        spotdl.download(song[0])
        sg.popup("Finished.")

window.close()