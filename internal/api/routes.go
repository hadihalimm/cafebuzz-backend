package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{

		account := v1.Group("/account")
		{
			account.POST("/login", s.accountHandler.Login)
			account.POST("/register", s.accountHandler.Register)
			account.GET("/:uuid", s.RequireAuth, s.accountHandler.GetAccountDetails)
			account.PUT("/:uuid", s.RequireAuth, s.accountHandler.UpdateAccountDetails)
			post := account.Group("/:uuid")
			{
				post.GET("/posts", s.postHandler.FindAllByCreator)
				post.GET("/post/:postID", s.postHandler.FindByID)
			}
		}
		cafe := v1.Group("/cafe")
		{
			cafe.POST("/login", s.cafeHandler.Login)
			cafe.POST("/register", s.cafeHandler.Register)
			cafe.GET("/:uuid", s.RequireAuth, s.cafeHandler.GetCafeDetails)
			cafe.PUT("/:uuid", s.RequireAuth, s.cafeHandler.UpdateCafeDetails)
			post := cafe.Group("/:uuid")
			{
				post.GET("/posts", s.postHandler.FindAllByCreator)
				post.GET("/post/:postID", s.postHandler.FindByID)
			}
		}
	}

	return r
}
