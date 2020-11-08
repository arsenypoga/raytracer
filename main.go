package main

import (
	"fmt"

	"github.com/arsenypoga/go-raytracer/pkg/units"
)

func main() {
	m := units.NewMatrix([][]float64{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	})
	fmt.Println(m.String())
	fmt.Println(m.Cofactor(2, 3))
	fmt.Println(m.Invert())
}
