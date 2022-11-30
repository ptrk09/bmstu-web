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

// getAllCalendarInfo godoc
// @Summary Get all calendar info
// @Tags calendar
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllCalendarInfoResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Router /calendar/ [get]
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

// createCalendarInfo godoc
// @Summary Create calendar info
// @Tags calendar
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param request body model.Calendar true "Calendar's listing id, date id and available status"
// @Router /calendar/ [post]
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

	ctx.JSON(http.StatusOK, responseWithId{id})
}

// updateCalendarInfo godoc
// @Summary Update calendar info
// @Tags calendar
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id path int true "calendar id"
// @Router /calendar/{id} [patch]
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

	calendarId, err := h.services.Calendar.UpdateCalendarInfo(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{calendarId})
}

// deleteCalendarInfo godoc
// @Summary Dalete calendar info
// @Tags calendar
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id path int true "calendar id"
// @Router /calendar/{id} [delete]
func (h *Handler) deleteCalendarInfo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	calendarId, err := h.services.Calendar.DeleteCalendarInfo(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{calendarId})
}
