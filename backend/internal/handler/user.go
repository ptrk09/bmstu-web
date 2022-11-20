package handler

import (
	"net/http"
	"strconv"
	"web/internal/model"

	"github.com/gin-gonic/gin"
)

type getUsersResponse struct {
	Data []model.User `json:"data"`
}

// getAllUsers godoc
// @Summary Get all users info
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} getUsersResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Router /user/ [get]
func (h *Handler) getAllUsers(ctx *gin.Context) {
	users, err := h.services.User.GetAllUsers()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getUsersResponse{
		Data: users,
	})
}

// getUserById godoc
// @Summary Get user info by user id
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} getUsersResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id query int true "user id"
// @Router /user/{id} [get]
func (h *Handler) getUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.services.User.GetUserById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getUsersResponse{
		Data: []model.User{user},
	})
}

// updateUser godoc
// @Summary Update user info
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id query int true "user id"
// @Router /user/{id} [patch]
func (h *Handler) updateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.UpdateUserInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.services.User.UpdateUser(id, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{userId})
}

// deleteUser godoc
// @Summary Delete user info
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id query int true "user id"
// @Router /user/{id} [delete]
func (h *Handler) deleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	userId, err := h.services.User.DeleteUser(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{userId})
}
