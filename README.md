# watcher

`watcher` provide some processing by consul events.

## example

```go
package main

import (
	"github.com/i2bskn/watcher"
	"log"
)

func eventProcess(events []watcher.Event) error {
	for _, event := range events {
		log.Println("Receive Payload:", event.ParsedPayload())
	}
}

func main() {
	watcher.ProcessWithEvents(eventProcess)
}
```

## service-operator

```
GOOS=linux GOARCH=amd64 make install
echo reload:nginx | service-operator // Execute `systemctl reload nginx`
```