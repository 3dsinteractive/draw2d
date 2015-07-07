// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 13/12/2010 by Laurent Le Goff

// Package draw2d is a pure go 2D vector graphics library with support
// for multiple output devices such as images (draw2d), pdf documents
// (draw2dpdf) and opengl (draw2dopengl), which can also be used on the
// google app engine.
//
// Features
//
// Operations in draw2d include stroking and filling polygons, arcs,
// Bézier curves, drawing images and text rendering with truetype fonts.
// All drawing operations can be transformed by affine transformations
// (scale, rotation, translation).
//
// Installation
//
// To install or update the package draw2d on your system, run:
//   go get -u github.com/llgcode/draw2d
//
// Quick Start
//
// Package draw2d itself provides a graphic context that can draw vector
// graphics and text on an image canvas. The following Go code
// generates a simple drawing and saves it to an image file:
//   // Initialize the graphic context on an RGBA image
//   dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
//   gc := draw2d.NewGraphicContext(dest)
//
//   // Set some properties
//   gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
//   gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
//   gc.SetLineWidth(5)
//
//   // Draw a closed shape
//   gc.MoveTo(10, 10) // should always be called first for a new path
//   gc.LineTo(100, 50)
//   gc.QuadCurveTo(100, 10, 10, 10)
//   gc.Close()
//   gc.FillStroke()
//
//   // Save to file
//   draw2d.SaveToPngFile(fn, dest)
//
// There are more examples here:
// https://github.com/llgcode/draw2d.samples
//
// Drawing on pdf documents is provided by the draw2dpdf package.
// Drawing on opengl is provided by the draw2dgl package.
// See subdirectories at the bottom of this page.
//
// Acknowledgments
//
// Laurent Le Goff wrote this library, inspired by postscript and
// HTML5 canvas. He implemented the image and opengl backend. Also
// he created a pure go Postscripter interpreter which can draw to a
// draw2d graphic context (https://github.com/llgcode/ps). Stani
// Michiels implemented the pdf backend.
//
// The package depends on freetype-go package for its rasterization
// algorithm.
//
// Packages using draw2d
//
// - https://github.com/llgcode/ps
//
// - https://github.com/gonum/plot
package draw2d
