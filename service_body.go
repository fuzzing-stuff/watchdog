package main

import (
	"fmt"
	"net/http"
	"time"
)

func pingHandler(c chan bool) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		c <- true
	}
}
func serviceBody(url string, ping string, timeout time.Duration) error {
	c := make(chan bool)

	http.HandleFunc("/", pingHandler(c))

	go func() {
		for {
			select {
			case <-c:
				continue
			case <-time.After(timeout):
				fmt.Println("No answer", timeout.String())
			}
		}
	}()

	return http.ListenAndServe(url, nil)
}
