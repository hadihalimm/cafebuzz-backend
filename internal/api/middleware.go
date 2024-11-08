package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
	_ "github.com/joho/godotenv/autoload"
)

func (s *Server) RequireAuth(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing."})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if (len(authToken) != 2) || (authToken[0] != "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format."})
		return
	}

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token."})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired."})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims["account_type"] == "personal" {
		var account models.PersonalAccount
		s.DB.Gorm.Where("UUID = ?", claims["uuid"]).Find(&account)
		if account.Account.UUID == uuid.Nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("currentAccount", account.Account.UUID)
	}
	if claims["account_type"] == "cafe" {
		var cafe models.CafeAccount
		s.DB.Gorm.Where("UUID = ?", claims["uuid"]).Find(&cafe)
		if cafe.Account.UUID == uuid.Nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("currentAccount", cafe.Account.UUID)
	}

	c.Set("accountType", claims["account_type"])
	c.Next()

}
