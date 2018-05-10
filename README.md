# scon-mpd

A Simple Controller Command Line for MPD

Basic Syntax:

	scon-mpd [command]

Commands:

	"play"		::	play current song
	"pause" 	::	pause,or un-pause
	"playpause"	::	if mpd is stop, play; if pause or playing, toggles pause
				<use with combined play/pause media key>
	"stop"		::	stop mpd
	"next" 		::	play next song
	"previous" 	::	play previous song

Note: Above basic commands produce NO output, purpose is for media key bindings.

CLI Commands:
	
	"status"	::	print status of toggles/controls


