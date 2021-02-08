package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	version string
)

func ping(c chan bool) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		c <- true
	}
}

func printVersionInfo() {
	fmt.Println("watchdog version in " + version)
}

func main() {
	timeout := flag.Duration("timeout", 30*time.Second, "Timeout for ping in sec (default: 30 sec)")
	url := flag.String("url", ":8090", "URL for check")
	verInfo := flag.Bool("version", false, "Print version info")
	flag.Parse()

	if *verInfo {
		printVersionInfo()
		return
	}

	c := make(chan bool)

	http.HandleFunc("/", ping(c))

	go func() {
		for {
			select {
			case <-c:
				continue
			case <-time.After(*timeout):
				fmt.Println("No answer", timeout.String())
			}
		}
	}()

	log.Fatal(http.ListenAndServe(*url, nil))
}
