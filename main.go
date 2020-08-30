package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		if _, err := fmt.Fprintln(os.Stderr, "Ошибка в библиотеке"); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	}
	hours, minutes, seconds := time.Clock()
	fmt.Printf("Текущее время: %d:%d:%d", hours, minutes, seconds)
}
