package main

import (
	"fmt"
	"task_manager/data"
	"task_manager/models"
	"task_manager/router"
)

var task_manager data.Task_manager = data.Task_manager{
	Tasks : make([]*models.Task, 0),
	NextId: 0,
}

func main(){
	fmt.Println("Server started")
	router.Run(task_manager)
}