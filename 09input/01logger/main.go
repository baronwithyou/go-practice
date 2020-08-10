package main

import (
	"os"

	"github.com/baronwithyou/go-practice/09input/logger"
)

func main() {
	file, err := os.OpenFile("./logger.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error(err)
		return
	}

	logger.SetOutput(file)

	logger.Info("Helo baron")
}
