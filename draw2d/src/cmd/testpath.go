package main

import (
	"draw2d"
	"fmt"
)

func main() {
	path := new(draw2d.Path)
	path.MoveTo(2.0, 3.0)
	path.LineTo(2.0, 3.0)
	path.QuadCurveTo(2.0, 3.0, 10, 20)
	path.CubicCurveTo(2.0, 3.0, 10, 20, 13, 23)
	path.Rect(2.0, 3.0, 100, 200)
	path.ArcTo(2.0, 3.0, 100, 200, 200, 300)
	fmt.Printf("%v\n", path)
}
