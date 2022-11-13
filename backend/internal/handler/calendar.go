package handler

import (
	"net/http"
	"strconv"
	"web/internal/model"

	"github.com/gin-gonic/gin"
)

type getAllCalendarInfoResponse struct {
	Data []model.Calendar `json:"data"`
}

func (h *Handler) getAllCalendarInfo(ctx *gin.Context) {
	calendar, err := h.services.Calendar.GetAllCalendarInfo()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllCalendarInfoResponse{
		Data: calendar,
	})
}

func (h *Handler) createCalendarInfo(ctx *gin.Context) {
	var input model.Calendar

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Calendar.CreateCalendarInfo(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateCalendarInfo(ctx *gin.Context) {
	var input model.Calendar

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input.ID = id

	if _, err := h.services.Calendar.UpdateCalendarInfo(input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteCalendarInfo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	_, err = h.services.Calendar.DeleteCalendarInfo(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
