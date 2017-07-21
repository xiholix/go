package main

import (
	"logger"
	"time"
)

func main() {
	logger.Init("log2")
	for {
		logger.StageError("hello world2222")
		logger.StageInfo("what's up2222")
		time.Sleep(2 * time.Second)
	}

}
