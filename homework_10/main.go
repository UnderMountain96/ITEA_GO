package main

import (
	"log"

	"github.com/UnderMountain96/ITEA_GO/http_server/handler"
	"github.com/UnderMountain96/ITEA_GO/http_server/middleware"
	"github.com/UnderMountain96/ITEA_GO/http_server/server"
)

func main() {
	logger := log.Default()

	requestLoggerMiddleware := middleware.NewRequestLogger(logger)
	// authenticateMiddleware := middleware.NewAuthenticate("validToken")

	homeHandler := handler.NewHomeHandler(logger)
	userHandler := handler.NewUserHandler(logger)

	apiServer := server.NewAPIServer(logger)
	apiServer.AddRoute("/", requestLoggerMiddleware.Wrap(homeHandler))
	apiServer.AddRoute("/user", requestLoggerMiddleware.Wrap(userHandler))

	if err := apiServer.Start(); err != nil {
		logger.Fatal(err)
	}
}
