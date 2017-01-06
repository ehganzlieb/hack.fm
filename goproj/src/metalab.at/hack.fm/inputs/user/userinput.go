package uinput

import (
	"container/list"
	"container/ring"
	"fmt"
	"metalab.at/hack.fm/inputs/user/playtask"
	"metalab.at/hack.fm/inputs/user/playtask/youtube"
	"metalab.at/hack.fm/inputs/user/playtask/speech"
	"net/http"
	"time"
)

type UserInput struct {
	revolver      *ring.Ring
	queue         *list.List
	newlyCreated  bool
	playing       bool
	killed        bool
	hybridQueue   chan *playtask.Playtask
	revolverQueue chan *playtask.Playtask
	queueQueue    chan *playtask.Playtask
}

type dummyTask byte

type TaskContainer struct {
	task *playtask.Playtask
}

func wrap(task *playtask.Playtask) *TaskContainer {
	return &TaskContainer{
		task,
	}
}

func (tc *TaskContainer) exec() {
	x := playtask.Playtask(*tc.task)
	x.Exec()
}

func (dt *dummyTask) Exec() {
	time.Sleep(5 * time.Second)
}

func NewUserInput() *UserInput {
	ring := ring.New(1)
	dt := dummyTask(1)
	tsk := playtask.Playtask(&dt)
	ring.Value = wrap(&tsk)
	ui := UserInput{
		ring,
		list.New(),
		true,
		true,
		false,
		make(chan *playtask.Playtask, 255),
		make(chan *playtask.Playtask, 255),
		make(chan *playtask.Playtask, 255),
	}

	http.Handle("/", &ui)
	go http.ListenAndServe(":1337", nil)
	return &ui
}

func (ui *UserInput) addTaskToRevolver(task *playtask.Playtask) {
	if ui.newlyCreated {
		ui.newlyCreated = false
		ui.revolver.Value = wrap(task)
	} else {
		newRing := ring.New(1)
		newRing.Value = wrap(task)
		ui.revolver = ui.revolver.Link(newRing)
	}

}

func (ui *UserInput) addTaskToQueue(task *playtask.Playtask) {
	ui.queue.PushBack(wrap(task)) //add to end of list

}

func (ui *UserInput) addTaskHybrid(task *playtask.Playtask) {

	if ui.newlyCreated {
		ui.revolver.Value = wrap(task)
		ui.newlyCreated = false
	} else {
		newEl := ring.New(1)
		newEl.Value = wrap(task)
		ui.revolver.Prev().Link(newEl)
	}
	ui.addTaskToQueue(task)
}

func (ui *UserInput) plugIn() {

	for !ui.killed { //terminate if killed

		for ui.playing { //while playing basically this

			//empty all chans
			chanEmtyFinished := false
			for !chanEmtyFinished {
				//var tmpTask *playtask.Playtask
				select {
				case qtask := <-ui.queueQueue:
					ui.addTaskToQueue(qtask)
					println("adding to queue")
				case revTask := <-ui.revolverQueue:
					ui.addTaskToRevolver(revTask)
					println("adding to revolver")
				case hybTask := <-ui.hybridQueue:
					ui.addTaskHybrid(hybTask)
					println("adding hybrid")
				default:
					chanEmtyFinished = true
					//makes select non-blocking
				}
			}

			for !(ui.queue.Len() <= 0) { //empty queue
				println("playing from queue")
				/*fr := ui.queue.Front()
				tc := fr.Value.(*TaskContainer)
				ui.queue.Remove(fr
				*/
				tc := ui.queue.Remove(ui.queue.Front()).(*TaskContainer)

				tc.exec()
			} //queue is empty, so we execute a task from the revolver
			tc := (ui.revolver.Value).(*TaskContainer)
			println("playing from revolver")
			println(ui.revolver.Len())
			tc.exec()
			ui.revolver = ui.revolver.Next()
		}
		println("sleeping")
		time.Sleep(5 * time.Second)

	}
}

func (ui *UserInput) PlugIn() {
	go ui.plugIn()
}

func (ui *UserInput) UnPlug() {
	ui.playing = false
	ui.killed = true
}

func (ui *UserInput) PlayPause() {
	ui.playing = !ui.playing
}

func (ui *UserInput) AddTaskHybrid(task *playtask.Playtask) {
	ui.hybridQueue <- task
}

func (ui *UserInput) AddTaskQueue(task *playtask.Playtask) {
	ui.queueQueue <- task
}

func (ui *UserInput) AddTaskRevolver(task *playtask.Playtask) {
	ui.revolverQueue <- task
}

func (ui *UserInput) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	defer resp.Write([]byte(html)) //get this out of the way

	switch req.PostFormValue("mediatype") {
	case "yt":
		fmt.Println("YT ID post")
		switch req.PostFormValue("ytmode") {
		case "1": //hybrid
			if !(req.PostFormValue("ytlink") == "") {
				tsk := playtask.Playtask(youtubetask.NewYoutubeTask(req.PostFormValue("ytlink")))
				ui.AddTaskHybrid(&tsk)
			}
		case "2": //revolver
			if !(req.PostFormValue("ytlink") == "") {
				tsk := playtask.Playtask(youtubetask.NewYoutubeTask(req.PostFormValue("ytlink")))
				ui.AddTaskRevolver(&tsk)
			}
		case "3": //queue
			if !(req.PostFormValue("ytlink") == "") {
				tsk := playtask.Playtask(youtubetask.NewYoutubeTask(req.PostFormValue("ytlink")))
				ui.AddTaskQueue(&tsk)
			}
		}
	case "tts":
		if(req.PostFormValue("ttstxt")!=""){
			tsk:=playtask.Playtask((speechtask.NewSpeechTask(req.PostFormValue("ttstxt"))))
			ui.AddTaskQueue(&tsk);
		}
		
	}
}
