package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/meziaris/gofinance/internal/app/controller"
	"github.com/meziaris/gofinance/internal/app/repository"
	"github.com/meziaris/gofinance/internal/app/service"
	"github.com/meziaris/gofinance/internal/pkg/config"
	"github.com/meziaris/gofinance/internal/pkg/middleware"
)

var (
	cfg    config.Config
	DBConn *sqlx.DB
)

func init() {
	loadConfig, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load app config")
	}
	cfg = loadConfig

	db, _ := sqlx.Connect(cfg.DBDriver, cfg.DBConnection)
	DBConn = db

}

func main() {
	// repository
	userRepo := repository.NewUserRepository(DBConn)
	authRepo := repository.NewAuthRepository(DBConn)

	// service
	tokenCreator := service.NewTokenCreator(
		cfg.JwtAccessTokenKey,
		cfg.JwtRefreshTokenKey,
		cfg.JwtAccessTokenDuration,
		cfg.RefreshTokenDuration,
	)
	userService := service.NewRegistrationService(userRepo)
	sessionService := service.NewSessionService(userRepo, authRepo, tokenCreator)

	// controller
	registrationController := controller.NewRegistrationController(userService)
	sessionController := controller.NewSessionController(sessionService, tokenCreator)

	// router
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Post("/auth/register", registrationController.Register)
	r.Post("/auth/login", sessionController.Login)
	r.Post("/auth/refresh", sessionController.Refresh)

	r.Route("/auth/logout", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(tokenCreator))
		r.Get("/", sessionController.Logout)
	})

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.ServerPort),
		Handler:           r,
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       3 * time.Minute,
		WriteTimeout:      5 * time.Minute,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
