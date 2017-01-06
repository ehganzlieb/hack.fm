package cinput


import(
	"os/exec"
	"time"
	"metalab.at/hack.fm/inputs/user"
	"metalab.at/hack.fm/inputs/user/playtask/speech"
	"metalab.at/hack.fm/inputs/user/playtask"
)

type CleanupInput struct{
	ui *uinput.UserInput
	hours int
	mins int
	pluggedIn bool
}


func NewCleanupInput(ui *uinput.UserInput,hour int, min int) *CleanupInput{
	ret := new(CleanupInput)
	ret.ui=ui
	ret.hours=hour
	ret.mins=min
	ret.pluggedIn=true
	return ret
}


func (ci *CleanupInput) plugIn(){
	ci.pluggedIn=true
	for ci.pluggedIn{
		if time.Now().Hour()==ci.hours&&time.Now().Minute()==ci.mins{
			ci.cleanupAnnouncement()
			time.Sleep(5*time.Second)
			for ci.pluggedIn{
				<- time.After(time.Hour*24)
				ci.cleanupAnnouncement()
			}
			
		}
	}
}

func (ci *CleanupInput) PlugIn(){
	go ci.plugIn()
}

func (ci *CleanupInput) UnPlug(){
	ci.pluggedIn=false
}



func (ci *CleanupInput)cleanupAnnouncement(){
	tsk:=playtask.Playtask(speechtask.NewSpeechTask("Now would be a good time for a cleanup"))
	exec.Command("killall","mpv").Run()
	exec.Command("killall", "espeak").Run()
	ci.ui.AddTaskQueue(&tsk)
}
