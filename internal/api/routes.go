package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("/register", s.accountHandler.Register)
		v1.POST("/login", s.accountHandler.Login)

		account := v1.Group("/account")
		{
			account.GET("/", s.RequireAuth, s.accountHandler.GetCurrentAccount)
		}
	}

	return r
}
