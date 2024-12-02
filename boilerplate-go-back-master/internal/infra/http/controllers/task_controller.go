package controllers

import   "github.com/BohdanBoriak/boilerplate-go-back"

type TaskController struct {
    taskService app.TaskService
}

func NewTaskContr011er(ts app.TaskService) TaskController{
	return TaskController{
		taskService: ts,
	}
}
	