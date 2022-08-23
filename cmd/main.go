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
	c, err := config.NewConfig(*configPath)
	if err != nil {
		log.Panic(err)
	}

	if *debug {
		c.Debug = true
	}

	// Start bot. This function is not expected to return.
	if err := bot.Start(c); err != nil {
		log.Panic(err)
	}
}
