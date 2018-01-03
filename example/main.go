package main

import "github.com/iain17/logger"

func main() {
	logger.AddOutput(logger.Stdout{
		MinLevel: logger.INFO, //logger.DEBUG,
		Colored:  true,
	})
	fileOut, err := logger.NewFileOut("/tmp/test.log")
	if err != nil {
		panic(err)
	}
	logger.AddOutput(fileOut)

	logger.Infof("test!")
}
