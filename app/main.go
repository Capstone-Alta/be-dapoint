package main

import (
	"dapoint-api/api"
	"dapoint-api/app/modules"
	"dapoint-api/config"
	"dapoint-api/util"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"

	//_ "dapoint-api/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Clean Hexa Sample API
// @version 1.0
// @description Berikut API yang digunakan untuk mini project
func main() {
	config := config.GetConfig()

	dbCon := util.NewConnectionDatabase(config)
	defer dbCon.CloseConnection()

	controllers := modules.RegisterModules(dbCon, config)

	e := echo.New()
	//e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin},
	}))

	handleSwag := echoSwagger.WrapHandler
	e.GET("/swagger/*", handleSwag)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "dapoint API")
	})
	e.POST("/callback", func(c echo.Context) error {
		var iface interface{}
		c.Bind(&iface)
		asByteJson, _ := json.Marshal(iface)
		fmt.Println("masuk : ", string(asByteJson))
		return c.JSON(200, "ok")
	})
	api.RegistrationPath(e, controllers)

	go func() {
		var appAddress string

		if config.App.Env == "dev" {
			appAddress = "127.0.0.1"
		} else {
			appAddress = "0.0.0.0"
		}
		address := fmt.Sprintf("%s:%d", appAddress, config.App.Port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	<-quit
}
