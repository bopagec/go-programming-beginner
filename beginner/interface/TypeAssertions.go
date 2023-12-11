package main

import "fmt"

func main() {
	test(email{
		message: message{
			isSubscribed: true,
			body:         "hello, there",
		},
		address: "charatha@gmail.com",
	})

	test(email{
		message: message{
			isSubscribed: false,
			body:         "hello, there",
		},
		address: "charatha@gmail.com",
	})

	test(email{
		message: message{
			isSubscribed: true,
			body:         "this meeting cloud have been an email",
		},
		address: "tyrone@gmail.com",
	})

	test(email{
		message: message{
			isSubscribed: false,
			body:         "this meeting cloud have been an email",
		},
		address: "tryone@gmail.com",
	})
	test(sms{
		message: message{
			isSubscribed: true,
			body:         "this meeting cloud have been an email",
		},
		number: "07983397030",
	})
	test(sms{
		message: message{
			isSubscribed: false,
			body:         "this meeting cloud have been an email",
		},
		number: "07983397030",
	})

	test(mms{message{
		isSubscribed: false,
		body:         "invalid mms mesage",
	}})
}

func test(e expenses) {
	address, cost := getExpensesReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: the email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("================================================================")
	case sms:
		fmt.Printf("Report: the sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("================================================================")
	default:
		fmt.Printf("Report: invalid expenses\n")
		fmt.Println("================================================================")
	}
}

type expenses interface {
	cost() float64
}

func getExpensesReport(e expenses) (string, float64) {
	email, ok := e.(email)
	if ok {
		return email.address, email.cost()
	}
	sms, ok := e.(sms)
	if ok {
		return sms.number, sms.cost()
	}
	return "", 0.0
}

type message struct {
	isSubscribed bool
	body         string
}

type email struct {
	message
	address string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}
	return float64(len(e.body)) * 0.01
}

type sms struct {
	message
	number string
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.1
	}
	return float64(len(s.body)) * 0.03
}

type mms struct {
	message
}

func (m mms) cost() float64 {
	if !m.isSubscribed {
		return float64(len(m.body)) * 0.1
	}
	return float64(len(m.body)) * 0.03
}
