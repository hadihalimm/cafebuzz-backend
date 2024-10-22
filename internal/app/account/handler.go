package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/request"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/response"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(c *gin.Context) {
	var input request.RegisterRequest
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
		Message: "Successfully created new account!",
		Data:    result,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *Handler) Login(c *gin.Context) {
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

func (h *Handler) GetCurrentAccount(c *gin.Context) {
	account, _ := c.Get("currentUser")
	response := response.Response{
		Success: true,
		Message: "Successfully retrieved current account.",
		Data:    account,
	}
	c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateAccountDetails(c *gin.Context) {
	var input request.AccountUpdateRequest
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

	account, _ := c.Get("currentUser")
	updatedAccount, err := h.service.Update(account.(uuid.UUID), input)
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
		Message: "Successfully updated the account.",
		Data:    updatedAccount,
	}
	c.JSON(http.StatusCreated, response)
}
