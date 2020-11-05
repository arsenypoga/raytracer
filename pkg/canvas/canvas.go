package canvas

import (
	"image"
	clr "image/color"
	"image/png"
	"os"
)

// Canvas is a canvas for drawing
type Canvas struct {
	Width  int
	Height int
	pixels [][]Color
}

// NewCanvas creates new Canvas with given Width and Height
func NewCanvas(width, height int) *Canvas {
	pixels := make([][]Color, height)
	for i := range pixels {
		pixels[i] = make([]Color, width)
	}

	return &Canvas{
		Width:  width,
		Height: height,
		pixels: pixels,
	}
}

// WritePixel writes Color at given coordinates
func (c Canvas) WritePixel(i, j int, color Color) *Canvas {
	c1 := c

	c1.pixels = make([][]Color, c.Height)
	for i := range c1.pixels {
		c1.pixels[i] = make([]Color, c.Width)
		copy(c1.pixels[i], c.pixels[i])
	}
	c1.pixels[i][j] = color

	return &c1
}

// GetPixel returns Color at given coordinates
func (c Canvas) GetPixel(i, j int) *Color {
	return &c.pixels[i][j]
}

// Save saves image at the given path
func (c Canvas) Save(filename string) error {
	img := image.NewRGBA64(image.Rectangle{image.Point{0, 0}, image.Point{c.Width, c.Height}})
	for i := range c.pixels {
		for j := range c.pixels[i] {
			color := clr.RGBA{
				R: uint8(255 * c.pixels[i][j].R),
				G: uint8(255 * c.pixels[i][j].G),
				B: uint8(255 * c.pixels[i][j].B),
				A: 255,
			}
			img.Set(i, j, color)
		}
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	png.Encode(file, img)
	return nil
}
