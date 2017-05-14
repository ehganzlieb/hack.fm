package main

import (
	"metalab.at/hack.fm/inputs/pulse"
	"metalab.at/hack.fm/inputs/user"
	//"metalab.at/hack.fm/inputs/user/playtask"
	"time"
	"metalab.at/hack.fm/inputs/cleanup"
	//"os/exec"
	//"fmt"
	//"metalab.at/hack.fm/inputs/user/playtask/youtube"
)

func main() {

	//heartbeat pulseinput, keeps fm dongle alive
	f := pulseinput.NewPulseInput(time.Second * 35)
	f.PlugIn()


	//test code
	defer println("DONE-----------------------------")
	ui:= uinput.NewUserInput()

	ui.PlugIn()

	ci := cinput.NewCleanupInput(ui, 21,45) //set time as needed
	ci.PlugIn()

	select{}
		 //wait indefinitely


}
