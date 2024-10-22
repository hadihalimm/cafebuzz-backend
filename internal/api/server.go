package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/cafebuzz-backend/internal/app/account"
	"github.com/hadihalimm/cafebuzz-backend/internal/config"
)

type Server struct {
	port           int
	DB             *config.Database
	accountHandler *account.Handler
}

func NewServer() (*http.Server, *Server) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	DB := config.ConnectToDatabase()
	DB.DropTable(&account.Account{})
	DB.AutoMigrate(&account.Account{})
	validate := validator.New()

	accountRepo := account.NewRepository(DB)
	accountService := account.NewService(accountRepo, validate)
	accountHandler := account.NewHandler(accountService)

	NewServer := &Server{
		port:           port,
		DB:             DB,
		accountHandler: accountHandler,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, NewServer
}
