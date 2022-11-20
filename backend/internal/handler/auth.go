package handler

import (
	"net/http"
	"web/internal/model"

	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type resultWithId struct {
	ID int `json:"id" binding:"required"`
}

type resultWithToken struct {
	Token string `json:"token" binding:"required"`
}

// signUp godoc
// @Summary Sign-up user
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} resultWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param request body model.User true "User's name, login, password and role"
// @Router /signUp/ [post]
func (h *Handler) signUp(ctx *gin.Context) {
	var input model.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resultWithId{id})
}

// signIn godoc
// @Summary Sign-in user
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} resultWithToken
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param request body signInInput true "User's login and password"
// @Router /signIn/ [post]
func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resultWithToken{token})
}
