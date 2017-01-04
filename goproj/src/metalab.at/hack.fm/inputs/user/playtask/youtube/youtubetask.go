package youtubetask


import(
	"os/exec"
	"fmt"
)


type YoutubeTask struct{
	videoId string
}


func NewYoutubeTask(videoId string) *YoutubeTask{
	yt := YoutubeTask{
		videoId,
	}
	return &yt
}


func (tsk YoutubeTask) Exec(){
	println("running yt")
	//dlink:= "https://youtu.be/"+tsk.videoId
	dlink:= tsk.videoId
	println(dlink)
	
	//cmd:= exec.Command("util/youtube.sh", dlink)
	cmd:= exec.Command("youtube-dl", "-x", "--exec", "mpv {}", dlink)
	fmt.Println(cmd)
	cmd.Run()
	//exec.Command("sh", "-c", "'youtube-dl -x --exec 'echo {}>fname' "+dlink+"'" ).Run()
	//exec.Command("sh", "-c", "mpv \"`cat fname`\"").Run()
	
	println("finished yt");
	
}
