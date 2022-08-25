package logging

import (
	"log"
	"os"

	"github.com/Shell32-Natsu/zuzu_bank/internal/config"
)

var logger = log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lmsgprefix)

func LogDebugf(format string, args ...any) {
	if !config.IsDebug() {
		return
	}
	logger.Printf(format, args...)
}

func LogDebug(s any) {
	LogDebugf("%s", s)
}
