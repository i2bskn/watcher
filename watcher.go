package watcher

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"os"
)

type ConsulEvent struct {
	ID            string
	Name          string
	Payload       string
	NodeFilter    string
	ServiceFilter string
	TagFilter     string
	Version       int
	LTime         int
}

func (e ConsulEvent) ParsedPayload() string {
	payload, _ := base64.StdEncoding.DecodeString(e.Payload)
	return string(payload)
}

func Process(fn func(string) error) error {
	payload, err := parseInput()
	if err != nil {
		return err
	}

	if len(payload) == 0 {
		log.Println("Received events is empty.")
		return nil
	}

	if err := fn(payload); err != nil {
		return err
	}

	log.Println("Event process finished successfully.")
	return nil
}

func parseInput() (string, error) {
	log.Println("Waiting for events from STDIN...")

	scanner := bufio.NewScanner(os.Stdin)
	var payload string
	for scanner.Scan() {
		if os.Getenv("CONSUL_INDEX") != "" {
			var events []ConsulEvent
			if err := json.Unmarshal(scanner.Bytes(), &events); err != nil {
				return "", err
			}

			if len(events) == 0 {
				return "", nil
			}

			event := events[len(events)]
			log.Println("Reveive event:", event.ID)

			payload = event.ParsedPayload()
		} else {
			payload = string(scanner.Bytes())
		}
	}

	log.Println("Parse payload:", payload)

	return payload, nil
}
