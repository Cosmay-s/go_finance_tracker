package main

import (
	"log"

	"github.com/cosmay-s/go_finance_tracker/config"
	"github.com/cosmay-s/go_finance_tracker/internal/database"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключение к бд: %v", err)
	}
	app := echo.New()

	_ = db

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Приложение запущено"+cfg.Port)
	})

	app.Logger.Fatal(app.Start(":" + cfg.Port))

}
