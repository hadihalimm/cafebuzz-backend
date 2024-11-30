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

// CreateFollowPersonal godoc
//
// @Summary Create a new follow by a personal account
// @Description create a new follow by a personal account
// @Tags follow
// @Accept json
// @Produce json
// @Param followerUUID header string true "Follower UUID"
// @Param followedUUID path string true "Followed UUID"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /follow/create/{followerUUID}/personal/{followedUUID} [post]
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

// CreateFollowCafe godoc
//
// @Summary Create a new follow by a cafe account
// @Description create a new follow by a cafe account
// @Tags follow
// @Accept json
// @Produce json
// @Param followerUUID header string true "Follower UUID"
// @Param followedUUID path string true "Followed UUID"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /follow/create/{followerUUID}/cafe/{followedUUID} [post]
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

// GetAllFollowing godoc
//
// @Summary Get all followings by uuid
// @Description get all followings by uuid
// @Tags follow
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /follow/following/{uuid} [get]
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

// GetAllFollowers godoc
//
// @Summary Get all followers by uuid
// @Description get all followers by uuid
// @Tags follow
// @Accept  json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /follow/followers/{uuid} [get]
func (h *FollowHandler) GetAllFollowers(c *gin.Context) {
	personalFollowers, cafeFollowers, err := h.service.FindFollowersByUUID(uuid.MustParse(c.Param("uuid")))
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
		Data:    []interface{}{personalFollowers, cafeFollowers},
	}
	c.JSON(http.StatusCreated, response)
}

// DeleteFollow godoc
//
// @Summary Delete a follow
// @Description delete a follow
// @Tags follow
// @Accept json
// @Produce json
// @Param followedUUID path string true "Followed UUID"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /follow/delete/{followerUUID}/{followedUUID} [delete]
func (h *FollowHandler) Delete(c *gin.Context) {
	followerUUID := uuid.MustParse(c.GetString("currentAccount"))
	followedUUID := uuid.MustParse(c.Param("followedUUID"))
	err := h.service.Delete(followerUUID, followedUUID)
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
		Message: "Successfully deleted the follow!",
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}
