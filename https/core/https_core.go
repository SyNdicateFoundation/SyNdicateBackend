package https_core

import (
	"SyNdicateBackend/common/configuration"
	"SyNdicateBackend/common/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

type HttpsServer struct {
	Router *gin.Engine
}

func middleware(context *gin.Context) {
	LogConnection(context)

	context.Next()
}

func (H *HttpsServer) ListenAndServe() {
	httpConfig := configuration.ConfigHolder.HTTPSServer

	if !httpConfig.Enabled {
		return
	}

	gin.SetMode(gin.ReleaseMode)

	H.Router = gin.New()
	H.Router.Use(middleware)
	H.Router.LoadHTMLGlob("assets/templates/*")
	H.Router.Static("/assets", "./assets")

	//Handling 404 error
	if unknw, ok := responses["not-found-screen"]; ok {
		H.Router.NoRoute(unknw.fn)
		delete(responses, "not-found-screen")
	}

	//Registering the Paths and responses
	for name, req := range responses {
		for _, address := range req.addresses {
			logger.Logger.Debug(fmt.Sprintf("Registering Route -> %s - %s <-", name, address))
			H.Router.Handle(req.method, address, req.fn)
		}
	}

	addr := fmt.Sprintf("%s:%d", configuration.ConfigHolder.HTTPSServer.Address, configuration.ConfigHolder.HTTPSServer.Port)

	LogInfo(fmt.Sprintf("Listening on %s", addr))

	var err error

	if httpConfig.TlsConfiguration.Enable {
		err = H.Router.RunTLS(addr, httpConfig.TlsConfiguration.CertFile, httpConfig.TlsConfiguration.KeyFile)
	} else {
		err = H.Router.Run(addr)
	}

	if err != nil {
		logger.Logger.Error(err)
	}
}
