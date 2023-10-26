package main

import (
	"AIAssistServer/controller"
	"AIAssistServer/db"
	"AIAssistServer/router"
	"fmt"
)

func main() {
	db.InitMysql()
	controller.NewClient()

	r := router.Router()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
