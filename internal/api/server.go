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
	db             config.Database
	accountHandler *account.Handler
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := config.ConnectToDatabase()
	validate := validator.New()

	accountRepo := account.NewRepository(db)
	accountService := account.NewService(accountRepo, validate)
	accountHandler := account.NewHandler(accountService)

	NewServer := &Server{
		port:           port,
		db:             db,
		accountHandler: accountHandler,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
