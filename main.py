import PySimpleGUI as sg
from spotdl import Spotdl

layout =    [[sg.Text('Enter Youtube URL:')],
            [sg.InputText(key='YT_URL')],
            [sg.Text('Enter Spotify URL:')],
            [sg.InputText(key='SPOT_URL')],
            [sg.Submit(), sg.Cancel()]]

window = sg.Window('Music Downloader', layout)

event, values = window.read()
window.close()

youtubeURL = values['YT_URL']
spotifyURL = values['SPOT_URL']

search  = youtubeURL + "|" + spotifyURL

spotdl = Spotdl(client_id='a', client_secret='a')
song = spotdl.search([search])
spotdl.download(song)