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
	postHandler    *handler.PostHandler
	followHandler  *handler.FollowHandler
}

func NewServer() (*http.Server, *Server) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	DB := config.ConnectToDatabase()
	DB.DropTable(&models.Account{}, &models.PersonalAccount{}, &models.CafeAccount{}, &models.Post{}, &models.Follow{})
	DB.AutoMigrate(&models.PersonalAccount{}, &models.CafeAccount{}, &models.Post{}, &models.Follow{})
	validate := validator.New()

	accountRepo := repository.NewPersonalAccountRepository(DB)
	accountService := services.NewPersonalAccountService(accountRepo, validate)
	accountHandler := handler.NewPersonalAccountHandler(accountService)

	cafeRepo := repository.NewCafeAccountRepository(DB)
	cafeService := services.NewCafeAccountService(cafeRepo, validate)
	cafeHandler := handler.NewCafeAccountHandler(cafeService)

	postRepo := repository.NewPostRepository(DB)
	postService := services.NewPostService(postRepo, validate)
	postHandler := handler.NewPostHandler(postService)

	followRepo := repository.NewFollowRepository(DB)
	followService := services.NewFollowService(followRepo, validate)
	followHandler := handler.NewFollowHandler(followService)

	NewServer := &Server{
		port:           port,
		DB:             DB,
		accountHandler: accountHandler,
		cafeHandler:    cafeHandler,
		postHandler:    postHandler,
		followHandler:  followHandler,
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
