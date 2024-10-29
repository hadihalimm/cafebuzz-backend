package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/request"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/response"
	"github.com/hadihalimm/cafebuzz-backend/internal/services"
)

type AccountHandler struct {
	service services.PersonalAccountService
}

func NewPersonalAccountHandler(service services.PersonalAccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (h *AccountHandler) Register(c *gin.Context) {
	var input request.AccountRegisterRequest
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

func (h *AccountHandler) Login(c *gin.Context) {
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

func (h *AccountHandler) GetAccountDetails(c *gin.Context) {
	accountUUID := c.Param("uuid")
	account, err := h.service.Details(uuid.MustParse(accountUUID))
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
		Message: "Successfully retrieved current account.",
		Data:    account,
	}
	c.JSON(http.StatusOK, response)
}

func (h *AccountHandler) UpdateAccountDetails(c *gin.Context) {
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

	accountUUID := c.Param("uuid")
	updatedAccount, err := h.service.Update(uuid.MustParse(accountUUID), input)
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
