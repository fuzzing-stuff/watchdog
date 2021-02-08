package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 30*time.Second, "Timeout for ping in sec (default: 30 sec)")
	listen := flag.String("listen", "", "URI to listen to")
	ping := flag.String("ping", "", "Ping watchdog")
	flag.Parse()

	log.Fatal(serviceBody(*listen, *ping, *timeout))
}
