package Logic

import (
	"ScrapBlueprint/Constants"
	"ScrapBlueprint/Structs"
)

func GenerateBinaryNumberInput(position Structs.Position, bitCount int) Structs.Body {
	var body = Structs.MakeBody()
	for i := 0; i < bitCount; i++ {
		color := "efefef"
		if i == 0 {
			color = "dfdfdf"
		}
		body.AddBlock(
			Structs.MakeLogicGate(color, Structs.Position{position.X + i, position.Y, position.Z}, Constants.OrLogicGate),
		)
	}
	return body
}

func GenerateBinaryNumberOutput(position Structs.Position, bitCount int) Structs.Body {
	var body = Structs.MakeBody()
	for i := 0; i < bitCount; i++ {
		color := "cfcfcf"
		if i == 0 {
			color = "afafaf"
		}
		body.AddBlock(
			Structs.MakeLogicGate(color, Structs.Position{position.X + i, position.Y, position.Z}, Constants.AndLogicGate),
		)
	}
	return body
}

func GenerateBinaryNumberGates(position Structs.Position, bitCount int, gateType Constants.LogicGateType) Structs.Body {
	var body = Structs.MakeBody()
	for i := 0; i < bitCount; i++ {
		color := "cfcfcf"
		if i == 0 {
			color = "afafaf"
		}
		body.AddBlock(
			Structs.MakeLogicGate(color, Structs.Position{position.X + i, position.Y, position.Z}, gateType),
		)
	}
	return body
}

func GenerateBinaryNumber(position Structs.Position, bitCount int, value uint64) Structs.Body {
	var body = Structs.MakeBody()
	// Dummy gate used to invert NOR gates
	var dummyGate = Structs.MakeLogicGate("000000", Structs.Position{X: position.X, Y: position.Y, Z: position.Z + 1}, Constants.AndLogicGate)
	body.AddBlock(dummyGate)
	for i := 0; i < bitCount; i++ {
		color := "efefef"
		if i == 0 {
			color = "1f1f1f"
		}
		gate := Structs.MakeLogicGate(color, Structs.Position{X: position.X + i, Y: position.Y, Z: position.Z}, Constants.NorLogicGate)
		if value&(1<<i) > 0 {
			dummyGate.AddConnection(gate)
		}
		body.AddBlock(gate)
	}
	return body
}

func GenerateBinaryNumberWithOverrideDummyGate(position Structs.Position, bitCount int, value uint64, dummyGate Structs.Block) Structs.Body {
	var body = Structs.MakeBody()
	body.AddBlock(dummyGate)
	for i := 0; i < bitCount; i++ {
		color := "efefef"
		if i == 0 {
			color = "dfdfdf"
		}
		gate := Structs.MakeLogicGate(color, Structs.Position{X: position.X + i, Y: position.Y, Z: position.Z}, Constants.NorLogicGate)
		if value&(1<<i) > 0 {
			dummyGate.AddConnection(gate)
		}
		body.AddBlock(gate)
	}
	return body
}
