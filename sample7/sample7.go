

package concurrent

var message string
var isWorkDone bool

func GetMessageFromGoRoutine() string {
	go doWork()
	for !isWorkDone {}
	return message
}

func doWork() {
	message = "Hello, world!"
	isWorkDone = true
}
