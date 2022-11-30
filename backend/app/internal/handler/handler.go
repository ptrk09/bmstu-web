package handler

import (
	"web/internal/service"

	_ "web/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("api/v1/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("api/v1/", h.userIdentity)
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
			listing.PATCH("/:id", h.updateListing)
			listing.DELETE("/:id", h.deleteListing)
		}

		listingDetailed := api.Group("listing_detailed")
		{
			listingDetailed.GET("/", h.getListingsDetailed)
			listingDetailed.POST("/", h.createListingDetailed)
			listingDetailed.PATCH("/:id", h.updateListingDetailed)
			listingDetailed.DELETE("/:id", h.deleteListingDetailed)
		}

		listingImage := api.Group("listing_image")
		{
			listingImage.GET("/", h.getListingImages)
			listingImage.POST("/", h.createListingImage)
			listingImage.PATCH("/:id", h.updateListingImage)
			listingImage.DELETE("/:id", h.deleteListingImage)
		}

		calendar := api.Group("calendar")
		{
			calendar.GET("/", h.getAllCalendarInfo)
			calendar.POST("/", h.createCalendarInfo)
			calendar.PATCH("/:id", h.updateCalendarInfo)
			calendar.DELETE("/:id", h.deleteCalendarInfo)
		}

		booking := api.Group("booking")
		{
			booking.GET("/", h.getBookings)
			booking.POST("/", h.createBooking)
			booking.PATCH("/:id", h.updateBooking)
			booking.DELETE("/:id", h.deleteBooking)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
