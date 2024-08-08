package main

import (
	"context"
	"fmt"
	"log"
	"task_manager/data"
	connect "task_manager/db_connection"
	"task_manager/router"

)

func main(){
	fmt.Println("Server started")
	client, err := connect.DB()
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	defer func(){
		err := client.Disconnect(context.TODO())
		if err != nil{
			log.Fatal(err.Error())
		}
		}()

	tm, er := data.NewTaskmanager(client)
	if er != nil{
		log.Fatal(er.Error())
		return
	}
	router.Run(tm)

}