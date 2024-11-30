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

// CreatePost godoc
//
// @Summary Create a new post
// @Description create a new post
// @Tags post
// @Accept json
// @Produce json
// @Param request body request.PostCreateRequest true "Post create request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /account/{uuid}/post/create [post]
// @Router /cafe/{uuid}/post/create [post]
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

// FindByID godoc
//
// @Summary Find a post by ID
// @Description find a post by ID
// @Tags post
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /account/{uuid}/post/{postID} [get]
// @Router /cafe/{uuid}/post/{postID} [get]
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

// FindAllByCreator godoc
//
// @Summary Find all posts by creator
// @Description find all posts by creator
// @Tags post
// @Accept json
// @Produce json
// @Param uuid path string true "Creator UUID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /account/{uuid}/posts [get]
// @Router /cafe/{uuid}/posts [get]
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

// DeletePost godoc
//
// @Summary Delete a post
// @Description delete a post
// @Tags post
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /account/{uuid}/post/{postID} [delete]
// @Router /cafe/{uuid}/post/{postID} [delete]
func (h *PostHandler) DeletePost(c *gin.Context) {
	postID, _ := strconv.ParseUint(c.Param("postID"), 10, 64)
	err := h.service.Delete(postID)
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
		Message: "Successfully deleted the post.",
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}
