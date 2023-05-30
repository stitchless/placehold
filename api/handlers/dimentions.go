package handlers

import (
	"bytes"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jpeizer/placehold/internal"
)

func SimpleImageHandler(c *gin.Context) {
	width := c.Param("width")
	height := c.Param("height")

	if height == "" {
		height = width
	}

	// convert to int
	h, err := strconv.Atoi(height)
	if err != nil {
		panic(err)
	}

	w, err := strconv.Atoi(width)
	if err != nil {
		panic(err)
	}

	buff := new(bytes.Buffer)

	internal.CreateImage(w, h, buff)

	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(buff.Bytes())))
	c.Writer.Write(buff.Bytes())
}
