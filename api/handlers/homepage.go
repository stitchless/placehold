package handlers

import (
	"github.com/gin-gonic/gin"
)

func HomePageHandler(c *gin.Context) {
	// serve the homepage and css
	c.HTML(200, "index.html", gin.H{})
}
