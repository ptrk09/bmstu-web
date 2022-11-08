package handler

import (
	"web/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("v1/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("v1/", h.userIdentity)
	{

		users := api.Group("users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}

		listings := api.Group("listings")
		{
			listings.GET("/", h.getListings)
			listings.POST("/", h.createListing)
			listings.PATCH("/", h.updateListing)
			listings.DELETE("/", h.deleteListing)
		}

		// teams := api.Group("teams")
		// {
		// 	teams.POST("/", h.createTeam)
		// 	teams.GET("/", h.getAllTeams)
		// 	teams.GET("/:id", h.getTeamById)
		// 	teams.PUT("/:id", h.updateTeam)
		// 	teams.DELETE("/:id", h.deleteTeam)
		// }

	}

	return router
}
