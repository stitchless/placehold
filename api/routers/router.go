package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jpeizer/placehold/api/handlers"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        false,
		AllowOrigins:           []string{"http://localhost:3000"},
		AllowOriginFunc:        nil,
		AllowMethods:           []string{"GET", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "Content-Type", "Accept", "X-Requested-With"},
		MaxAge:                 time.Minute * 2,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
	}))

	image := r.Group("/i")
	{
		image.GET("/:width/:height", handlers.SimpleImageHandler)
		image.GET("/:width", handlers.SimpleImageHandler)
	}

	return r
}
