/*************************************.
*	Author : Ameya Salagre
*	github : github.com/ameyasalagre
*	Entry Point of Repository
**************************************/
package main

import (
	"server"
	"utils"
)
/************************.
*	Call to The Server
************************/
var Version = "1.0.0"

func main() {
	logger := utils.New().With(nil, "version", Version)
	logger.Info("Server started")
	server.Start(logger)
}
