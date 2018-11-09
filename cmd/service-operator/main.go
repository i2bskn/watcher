package main

import (
	"fmt"
	"github.com/i2bskn/watcher"
	"log"
	"os/exec"
	"strings"
)

func systemctl(events []watcher.Event) error {
	if len(events) > 0 {
		event := events[0]

		params, err := parsePayload(event.ParsedPayload())
		if err != nil {
			log.Println("Parsing payload:", err)
			return err
		}

		log.Println("Execute systemctl:", params)
		out, err := exec.Command("systemctl", params...).Output()
		if err != nil {
			log.Println("Executing command:", err)
			return err
		}

		log.Println(string(out))
		log.Println("Operation successfully completed")
	}

	return nil
}

func parsePayload(payload string) ([]string, error) {
	params := strings.Split(payload, ":")
	if len(params) != 2 {
		return []string{}, fmt.Errorf("Invalid payload: %v", payload)
	}

	return params, nil
}

func main() {
	watcher.ProcessWithEvents(systemctl)
}
