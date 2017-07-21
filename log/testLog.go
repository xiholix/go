package main

import (
	"logger"
	"strconv"
	"time"
)

func OneProcess(i string) {
	for {
		logger.StageError("hello world " + i)
		logger.StageInfo("what's up " + i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	logger.Init("log")

	for i := 0; i < 20; i++ {
		go OneProcess(strconv.Itoa(i))
	}

	for {

	}
}
