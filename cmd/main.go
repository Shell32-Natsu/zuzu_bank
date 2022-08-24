package main

import (
	"flag"
	"log"

	"github.com/Shell32-Natsu/zuzu_bank/internal/bot"
	"github.com/Shell32-Natsu/zuzu_bank/internal/config"
)

func main() {
	// init flags
	configPath := flag.String("config", "a", "path to the configuration file")
	debug := flag.Bool("debug", false, "enable debug")
	flag.Parse()

	if *configPath == "" {
		log.Panicln("must provide config file.")
	}
	err := config.InitConfig(*configPath)
	if err != nil {
		log.Panic(err)
	}

	config.SetDebug(*debug)

	// Start bot. This function is not expected to return.
	if err := bot.Start(); err != nil {
		log.Panic(err)
	}
}
