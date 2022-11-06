package handler

// func (h *Handler) createTeam(ctx *gin.Context) {
// 	id, err := getUserId(ctx)
// 	if err != nil {
// 		return
// 	}

// 	var input model.Team
// 	if err := ctx.BindJSON(&input); err != nil {
// 		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	idTeam, err := h.services.Team.Create(id, input)
// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"id": idTeam,
// 	})
// }

// type getAllListsResponse struct {
// 	Data []model.Team `json:"data"`
// }

// func (h *Handler) getAllTeams(ctx *gin.Context) {
// 	userId, err := getUserId(ctx)
// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	teams, err := h.services.Team.GetAll(userId)
// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, getAllListsResponse{
// 		Data: teams,
// 	})

// }

// func (h *Handler) getTeamById(ctx *gin.Context) {

// }

// func (h *Handler) updateTeam(ctx *gin.Context) {

// }

// func (h *Handler) deleteTeam(ctx *gin.Context) {

// }
