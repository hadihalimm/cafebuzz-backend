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

// RegisterPersonalAccount godoc
//
// @Summary Register a new personal account
// @Description create a new personal account
// @Tags personalAccount
// @Accept json
// @Produce json
// @Param request body request.AccountRegisterRequest true "Account register request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /account/register [post]
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

// LoginPersonalAccount godoc
//
// @Summary Login a personal account
// @Description authenticate & authorize a personal account
// @Tags personalAccount
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Router /account/login [post]
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

// GetAccountDetails godoc
//
// @Summary Retrieve current account details
// @Description get current account details
// @Tags personalAccount
// @Accept json
// @Produce json
// @Param uuid path string true "Account UUID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Failure 404 {object} response.Response "Not Found"
// @Router /account{uuid} [get]
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

// UpdateAccountDetails godoc
//
// @Summary Update current account details
// @Description update current account details
// @Tags personalAccount
// @Accept json
// @Produce json
// @Param request body request.AccountUpdateRequest true "Account update request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Failure 404 {object} response.Response "Not Found"
// @Router /account{uuid} [put]
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

// DeleteAccount godoc
//
// @Summary Delete current account
// @Description delete current account
// @Tags personalAccount
// @Accept json
// @Produce json
// @Param uuid path string true "Account UUID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response "Bad Request"
// @Failure 404 {object} response.Response "Not Found"
// @Router /account{uuid} [delete]
func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	accountUUID := c.Param("uuid")
	err := h.service.Delete(uuid.MustParse(accountUUID))
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
		Message: "Successfully deleted the account.",
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}
