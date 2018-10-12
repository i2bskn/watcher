package watcher

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
)

type Event struct {
	ID            string
	Name          string
	Payload       string
	OriginPayload string
}

func (e Event) ParsedPayload() string {
	if len(e.OriginPayload) > 0 {
		return e.OriginPayload
	}

	payload, _ := base64.StdEncoding.DecodeString(e.Payload)
	return string(payload)
}

func ProcessWithEvents(fn func([]Event) error) error {
	log.Println("Waiting for events from STDIN...")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var events []Event
		if err := json.Unmarshal(scanner.Bytes(), &events); err != nil {
			events = append(events, Event{OriginPayload: string(scanner.Bytes())})
		}

		if err := fn(events); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Reading STDIN:", err)
		return err
	}

	return nil
}
