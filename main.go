package main

import (
	"sync"
)

var wg sync.WaitGroup

var chSend chan (TopicMessage)
var chDone chan (bool)

func main() {

	chSend = make(chan TopicMessage)
	chDone = make(chan bool)

	wg = sync.WaitGroup{}

	wg.Add(2)

	// http server
	go func() {
		defer wg.Done()
		initHttpServer()
	}()

	// web socket server
	go func() {
		defer wg.Done()
		initWebSocketServer()
	}()

	wg.Wait()
}
