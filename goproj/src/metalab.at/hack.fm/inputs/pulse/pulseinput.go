package pulseinput


import(
	"time"
	"os/exec"
)

type PulseInput struct {
	tckr *time.Ticker
	dur time.Duration
	terminated chan int
}


func NewPulseInput(pulseTime time.Duration) *PulseInput{
	var pi PulseInput
	pi.tckr=nil
	pi.dur=pulseTime
	return &pi
}


func (pi *PulseInput) plugIn(){
	pi.tckr=time.NewTicker(pi.dur)
	
	quit:=false
	
	for(!quit){
		select{
		case <-pi.tckr.C:
			println("PLING")
			exec.Command("play", "../res/bell.mp3").Run();
			//bellRng.Run()
		case <-pi.terminated:
			quit=true
			println("TERMINATED");
		}
	}
	
}


func (pi *PulseInput) PlugIn(){
	go pi.plugIn()
}

func (pi *PulseInput) UnPlug(){
	pi.terminated<-0xFF
}
