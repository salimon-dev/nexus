package mail

import "fmt"

type SendEmailJob struct {
	To      string
	Subject string
	Body    string
}

var EmailQueue chan SendEmailJob
var Done chan bool

func emailWorker() {
	for email := range EmailQueue {
		fmt.Printf("proccessing email to %s", email.To)
		err := SendRawEmail(email.To, email.Subject, email.Body)
		if err != nil {
			fmt.Printf("Failed to send email to %s: %v\n", email.To, err)
		}
	}
	Done <- true
}

func SetupEmailQueue() {
	EmailQueue = make(chan SendEmailJob, 32)
	Done = make(chan bool)

	go emailWorker()
}
