// Copyright 2010 The postscript-go Authors. All rights reserved.
// created: 13/12/2010 by Laurent Le Goff

package postscript

import (
	"log"
	"image"
	"draw2d.googlecode.com/svn/trunk/draw2d/src/pkg/draw2d"
	"math"
)


// begin Primitive Operator implementation


//Path Construction Operators
func newpath(interpreter *Interpreter) {
	interpreter.GetGraphicContext().BeginPath()
}

func closepath(interpreter *Interpreter) {
	interpreter.GetGraphicContext().Close()
}

func currentpoint(interpreter *Interpreter) {
	x, y := interpreter.GetGraphicContext().LastPoint()
	interpreter.Push(x)
	interpreter.Push(y)
}

func moveto(interpreter *Interpreter) {
	y := interpreter.PopFloat()
	x := interpreter.PopFloat()
	interpreter.GetGraphicContext().MoveTo(x, y)
}

func rmoveto(interpreter *Interpreter) {
	y := interpreter.PopFloat()
	x := interpreter.PopFloat()
	interpreter.GetGraphicContext().RMoveTo(x, y)
}

func lineto(interpreter *Interpreter) {
	y := interpreter.PopFloat()
	x := interpreter.PopFloat()
	interpreter.GetGraphicContext().LineTo(x, y)
}

func rlineto(interpreter *Interpreter) {
	y := interpreter.PopFloat()
	x := interpreter.PopFloat()
	interpreter.GetGraphicContext().RLineTo(x, y)
}

func curveto(interpreter *Interpreter) {
	cy3 := interpreter.PopFloat()
	cx3 := interpreter.PopFloat()
	cy2 := interpreter.PopFloat()
	cx2 := interpreter.PopFloat()
	cy1 := interpreter.PopFloat()
	cx1 := interpreter.PopFloat()
	interpreter.GetGraphicContext().CubicCurveTo(cx1, cy1, cx2, cy2, cx3, cy3)
}

func rcurveto(interpreter *Interpreter) {
	cy3 := interpreter.PopFloat()
	cx3 := interpreter.PopFloat()
	cy2 := interpreter.PopFloat()
	cx2 := interpreter.PopFloat()
	cy1 := interpreter.PopFloat()
	cx1 := interpreter.PopFloat()
	interpreter.GetGraphicContext().RCubicCurveTo(cx1, cy1, cx2, cy2, cx3, cy3)
}

func clippath(interpreter *Interpreter) {

}

func stroke(interpreter *Interpreter) {
	interpreter.GetGraphicContext().Stroke()
}

func fill(interpreter *Interpreter) {
	interpreter.GetGraphicContext().Fill()
}

func gsave(interpreter *Interpreter) {
	interpreter.GetGraphicContext().Save()
}

func grestore(interpreter *Interpreter) {
	interpreter.GetGraphicContext().Restore()
}

func setgray(interpreter *Interpreter) {
	gray := interpreter.PopFloat()
	color := image.RGBAColor{uint8(gray * 0xff), uint8(gray * 0xff), uint8(gray * 0xff), 0xff}
	interpreter.GetGraphicContext().SetStrokeColor(color)
	interpreter.GetGraphicContext().SetFillColor(color)
}

func setrgbcolor(interpreter *Interpreter) {
	blue := interpreter.PopFloat()
	green := interpreter.PopFloat()
	red := interpreter.PopFloat()
	color := image.RGBAColor{uint8(red * 0xff), uint8(green * 0xff), uint8(blue * 0xff), 0xff}
	interpreter.GetGraphicContext().SetStrokeColor(color)
	interpreter.GetGraphicContext().SetFillColor(color)
}

func hsbtorgb(hue, saturation, brightness float) (red, green, blue int) {
	var fr, fg, fb float
	if saturation == 0 {
		fr, fg, fb = brightness, brightness, brightness
	} else {
		H := (hue - float(math.Floor(float64(hue)))) * 6
		I := int(math.Floor(float64(H)))
		F := H - float(I)
		M := brightness * (1 - saturation)
		N := brightness * (1 - saturation*F)
		K := brightness * (1 - saturation*(1-F))

		switch I {
		case 0:
			fr = brightness
			fg = K
			fb = M
		case 1:
			fr = N
			fg = brightness
			fb = M
		case 2:
			fr = M
			fg = brightness
			fb = K
		case 3:
			fr = M
			fg = N
			fb = brightness
		case 4:
			fr = K
			fg = M
			fb = brightness
		case 5:
			fr = brightness
			fg = M
			fb = N
		default:
			fr, fb, fg = 0, 0, 0
		}
	}

	red = int(fr*255. + 0.5)
	green = int(fg*255. + 0.5)
	blue = int(fb*255. + 0.5)
	return
}

func sethsbcolor(interpreter *Interpreter) {
	brightness := interpreter.PopFloat()
	saturation := interpreter.PopFloat()
	hue := interpreter.PopFloat()
	red, green, blue := hsbtorgb(hue, saturation, brightness)
	color := image.RGBAColor{uint8(red), uint8(green), uint8(blue), 0xff}
	interpreter.GetGraphicContext().SetStrokeColor(color)
	interpreter.GetGraphicContext().SetFillColor(color)
}

func setcmybcolor(interpreter *Interpreter) {
	black := interpreter.PopFloat()
	yellow := interpreter.PopFloat()
	magenta := interpreter.PopFloat()
	cyan := interpreter.PopFloat()

	/*  cyan = cyan / 255.0;   
	    magenta = magenta / 255.0;   
	    yellow = yellow / 255.0;   
	    black = black / 255.0;   */

	red := cyan*(1.0-black) + black
	green := magenta*(1.0-black) + black
	blue := yellow*(1.0-black) + black

	red = (1.0-red)*255.0 + 0.5
	green = (1.0-green)*255.0 + 0.5
	blue = (1.0-blue)*255.0 + 0.5

	color := image.RGBAColor{uint8(red), uint8(green), uint8(blue), 0xff}
	interpreter.GetGraphicContext().SetStrokeColor(color)
	interpreter.GetGraphicContext().SetFillColor(color)
}

func setdash(interpreter *Interpreter) {
	offset := interpreter.PopInt()
	dash := interpreter.PopArray()
	log.Printf("dash: %v, offset: %d \n", dash, offset)
}

func setlinejoin(interpreter *Interpreter) {
	linejoin := interpreter.PopInt()
	switch linejoin {
	case 0:
		interpreter.GetGraphicContext().SetLineJoin(draw2d.MiterJoin)
	case 1:
		interpreter.GetGraphicContext().SetLineJoin(draw2d.RoundJoin)
	case 2:
		interpreter.GetGraphicContext().SetLineJoin(draw2d.BevelJoin)
	}
}

func setlinecap(interpreter *Interpreter) {
	linecap := interpreter.PopInt()
	switch linecap {
	case 0:
		interpreter.GetGraphicContext().SetLineCap(draw2d.ButtCap)
	case 1:
		interpreter.GetGraphicContext().SetLineCap(draw2d.RoundCap)
	case 2:
		interpreter.GetGraphicContext().SetLineCap(draw2d.SquareCap)
	}
}

func setmiterlimit(interpreter *Interpreter) {
	interpreter.PopInt()
}

func setlinewidth(interpreter *Interpreter) {
	interpreter.GetGraphicContext().SetLineWidth(interpreter.PopFloat())
}

func showpage(interpreter *Interpreter) {

}

func show(interpreter *Interpreter) {
	s := interpreter.PopString()
	interpreter.GetGraphicContext().FillString(s)
}

func findfont(interpreter *Interpreter) {

}

func scalefont(interpreter *Interpreter) {

}

func setfont(interpreter *Interpreter) {

}

func stringwidth(interpreter *Interpreter) {
	interpreter.Push(10.0)
	interpreter.Push(10.0)
}

func setflat(interpreter *Interpreter) {
	interpreter.Pop()
}

func currentflat(interpreter *Interpreter) {
	interpreter.Push(1.0)
}


// Coordinate System and Matrix operators
func matrix(interpreter *Interpreter) {
	interpreter.Push(draw2d.NewIdentityMatrix())
}

func transform(interpreter *Interpreter) {
	y := interpreter.PopFloat()
	x := interpreter.PopFloat()
	interpreter.GetGraphicContext().GetMatrixTransform().Transform(&x, &y)
	interpreter.Push(x)
	interpreter.Push(y)
}

func itransform(interpreter *Interpreter) {
	y := interpreter.PopFloat()
	x := interpreter.PopFloat()
	interpreter.GetGraphicContext().GetMatrixTransform().InverseTransform(&x, &y)
	interpreter.Push(x)
	interpreter.Push(y)
}

func translate(interpreter *Interpreter) {
	y := interpreter.PopFloat()
	x := interpreter.PopFloat()
	interpreter.GetGraphicContext().Translate(x, y)
}

func rotate(interpreter *Interpreter) {
	angle := interpreter.PopFloat()
	interpreter.GetGraphicContext().Rotate(angle * (math.Pi / 180.0))
}

func scale(interpreter *Interpreter) {
	y := interpreter.PopFloat()
	x := interpreter.PopFloat()
	interpreter.GetGraphicContext().Scale(x, y)
}


func initDrawingOperators(interpreter *Interpreter) {

	interpreter.SystemDefine("stroke", NewOperator(stroke))
	interpreter.SystemDefine("fill", NewOperator(fill))
	interpreter.SystemDefine("show", NewOperator(show))
	interpreter.SystemDefine("showpage", NewOperator(showpage))

	interpreter.SystemDefine("findfont", NewOperator(findfont))
	interpreter.SystemDefine("scalefont", NewOperator(scalefont))
	interpreter.SystemDefine("setfont", NewOperator(setfont))
	interpreter.SystemDefine("stringwidth", NewOperator(stringwidth))

	// Graphic state operators
	interpreter.SystemDefine("gsave", NewOperator(gsave))
	interpreter.SystemDefine("grestore", NewOperator(grestore))
	interpreter.SystemDefine("setrgbcolor", NewOperator(setrgbcolor))
	interpreter.SystemDefine("sethsbcolor", NewOperator(sethsbcolor))
	interpreter.SystemDefine("setcmybcolor", NewOperator(setcmybcolor))
	interpreter.SystemDefine("setcmykcolor", NewOperator(setcmybcolor))
	interpreter.SystemDefine("setgray", NewOperator(setgray))
	interpreter.SystemDefine("setdash", NewOperator(setdash))
	interpreter.SystemDefine("setlinejoin", NewOperator(setlinejoin))
	interpreter.SystemDefine("setlinecap", NewOperator(setlinecap))
	interpreter.SystemDefine("setmiterlimit", NewOperator(setmiterlimit))
	interpreter.SystemDefine("setlinewidth", NewOperator(setlinewidth))
	// Graphic state operators device dependent
	interpreter.SystemDefine("setflat", NewOperator(setflat))
	interpreter.SystemDefine("currentflat", NewOperator(currentflat))

	// Coordinate System and Matrix operators
	interpreter.SystemDefine("matrix", NewOperator(transform))
	interpreter.SystemDefine("transform", NewOperator(transform))
	interpreter.SystemDefine("itransform", NewOperator(itransform))
	interpreter.SystemDefine("translate", NewOperator(translate))
	interpreter.SystemDefine("rotate", NewOperator(rotate))
	interpreter.SystemDefine("scale", NewOperator(scale))

	//Path Construction Operators
	interpreter.SystemDefine("newpath", NewOperator(newpath))
	interpreter.SystemDefine("closepath", NewOperator(closepath))
	interpreter.SystemDefine("currentpoint", NewOperator(currentpoint))
	interpreter.SystemDefine("moveto", NewOperator(moveto))
	interpreter.SystemDefine("rmoveto", NewOperator(rmoveto))
	interpreter.SystemDefine("lineto", NewOperator(lineto))
	interpreter.SystemDefine("rlineto", NewOperator(rlineto))
	interpreter.SystemDefine("curveto", NewOperator(curveto))
	interpreter.SystemDefine("rcurveto", NewOperator(rcurveto))
	interpreter.SystemDefine("clippath", NewOperator(clippath))
}
