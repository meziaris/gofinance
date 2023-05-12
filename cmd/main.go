package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/meziaris/gofinance/internal/app/controller"
	"github.com/meziaris/gofinance/internal/app/repository"
	"github.com/meziaris/gofinance/internal/app/service"
	"github.com/meziaris/gofinance/internal/pkg/config"
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

	// service
	userService := service.NewRegistrationService(userRepo)

	// controller
	registrationController := controller.NewRegistrationController(userService)

	// router
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Post("/auth/register", registrationController.Register)

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
