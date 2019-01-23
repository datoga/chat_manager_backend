package main

import (
	"log"
	"os"

	"github.com/datoga/chat_manager_backend/model"

	"github.com/datoga/chat_manager_backend/db"

	"github.com/datoga/chat_manager_backend/service"
)

const mongoConnStr = "mongodb://chatstorer:1123581321chatstore@ds111065.mlab.com:11065/chatstore"
const redisConnStr = "redis-17420.c60.us-west-1-2.ec2.cloud.redislabs.com:17420"
const redisPass = "KEvnu3QoRJguUqPyeWyNYApt5ZTfgsGa"
const redisDB = 0

func main() {
	dbchatstorer := db.NewChatStorerMongo(mongoConnStr)

	log.Println("Connecting mongo")

	err := dbchatstorer.Connect()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("Is connected to Mongo")

	defer dbchatstorer.Disconnect()

	dbchatmanager := db.NewChatManagerRedis(
		redisConnStr,
		redisPass,
		redisDB,
	)

	err = dbchatmanager.Connect()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer dbchatmanager.Disconnect()

	chatmanagerservice := service.NewChatManagerService(dbchatmanager)

	var chat *model.Chat

	chat, err = chatmanagerservice.CreateChat("nombre", "descripcion")

	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	log.Println(chat)

	chatstorerservice := service.NewChatStorerService(dbchatstorer)

	var entry *model.ChatEntry

	entry, err = chatstorerservice.StoreChat(chat.ID, "Nickname", []byte("example of message"))

	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	log.Println(entry)

	log.Println("Finish program")
}
