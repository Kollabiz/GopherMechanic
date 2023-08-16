package Generators

import (
	"ScrapBlueprint/Structs"
	"math"
)

func GenerateLine(startPoint, endPoint Structs.Position, block, color string, lineSize int) Structs.Body {
	dx := float64(endPoint.X - startPoint.X)
	dy := float64(endPoint.Y - startPoint.Y)
	dz := float64(endPoint.Z - startPoint.Z)
	var step float64
	if math.Abs(float64(dx)) >= math.Abs(float64(dy)) {
		if math.Abs(float64(dz)) >= math.Abs(float64(dx)) {
			step = float64(dz)
		} else {
			step = float64(dx)
		}
	} else {
		if math.Abs(float64(dz)) >= math.Abs(float64(dy)) {
			step = float64(dz)
		} else {
			step = float64(dy)
		}
	}
	dx /= step
	dy /= step
	dz /= step
	x := float64(startPoint.X)
	y := float64(startPoint.Y)
	z := float64(startPoint.Z)
	body := Structs.MakeBody()
	for i := float64(0); i < step; i++ {
		// Generating hollow spheres to reduce blueprint size
		sphere := GenerateSphereNaiveHollow(Structs.Position{int(x), int(y), int(z)}, lineSize, 1, block, color)
		body.Join(&sphere)
		x += dx
		y += dy
		z += dz
		i++
	}
	return body
}
