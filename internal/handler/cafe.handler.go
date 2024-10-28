package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/request"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/response"
	"github.com/hadihalimm/cafebuzz-backend/internal/services"
)

type CafeHandler struct {
	service services.CafeService
}

func NewCafeHandler(service services.CafeService) *CafeHandler {
	return &CafeHandler{service: service}
}

func (h *CafeHandler) Register(c *gin.Context) {
	var input request.CafeRegisterRequest
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

	result, err := h.service.Register(input)
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
		Message: "Successfully created new cafe!",
		Data:    result,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *CafeHandler) Login(c *gin.Context) {
	var input request.LoginRequest
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

	result, err := h.service.Login(input)
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
		Message: "Login successful!",
		Data: gin.H{
			"token": result,
		},
	}
	c.JSON(http.StatusCreated, response)
}

func (h *CafeHandler) UpdateCafeDetails(c *gin.Context) {
	var input request.CafeUpdateRequest
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

	cafe, _ := c.Get("currentUser")
	updatedCafe, err := h.service.Update(cafe.(uuid.UUID), input)
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
		Message: "Successfully updated the cafe.",
		Data:    updatedCafe,
	}
	c.JSON(http.StatusCreated, response)
}
