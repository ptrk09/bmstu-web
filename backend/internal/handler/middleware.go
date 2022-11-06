package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	usertIdCtx          = "userId"
	userRoleCtx         = "role"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userCtx, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	ctx.Set(usertIdCtx, userCtx.UserId)
	ctx.Set(userRoleCtx, userCtx.Role)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(usertIdCtx)

	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id is not found")
		return 0, errors.New("user id is not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is not found")
	}

	return idInt, nil
}

func getUserRole(ctx *gin.Context) (string, error) {
	role, ok := ctx.Get(userRoleCtx)

	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user role is not found")
		return "", errors.New("user role is not found")
	}

	stringRole, ok := role.(string)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user role is of invalid type")
		return "", errors.New("user role is not found")
	}

	return stringRole, nil
}
