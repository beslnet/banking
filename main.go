package main

import (
	"github.com/beslnet/banking/app"
	"github.com/beslnet/banking/app/logger"
)

func main() {
	logger.Info("Starting our application")
	app.Start()
}
