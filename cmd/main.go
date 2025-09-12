package main

import (
	"github.com/cosmay-s/go_finance_tracker/config"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Приложение запущено"+cfg.Port)
	})

	app.Logger.Fatal(app.Start(":" + cfg.Port))

}
