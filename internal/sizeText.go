package internal

import (
	"image/color"
	"log"
	"os"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

func WriteText(dc *gg.Context, text string, x float64, y float64) {
	// Ignore text if it's too small.
	if x < 100 {
		return
	}

	fontBytes, err := os.ReadFile("assets/Poppins-Regular.ttf")
	if err != nil {
		panic(err)
	}

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	fontSize := 48.0

	if x < 250 || y < 100 {
		fontSize = 24.0
	}

	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})

	dc.SetFontFace(face)

	dc.SetColor(color.RGBA{R: 0, G: 0, B: 0, A: 255})
	n := 2 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := x/2 + float64(dx)
			y := y/2 + float64(dy)
			dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
		}
	}

	dc.SetColor(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	dc.DrawStringAnchored(text, float64(x)/2, float64(y)/2, 0.5, 0.5)
}
