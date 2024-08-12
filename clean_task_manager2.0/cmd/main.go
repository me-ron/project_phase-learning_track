package main

import (
	"context"
	"fmt"
	"log"
	"task_manager/database"
	"task_manager/delivery/routes"
)

func main(){
	fmt.Println("Server started")
	client, err := database.DB()
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

	db := client.Database("task_manager")
	routes.Run(db)

}