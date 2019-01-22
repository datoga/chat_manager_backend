package main

import (
	"log"
	"os"

	"github.com/datoga/chat_manager_backend/db"

	"github.com/datoga/chat_manager_backend/service"
)

func main() {

	connstr := "mongodb://chatmanager:1123581321chatmanager@ds159624.mlab.com:59624/chatmanager"

	dbchatmanager := db.NewMongoAccessor(connstr)

	log.Println("Connecting mongo")

	err := dbchatmanager.Connect()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	chatmanager := service.NewChatManagerService(dbchatmanager)

	chat, err := chatmanager.NewChat("prueba", "esta es una prueba")

	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	log.Println(chat)

	log.Println("Finish program")
}
