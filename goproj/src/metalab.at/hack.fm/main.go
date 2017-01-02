package main

import (
	"metalab.at/hack.fm/inputs/pulse"
	"metalab.at/hack.fm/inputs/user"
	//"metalab.at/hack.fm/inputs/user/playtask"
	"time"
	//"os/exec"
	//"fmt"
	//"metalab.at/hack.fm/inputs/user/playtask/youtube"
)

func main() {

	//heartbeat pulseinput, keeps fm dongle alive
	f := pulseinput.NewPulseInput(time.Second * 35)
	f.PlugIn()
	
	
	//test code
	
	ui:= uinput.NewUserInput()
	
	ui.PlugIn()
	select{} //wait indefinitely
	println("DONE-----------------------------")


}
