//+build ignore

package concurrent

import "runtime"

func GetMessageFromGoRoutine() string {
	runtime.GOMAXPROCS(1) // Just for demonstration – not needed.
	messageChannel := make(chan string)
	go doWork(messageChannel)
	return <-messageChannel
}

func doWork(messageChannel chan string) {
	messageChannel <- "Hello, world!"
}
