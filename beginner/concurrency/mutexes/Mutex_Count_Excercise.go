package main

import (
	"fmt"
	"sort"
	"sync"
)

type safeCounter struct {
	counts map[string]int
	mux    *sync.Mutex
}

type emailStruct struct {
	email string
	count int
}

func (sc safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	count := sc.counts[key]
	count++
	sc.counts[key] = count
}

func (sc safeCounter) val(key string) int {
	return sc.counts[key]
}

func testEmail(sc safeCounter, emailStructs []emailStruct) {
	emails := make(map[string]struct{})

	var wg sync.WaitGroup

	for _, emailS := range emailStructs {
		emails[emailS.email] = struct{}{}
		for i := 0; i < emailS.count; i++ {
			wg.Add(1)
			go func(s emailStruct) {
				sc.inc(s.email)
				wg.Done()
			}(emailS)
		}
	}
	wg.Wait()

	emailSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailSorted = append(emailSorted, email)
	}
	sort.Strings(emailSorted)

	for _, email := range emailSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=====================================")
}

func main() {
	sc := safeCounter{
		counts: make(map[string]int),
		mux:    &sync.Mutex{},
	}

	testEmail(sc, []emailStruct{
		{
			email: "john@example.com",
			count: 23,
		},
		{
			email: "john@example.com",
			count: 29,
		},
		{
			email: "jill@example.com",
			count: 31,
		},
		{
			email: "jill@example.com",
			count: 67,
		},
	})
	testEmail(sc, []emailStruct{
		{
			email: "kaden@example.com",
			count: 23,
		},
		{
			email: "george@example.com",
			count: 126,
		},
		{
			email: "kaden@example.com",
			count: 31,
		},
		{
			email: "george@example.com",
			count: 453,
		},
	})
}

func test(sc safeCounter, emails []emailStruct) {

}
