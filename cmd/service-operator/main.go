package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/i2bskn/watcher"
)

const (
	Version = "20181221"
)

func systemctl(payload string) error {
	params := strings.Split(payload, ":")
	if len(params) != 2 {
		return fmt.Errorf("Invalid payload: %v", payload)
	}

	log.Println("Execute systemctl:", params)
	out, err := exec.Command("systemctl", params...).Output()
	if err != nil {
		log.Println("Executing command error:", err)
		return err
	}

	log.Println(string(out))
	log.Println("Operation successfully completed")

	return nil
}

func main() {
	var printVersion bool
	flag.BoolVar(&printVersion, "version", false, "Print version")
	flag.Parse()

	if printVersion {
		fmt.Println("service-operator", Version)
		os.Exit(0)
	}

	if err := watcher.Process(systemctl); err != nil {
		log.Fatal(err)
	}
}
