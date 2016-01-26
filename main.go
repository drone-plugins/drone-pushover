package main

import (
	"fmt"
)

var (
	buildDate string
)

func main() {
	fmt.Printf("Drone Pushover Plugin built at %s\n", buildDate)
}
