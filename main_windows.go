package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/windows/svc"
)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"Help message\n\n",
		errmsg, os.Args[0])
	os.Exit(2)
}

func main() {
	const serviceName = "watchdog"

	inService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("Error checking of service status: %v", err)
	}
	if inService {
		runService(serviceName, false)
		return
	}

	if len(os.Args) < 2 {
		usage("no command specified")
	}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "debug":
		runService(serviceName, true)
		return
	case "install":
		err = installService(serviceName, "Watchdog service")
	case "remove":
		err = removeService(serviceName)
	case "start":
		err = startService(serviceName)
	case "stop":
		err = controlService(serviceName, svc.Stop, svc.Stopped)
	case "pause":
		err = controlService(serviceName, svc.Pause, svc.Paused)
	case "continue":
		err = controlService(serviceName, svc.Continue, svc.Running)
	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, serviceName, err)
	}
	return
}
