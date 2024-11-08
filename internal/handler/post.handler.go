package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/request"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/response"
	"github.com/hadihalimm/cafebuzz-backend/internal/services"
)

type PostHandler struct {
	service services.PostService
}

func NewPostHandler(service services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) Create(c *gin.Context) {
	var input request.PostCreateRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := response.Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	creatorUUID := uuid.MustParse(c.Param("uuid"))
	creatorType := c.GetString("userType")
	result, err := h.service.Create(input, creatorUUID, creatorType)
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
		Message: "Successfully created new post!",
		Data:    result,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *PostHandler) FindByID(c *gin.Context) {
	postID, _ := strconv.ParseUint(c.Param("postID"), 10, 64)
	post, err := h.service.FindByID(postID)
	if err != nil {
		response := response.Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := response.Response{
		Success: true,
		Message: "Successfully retrieved current cafe",
		Data:    post,
	}
	c.JSON(http.StatusOK, response)
}

func (h *PostHandler) FindAllByCreator(c *gin.Context) {
	creatorUUID := uuid.MustParse(c.Param("uuid"))
	posts, err := h.service.FindAllByCreator(creatorUUID)
	if err != nil {
		response := response.Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := response.Response{
		Success: true,
		Message: "Successfully retrieved current cafe",
		Data:    posts,
	}
	c.JSON(http.StatusOK, response)
}
