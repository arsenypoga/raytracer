package canvas

// Color represents a color stucture
type Color struct {
	R float64
	G float64
	B float64
}

// Add adds two colors
func (c Color) Add(c1 *Color) *Color {
	return &Color{
		R: c.R + c1.R,
		G: c.G + c1.G,
		B: c.B + c1.B,
	}
}

// Substract substracts two colors
func (c Color) Substract(c1 *Color) *Color {
	return &Color{
		R: c.R - c1.R,
		G: c.G - c1.G,
		B: c.B - c1.B,
	}
}

// Scale scales the color by number
func (c Color) Scale(s float64) *Color {
	return &Color{
		R: c.R * s,
		G: c.G * s,
		B: c.B * s,
	}
}

// Multiply multiplies two colors.
func (c Color) Multiply(c1 *Color) *Color {
	return &Color{
		R: c.R * c1.R,
		G: c.G * c1.G,
		B: c.B * c1.B,
	}
}

// Clamp normalizes the color to be in range between 0 to 1
func (c Color) Clamp() *Color {
	c1 := c
	if c1.R > 1 {
		c1.R = 1
	}
	if c1.G > 1 {
		c1.G = 1
	}
	if c1.B > 1 {
		c1.B = 1
	}

	if c1.R < 0 {
		c1.R = 0
	}
	if c1.G < 0 {
		c1.G = 0
	}
	if c1.B < 0 {
		c1.B = 0
	}
	return &c1
}
