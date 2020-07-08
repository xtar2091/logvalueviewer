package main

type TaskList struct {
	Done chan struct{}
	WorkList chan string
}

var task = &TaskList{
	Done: make(chan struct{}),
	WorkList: make(chan string),
}

var taskListMaxSize = 10000

func (obj *TaskList) CloseWorkList() {
	close(obj.WorkList)
}

func (obj *TaskList) Push(line string) bool {
	ok := false
	if len(obj.WorkList) < taskListMaxSize {
		obj.WorkList <- line
		ok = true
	}
	return ok
}

func (obj *TaskList) Stop() {
	close(obj.Done)
}

func (obj *TaskList) Wait() {
	for range obj.Done {

	}
}
