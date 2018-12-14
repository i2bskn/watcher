# watcher

`watcher` provide some processing by consul events.

## example

```go
package main

import (
	"github.com/i2bskn/watcher"
	"log"
)

func eventProcess(payload string) error {
	log.Println("Receive Payload:", payload)
}

func main() {
	watcher.Process(eventProcess)
}
```