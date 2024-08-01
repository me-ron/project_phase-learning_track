package main

import(
	"fmt"
	"task_manager/router"
)

func main(){
	fmt.Println("Server started")
	router.Run()
}