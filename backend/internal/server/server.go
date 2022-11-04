package server

import (
	"context"
	"net/http"
	"time"
)

// import (
// 	"web/config"
// 	"web/internal/services"

// 	"github.com/gin-gonic/gin"
// )

// func SetupServer(
// 	cfg *config.Config,
// 	userService services.UserService,
// 	playerService services.PlyerService,
// 	teamService services.TeamService,
// ) *gin.Engine {
// 	r := gin.Default()
// 	r.Use(
// 		CORSMiddleware(),
// 		gin.Logger(),
// 		gin.Recovery(),
// 	)

// 	// r.GET("/ping", service.Ping)

// 	api := r.Group("/api/v1")

// 	user := api.Group("/users")
// 	user.GET("/", userService.GetUsers)
// 	// user.POST("/", userService.AddUser)
// 	// user.DELETE("/:email", service.DeleteUser)
// 	// user.PATCH("/role/:email", service.MakeMod)

// 	// user.POST("/login", service.Login)

// 	player := api.Group("/players")
// 	player.GET("/", playerService.GetPlayers)

// 	team := api.Group("/teams")
// 	team.GET("/", teamService.GetTeams)

// 	return r
// }

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT, PATCH")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
