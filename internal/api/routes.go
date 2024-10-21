package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("/account", s.accountHandler.Register)
	}

	return r
}
