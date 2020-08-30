package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("%v", err)
	}
	hours, minutes, seconds := time.Clock()
	fmt.Printf("Текущее время: %d:%d:%d", hours, minutes, seconds)
}
