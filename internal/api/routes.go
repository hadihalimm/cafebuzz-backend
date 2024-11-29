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
			account.DELETE("/:uuid", s.RequireAuth, s.accountHandler.DeleteAccount)
			post := account.Group("/:uuid")
			{
				post.POST("/post/create", s.RequireAuth, s.postHandler.Create)
				post.GET("/posts", s.RequireAuth, s.postHandler.FindAllByCreator)
				post.GET("/post/:postID", s.RequireAuth, s.postHandler.FindByID)
				post.DELETE("/post/:postID", s.RequireAuth, s.postHandler.DeletePost)
			}
		}
		cafe := v1.Group("/cafe")
		{
			cafe.POST("/login", s.cafeHandler.Login)
			cafe.POST("/register", s.cafeHandler.Register)
			cafe.GET("/:uuid", s.RequireAuth, s.cafeHandler.GetCafeDetails)
			cafe.PUT("/:uuid", s.RequireAuth, s.cafeHandler.UpdateCafeDetails)
			cafe.DELETE("/:uuid", s.RequireAuth, s.cafeHandler.DeleteCafe)
			post := cafe.Group("/:uuid")
			{
				post.POST("/post/create", s.RequireAuth, s.postHandler.Create)
				post.GET("/posts", s.RequireAuth, s.postHandler.FindAllByCreator)
				post.GET("/post/:postID", s.RequireAuth, s.postHandler.FindByID)
				post.DELETE("/post/:postID", s.RequireAuth, s.postHandler.DeletePost)
			}
		}
		follow := v1.Group("/follow")
		{
			follow.POST("/create/:followerUUID/personal/:followedUUID", s.RequireAuth, s.followHandler.CreateFollowPersonal)
			follow.POST("/create/:followerUUID/cafe/:followedUUID", s.RequireAuth, s.followHandler.CreateFollowCafe)
			follow.GET("/following/:uuid", s.RequireAuth, s.followHandler.GetAllFollowing)
			follow.GET("/followers/:uuid", s.RequireAuth, s.followHandler.GetAllFollowers)
			follow.DELETE("/delete/:followerUUID/:followedUUID", s.RequireAuth, s.followHandler.Delete)
		}
	}

	return r
}
