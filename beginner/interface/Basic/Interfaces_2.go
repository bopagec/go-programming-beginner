package main

import (
	"fmt"
	"time"
)

func main() {
	testMsg(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})

	testMsg(birthdayMessage{
		birthdayTime:  time.Date(2014, 03, 31, 0, 0, 0, 0, time.UTC),
		recipientName: "Tyrone Bopage",
	})

	testMsg(sendingReport{
		reportName:    "Second Report",
		numberOfSends: 15,
	})

	testMsg(birthdayMessage{
		birthdayTime:  time.Date(1979, 11, 30, 0, 0, 0, 0, time.UTC),
		recipientName: "Nush Wijayakoon",
	})
}
func testMsg(msg message) {
	sendMessage(msg)
	fmt.Println("===================================")
}

type message interface {
	getMessage() string
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi, %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("Your '%s' report is ready. You've sent %v messages", sr.reportName, sr.numberOfSends)
}
