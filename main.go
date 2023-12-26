package main

import (
	"learnGo/config"
	"learnGo/misc"
)

func initializeMinicalcRouter() {
	config.ConfigMinicalcRouter()
}

func initializeBasicRouter() {
	config.ConfigBasicPost()
}

func main() {
	// Create a channel to signal when each function is done
	done1 := make(chan bool)
	done2 := make(chan bool)

	// Execute ByteSliceTutorial and wait for it to finish
	go func() { // Anonymous function
		misc.ByteSliceTutorial()
		done1 <- true
	}() //Execute/invoke the ananymous function right away.

	/*
		This operation blocks the current goroutine until a value
		is received from the done channel.When a value is sent into
		the done channel (e.g., done <- true), the <-done operation
		unblocks and continues with the value that was received.
		In the case of done <- true, it will receive true.
	*/
	<-done1

	go func() {
		misc.PrintStructTags()
		done2 <- true
	}()

	// Wait for PrintStructTags to finish
	<-done2

	go initializeMinicalcRouter()
	initializeBasicRouter()
}
