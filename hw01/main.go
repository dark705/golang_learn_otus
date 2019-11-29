package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

const (
	ntpHost         = "3.europe.pool.ntp.org"
	layOutFormat    = "2006-01-02 15:04:05"
	successExitCode = 0
	errorExitCode   = 1
)

func main() {
	currentTime, err := ntp.Time(ntpHost)
	if err != nil {
		fmt.Println(err)
		os.Exit(errorExitCode)
	}

	localTime := time.Now()

	fmt.Println("Текущее время:", localTime.Format(layOutFormat))
	fmt.Println("Точное  время:", currentTime.Format(layOutFormat))
	os.Exit(successExitCode)
}
