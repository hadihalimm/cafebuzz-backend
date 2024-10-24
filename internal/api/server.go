package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/cafebuzz-backend/internal/config"
	"github.com/hadihalimm/cafebuzz-backend/internal/handler"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
	"github.com/hadihalimm/cafebuzz-backend/internal/repository"
	"github.com/hadihalimm/cafebuzz-backend/internal/services"
)

type Server struct {
	port           int
	DB             *config.Database
	accountHandler *handler.AccountHandler
}

func NewServer() (*http.Server, *Server) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	DB := config.ConnectToDatabase()
	DB.DropTable(&models.Account{})
	DB.AutoMigrate(&models.Account{})
	validate := validator.New()

	accountRepo := repository.NewAccountRepository(DB)
	accountService := services.NewAccountService(accountRepo, validate)
	accountHandler := handler.NewAccountHandler(accountService)

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
