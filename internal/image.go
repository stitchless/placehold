package internal

import (
	"bytes"
	"image/color"
	"math/rand"
	"strconv"

	"github.com/fogleman/gg"
)

type RGB struct {
	R, G, B int
}

func CreateImage(W, H int, writer *bytes.Buffer) {
	dc := gg.NewContext(W, H)

	// Draw the background.
	dc.DrawRectangle(0, 0, float64(W), float64(H))
	dc.SetColor(toPastel(randomColor()))
	dc.Fill()

	label := strconv.Itoa(W) + "X" + strconv.Itoa(H)

	WriteText(dc, label, float64(W), float64(H))

	err := dc.EncodePNG(writer)
	if err != nil {
		panic(err)
	}
}

func randomColor() color.RGBA {
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func toPastel(c color.RGBA) color.RGBA {
	return color.RGBA{
		R: uint8((uint16(c.R) + 255*2) / 3),
		G: uint8((uint16(c.G) + 255*2) / 3),
		B: uint8((uint16(c.B) + 255*2) / 3),
		A: c.A,
	}
}
