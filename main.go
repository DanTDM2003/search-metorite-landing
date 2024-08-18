package main

// import "log"

// type Notifier interface {
// 	Send(message string) error
// 	Validate() error
// }

// type EmailNotifier struct {
// 	// fields related to email notification
// }

// func (e EmailNotifier) Send(message string) error {
// 	// logic to send an email
// 	return nil
// }

// type SMSNotifier struct {
// 	// fields related to SMS notification
// }

// func (s SMSNotifier) Send(message string) error {
// 	// logic to send an SMS
// 	return nil
// }

// func NotifyAll(notifiers []Notifier, message string) {
// 	for _, notifier := range notifiers {
// 		err := notifier.Send(message)
// 		if err != nil {
// 			log.Println("Failed to send message:", err)
// 		}
// 	}
// }

// func main() {
// 	notifiers := []Notifier{EmailNotifier{}, SMSNotifier{}}

// 	NotifyAll(notifiers, "Hello, World!")
// }
