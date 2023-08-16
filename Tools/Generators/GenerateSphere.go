package Generators

import "ScrapBlueprint/Structs"

func GenerateSphereNaive(position Structs.Position, radius int, block, color string) Structs.Body {
	startPoint := Structs.Position{
		X: position.X - radius,
		Y: position.Y - radius,
		Z: position.Z - radius,
	}
	endPoint := Structs.Position{
		X: position.X + radius + 1,
		Y: position.Y + radius + 1,
		Z: position.Z + radius + 1,
	}
	body := Structs.MakeBody()

	for x := startPoint.X; x < endPoint.X; x++ {
		for y := startPoint.Y; y < endPoint.Y; y++ {
			for z := startPoint.Z; z < endPoint.Z; z++ {
				dist := position.DistanceTo(Structs.Position{x, y, z})
				if dist <= float64(radius) {
					tile := Structs.MakeBlock(block, color, Structs.Position{x, y, z})
					body.AddBlock(tile)
				}
			}
		}
	}

	return body
}

func GenerateSphereNaiveHollow(position Structs.Position, radius, borderWidth int, block, color string) Structs.Body {
	startPoint := Structs.Position{
		X: position.X - radius,
		Y: position.Y - radius,
		Z: position.Z - radius,
	}
	endPoint := Structs.Position{
		X: position.X + radius + 1,
		Y: position.Y + radius + 1,
		Z: position.Z + radius + 1,
	}
	body := Structs.MakeBody()

	for x := startPoint.X; x < endPoint.X; x++ {
		for y := startPoint.Y; y < endPoint.Y; y++ {
			for z := startPoint.Z; z < endPoint.Z; z++ {
				dist := position.DistanceTo(Structs.Position{x, y, z})
				if dist <= float64(radius) && dist >= float64(radius-borderWidth) {
					tile := Structs.MakeBlock(block, color, Structs.Position{x, y, z})
					body.AddBlock(tile)
				}
			}
		}
	}

	return body
}
