package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Action string `json:"action"`
	Topic  string `json:"topic"`
}

func initHttpServer() {
	router := gin.Default()
	router.POST("/", postMessage)
	router.Run("localhost:3001")
}

func postMessage(c *gin.Context) {
	var msg Message

	c.BindJSON(&msg)

	fmt.Printf("action:%s, topic: '%s'\n", msg.Action, msg.Topic)

	switch msg.Action {
	case "quit":
		chDone <- true
		fmt.Println("http server recieved quit signal")
		wg.Done()
	case "send":
		send(msg)
	}
}

func send(msg Message) {
	chMsg := TopicMessage{
		Topic:   msg.Topic,
		Payload: "tbd",
	}
	fmt.Printf("http server sends on topic '%s'\n", msg.Topic)
	chSend <- chMsg
}
