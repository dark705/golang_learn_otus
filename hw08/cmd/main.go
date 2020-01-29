package main

import (
	"fmt"
	"os"

	"github.com/dark705/otus/hw08/internal/Calendar/Calendar"
	"github.com/dark705/otus/hw08/internal/Config"
	"github.com/dark705/otus/hw08/internal/Logger"
	"github.com/dark705/otus/hw08/internal/Storage"
)

func main() {

	config, err := Config.ReadFromFile("config/config.yaml")
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
	logger := Logger.GetLogger(config)

	inMemory := Storage.InMemory{}
	inMemory.Init()
	_ = Calendar.Calendar{Config: config, Storage: &inMemory, Logger: logger}
}
