package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	usertCtx            = "userId"
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

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	ctx.Set(usertCtx, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(usertCtx)

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
