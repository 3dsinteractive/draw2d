// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 21/11/2010 by Laurent Le Goff, Stani Michiels

// Load a png image and rotate it
package main

import (
	"math"

	"github.com/stanim/draw2d"
	"github.com/stanim/draw2d/pdf2d"
)

func main() {
	// Margin between the image and the frame
	const margin = 30
	// Line width od the frame
	const lineWidth = 3

	// Initialize the graphic context on a pdf document
	dest := pdf2d.NewPdf("P", "mm", "A3")
	gc := pdf2d.NewGraphicContext(dest)
	// Size of destination image
	dw, dh := dest.GetPageSize()
	// Draw frame
	draw2d.RoundRect(gc, lineWidth, lineWidth, dw-lineWidth, dh-lineWidth, 100, 100)
	gc.SetLineWidth(lineWidth)
	gc.FillStroke()

	// load the source image
	source, err := draw2d.LoadFromPngFile("gopher.png")
	if err != nil {
		panic(err)
	}
	// Size of source image
	sw, sh := float64(source.Bounds().Dx()), float64(source.Bounds().Dy())
	// Draw image to fit in the frame
	// TODO Seems to have a transform bug here on draw image
	scale := math.Min((dw-margin*2)/sw, (dh-margin*2)/sh)
	gc.Save()
	gc.Translate(margin, margin)
	gc.Scale(scale, scale)

	gc.DrawImage(source)
	gc.Restore()

	// Save to pdf
	pdf2d.SaveToPdfFile("frame-image.pdf", dest)
}
