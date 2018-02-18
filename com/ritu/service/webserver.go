package service

import (
	"log"
	"net/http"

	"github.com/rituK/com/ritu/controller"
	"github.com/rituK/com/ritu/dao"
)

func StartWebServer(port string) {

	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting HTTP service at " + port)
	intializeBoltConnection()
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())

	}

}

func intializeBoltConnection() {
	controller.DBClient = &dao.BoltClient{}
	controller.DBClient.OpenBoltDB()
	controller.DBClient.Seed()
}
