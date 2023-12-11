package main

import "fmt"

func main() {
	testReport(emailMsg{
		messageModel: messageModel{
			isSubscribed: true,
			body:         "hello, there",
		},
		address: "charatha@gmail.com",
	})

	testReport(emailMsg{
		messageModel: messageModel{
			isSubscribed: false,
			body:         "hello, there",
		},
		address: "charatha@gmail.com",
	})

	testReport(emailMsg{
		messageModel: messageModel{
			isSubscribed: true,
			body:         "this meeting cloud have been an email",
		},
		address: "tyrone@gmail.com",
	})

	testReport(emailMsg{
		messageModel: messageModel{
			isSubscribed: false,
			body:         "this meeting cloud have been an email",
		},
		address: "tryone@gmail.com",
	})
	testReport(smsMsg{
		messageModel: messageModel{
			isSubscribed: true,
			body:         "this meeting cloud have been an email",
		},
		number: "07983397030",
	})
	testReport(smsMsg{
		messageModel: messageModel{
			isSubscribed: false,
			body:         "this meeting cloud have been an email",
		},
		number: "07983397030",
	})

	testReport(mmsMsg{
		messageModel: messageModel{
			isSubscribed: false,
			body:         "invalid mms mesage",
		}})
}

func testReport(e expensesInt) {
	address, cost := getExpensesReportWithTypeSwitching(e)
	switch e.(type) {
	case emailMsg:
		fmt.Printf("Report: the email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("================================================================")
	case smsMsg:
		fmt.Printf("Report: the sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("================================================================")
	default:
		fmt.Printf("Report: invalid expenses\n")
		fmt.Println("================================================================")
	}
}

type messageModel struct {
	isSubscribed bool
	body         string
}

type expensesInt interface {
	calCost() float64
}

func getExpensesReportWithTypeSwitching(e expensesInt) (string, float64) {
	switch v := e.(type) {
	case emailMsg:
		return v.address, v.calCost()
	case smsMsg:
		return v.number, v.calCost()
	default:
		return "invalid msg", 0.0
	}

}

// we're using this from TypeAssertions.go file
/*type message struct {
	isSubscribed bool
	body         string
}*/
type emailMsg struct {
	messageModel
	address string
}

func (e emailMsg) calCost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}
	return float64(len(e.body)) * 0.01
}

type smsMsg struct {
	messageModel
	number string
}

func (s smsMsg) calCost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.1
	}
	return float64(len(s.body)) * 0.03
}

type mmsMsg struct {
	messageModel
}

func (m mmsMsg) calCost() float64 {
	if !m.isSubscribed {
		return float64(len(m.body)) * 0.1
	}
	return float64(len(m.body)) * 0.03
}
