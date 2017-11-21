package main

import (
	"net/http"

	"github.com/byuoitav/authmiddleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	port := ":5000"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	secure.GET("/health", health)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)

}

func health(context echo.Context) error {

	return context.JSON(http.StatusOK, "the fleet has moved out of lightspeed and we're preparing to - augh!")

}
