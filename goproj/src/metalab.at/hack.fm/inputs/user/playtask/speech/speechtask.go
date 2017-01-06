package speechtask


import(
	"os/exec"
	"fmt"
)


type SpeechTask struct{
	text string
}



func NewSpeechTask(text string) *SpeechTask{
	ret := new(SpeechTask)
	ret.text=text
	
	return ret
}



func (st *SpeechTask) Exec(){
	fmt.Println("Starting TTS", st.text)
	cmd:= exec.Command("espeak",st.text)
	fmt.Println(cmd)
	cmd.Run()
	fmt.Println("Finished TTS")
	
	
}



