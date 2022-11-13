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

		listing := api.Group("listing")
		{
			listing.GET("/", h.getListings)
			listing.POST("/", h.createListing)
			listing.PATCH("/", h.updateListing)
			listing.DELETE("/", h.deleteListing)
		}

		listingDetailed := api.Group("listing_detailed")
		{
			listingDetailed.GET("/", h.getListingsDetailed)
			listingDetailed.POST("/", h.createListingDetailed)
			listingDetailed.PATCH("/", h.updateListingDetailed)
			listingDetailed.DELETE("/", h.deleteListingDetailed)
		}

		listingImage := api.Group("listings_images")
		{
			listingImage.GET("/", h.getListingImages)
			listingImage.POST("/", h.createListingImage)
			listingImage.PATCH("/", h.updateListingImage)
			listingImage.DELETE("/", h.deleteListingImage)
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
