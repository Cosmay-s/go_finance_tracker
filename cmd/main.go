package main

import (
	"log"

	"github.com/cosmay-s/go_finance_tracker/config"
	"github.com/cosmay-s/go_finance_tracker/internal/database"
	"github.com/cosmay-s/go_finance_tracker/internal/handlers"
	"github.com/cosmay-s/go_finance_tracker/internal/middleware"
	"github.com/cosmay-s/go_finance_tracker/internal/models"
	"github.com/cosmay-s/go_finance_tracker/internal/repositories"
	"github.com/cosmay-s/go_finance_tracker/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
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

	app.Use(echoMiddleware.Logger())
	app.Use(echoMiddleware.Recover())
	app.Use(echoMiddleware.CORS())

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, cfg.JwtSecret)
	authHandler := handlers.NewAuthHandler(authService)

	jwtMiddleware := middleware.JWTMiddleware(authService)

	//ручки
	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Приложение запущено - порт: "+cfg.Port)
	})

	app.POST("/register", authHandler.Register)
	app.POST("/login", authHandler.Login)

	protected := app.Group("/api")
	protected.Use(jwtMiddleware)

	// ручка для проверки
	protected.GET("/profile", func(c echo.Context) error {
		userID, err := services.GetUserIDFromContext(c)
		if err != nil {
			return echo.NewHTTPError(401, err.Error())
		}

		return c.JSON(200, map[string]interface{}{
			"message": "Доступ разрешен",
			"user_id": userID,
		})
	})
	app.Logger.Fatal(app.Start(":" + cfg.Port))
}
