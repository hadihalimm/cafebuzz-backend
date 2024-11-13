package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/response"
	"github.com/hadihalimm/cafebuzz-backend/internal/services"
)

type FollowHandler struct {
	service services.FollowService
}

func NewFollowHandler(service services.FollowService) *FollowHandler {
	return &FollowHandler{service: service}
}

func (h *FollowHandler) CreateFollowPersonal(c *gin.Context) {
	followerUUID := uuid.MustParse(c.GetString("currentAccount"))
	result, err := h.service.Create(followerUUID, uuid.MustParse(c.Param("followedUUID")), "personal")
	if err != nil {
		response := response.Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.Response{
		Success: true,
		Message: "Successfully created new follow!",
		Data:    result,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *FollowHandler) CreateFollowCafe(c *gin.Context) {
	followerUUID := uuid.MustParse(c.GetString("currentAccount"))
	result, err := h.service.Create(followerUUID, uuid.MustParse(c.Param("followedUUID")), "cafe")
	if err != nil {
		response := response.Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.Response{
		Success: true,
		Message: "Successfully created new follow!",
		Data:    result,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *FollowHandler) GetAllFollowing(c *gin.Context) {
	personalFollowing, cafeFollowing, err := h.service.FindFollowingsByUUID(uuid.MustParse(c.Param("uuid")))
	if err != nil {
		response := response.Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.Response{
		Success: true,
		Message: "Successfully retrieved all followings!",
		Data:    []interface{}{personalFollowing, cafeFollowing},
	}
	c.JSON(http.StatusCreated, response)
}

func (h *FollowHandler) GetAllFollowers(c *gin.Context) {
	personalFollowing, cafeFollowing, err := h.service.FindFollowersByUUID(uuid.MustParse(c.Param("uuid")))
	if err != nil {
		response := response.Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.Response{
		Success: true,
		Message: "Successfully retrieved all followers!",
		Data:    []interface{}{personalFollowing, cafeFollowing},
	}
	c.JSON(http.StatusCreated, response)
}
