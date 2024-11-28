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
	service services.CafeAccountService
}

func NewCafeAccountHandler(service services.CafeAccountService) *CafeHandler {
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

func (h *CafeHandler) GetCafeDetails(c *gin.Context) {
	cafeUUID := c.Param("uuid")
	cafe, err := h.service.Details(uuid.MustParse(cafeUUID))
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
		Data:    cafe,
	}
	c.JSON(http.StatusOK, response)
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

	cafeUUID := c.Param("uuid")
	updatedCafe, err := h.service.Update(uuid.MustParse(cafeUUID), input)
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

func (h *CafeHandler) DeleteCafe(c *gin.Context) {
	cafeUUID := c.Param("uuid")
	err := h.service.Delete(uuid.MustParse(cafeUUID))
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
		Message: "Successfully deleted the cafe.",
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}
