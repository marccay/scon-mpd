# scon-mpd

A Simple Controller Command Line for MPD

Basic Syntax:

	scon-mpd [command]

Commands:

	"play"				play current song
	"pause" 			pause,or un-pause
	"playpause"			if mpd is stop, play; if pause or playing, toggles pause
						<use with combined play/pause media key>
	"stop"				stop mpd
	"next" 				play next song
	"previous" 			play previous song
	: "smart" | "s"				smart previous, rewind current song first

Note: Above basic commands produce NO output, purpose is for media key bindings.

CLI Commands:
	
	scon-mpd [command] [option]

		"currentsong" | "current		print info on current song

		"repeat"				toggle repeat
		: "on" | "off"

		"random"				toggle random
		: "on" | "off"

		"single"				toggle single mode (repeat song)
		: "on" | "off"
	
		"consume"				toggle consume mode (remove from playlist after play)
		: "on" | "off"

		"crossfade" | "fade" | "xfade"		toggle crossfade (ON or default 10 seconds)
		: "off" | [seconds]				[seconds] = number of seconds



		"status"				print status of toggles/controls
		: "state"					play state: play/pause/stopped
		: "repeat"					repeat toggle: on/off
		: "random" | "rand"				random toggle: on/off
		: "single"					single repeat: on/off
		: "consume"					consume song from playlist after play: on/off
		: "crossfade" | "fade" | "xfade"		crossfade: on/off
		: "time"					prints elapsed/duration
		: "elapsed"					elapsed time of song in minutes
		: "duration"					total duration of song in minutes
	
