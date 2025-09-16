package main

import (
	"log"

	"github.com/cosmay-s/go_finance_tracker/config"
	"github.com/cosmay-s/go_finance_tracker/internal/database"
	"github.com/cosmay-s/go_finance_tracker/internal/handlers"
	"github.com/cosmay-s/go_finance_tracker/internal/models"
	"github.com/cosmay-s/go_finance_tracker/internal/repositories"
	"github.com/cosmay-s/go_finance_tracker/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключение к бд: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	app := echo.New()

	app.Validator = &CustomValidator{validator: validator.New()}

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	userRepo := repositories.NewUserRepository(db)
	authSevice := services.NewAuthService(userRepo, cfg.JwtSecret)
	authHandler := handlers.NewAuthHandler(authSevice)

	//ручки
	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Приложение запущено - порт:"+cfg.Port)
	})

	app.POST("/register", authHandler.Register)
	app.POST("/login", authHandler.Login)

	app.Logger.Fatal(app.Start(":" + cfg.Port))

}
