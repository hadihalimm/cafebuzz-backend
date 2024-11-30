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

// RegisterCafeAccount godoc
//
// @Summary Register a new cafe account
// @Description create a new cafe account
// @Tags cafeAccount
// @Accept json
// @Produce json
// @Param request body request.CafeRegisterRequest true "Cafe register request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /cafe/register [post]
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

// LoginCafeAccount godoc
//
// @Summary Login a cafe account
// @Description authenticate & authorize a cafe account
// @Tags cafeAccount
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /cafe/login [post]
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

// GetCafeDetails godoc
//
// @Summary Get current cafe details
// @Description get current cafe details
// @Tags cafeAccount
// @Accept json
// @Produce json
// @Param uuid path string true "Cafe UUID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /cafe{uuid} [get]
func (h *CafeHandler) GetCafeDetails(c *gin.Context) {
	cafeUUID := c.Param("uuid")
	cafe, err := h.service.Details(uuid.MustParse(cafeUUID))
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
		Message: "Successfully retrieved current cafe",
		Data:    cafe,
	}
	c.JSON(http.StatusOK, response)
}

// UpdateCafeDetails godoc
//
// @Summary Update current cafe details
// @Description update current cafe details
// @Tags cafeAccount
// @Accept json
// @Produce json
// @Param request body request.CafeUpdateRequest true "Cafe update request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /cafe{uuid} [put]
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

// DeleteCafe godoc
//
// @Summary Delete current cafe
// @Description delete current cafe
// @Tags cafeAccount
// @Accept json
// @Produce json
// @Param uuid path string true "Cafe UUID"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /cafe{uuid} [delete]
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
