package main

import (
	"logger"
)

func main() {
	log := logger.Logger{Log_dir: "log"}
	log.Init()

	log.StageError("hello world")
	log.StageInfo("what's up")
}
