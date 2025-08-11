package main

import (
	command "SyNdicateBackend/common"
	"SyNdicateBackend/common/configuration"
	"SyNdicateBackend/common/logger"
	httpscore "SyNdicateBackend/https/core"
	"github.com/muesli/termenv"
	"os"
)

var (
	HttpsServer httpscore.HttpsServer
	terminal    = termenv.NewOutput(os.Stdout)
)

func main() {
	logger.SetupLogger()
	configuration.SetupConfig()

	HttpsServer.ListenAndServe()

	logger.Logger.Info("running rescue mode since https option is off")
	logger.Logger.Info("type \"help\" for command list")

	command.Loop(terminal)
}
