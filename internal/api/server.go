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
	cafeHandler    *handler.CafeHandler
}

func NewServer() (*http.Server, *Server) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	DB := config.ConnectToDatabase()
	DB.DropTable(&models.Account{}, &models.PersonalAccount{}, &models.CafeAccount{})
	DB.AutoMigrate(&models.Account{}, &models.PersonalAccount{}, &models.CafeAccount{})
	validate := validator.New()

	accountRepo := repository.NewPersonalAccountRepository(DB)
	accountService := services.NewPersonalAccountService(accountRepo, validate)
	accountHandler := handler.NewPersonalAccountHandler(accountService)

	cafeRepo := repository.NewCafeAccountRepository(DB)
	cafeService := services.NewCafeAccountService(cafeRepo, validate)
	cafeHandler := handler.NewCafeAccountHandler(cafeService)

	NewServer := &Server{
		port:           port,
		DB:             DB,
		accountHandler: accountHandler,
		cafeHandler:    cafeHandler,
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
