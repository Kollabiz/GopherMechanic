package Logic

import (
	"ScrapBlueprint/Constants/GateType"
	"ScrapBlueprint/Structs"
)

func GenerateBinaryNumberInput(position Structs.Position, bitCount, inputIndex int) Structs.Body {
	var body = Structs.MakeBody()
	for i := 0; i < bitCount; i++ {
		color := "efefef"
		if i == 0 {
			color = "dfdfdf"
		}
		gate := Structs.MakeLogicGate(color, Structs.Position{position.X + i, position.Y, position.Z}, GateType.OrLogicGate)
		gate.Attributes.SetInt("BitIndex", i)
		gate.Attributes.SetInt("InputIndex", inputIndex)
		body.AddBlock(
			gate,
		)
	}
	return body
}

func GenerateBinaryNumberOutput(position Structs.Position, bitCount, outputIndex int) Structs.Body {
	var body = Structs.MakeBody()
	for i := 0; i < bitCount; i++ {
		color := "cfcfcf"
		if i == 0 {
			color = "afafaf"
		}
		gate := Structs.MakeLogicGate(color, Structs.Position{position.X + i, position.Y, position.Z}, GateType.AndLogicGate)
		gate.Attributes.SetInt("BitIndex", i)
		gate.Attributes.SetInt("InputIndex", outputIndex)
		body.AddBlock(
			gate,
		)
	}
	return body
}

func GenerateBinaryNumberGates(position Structs.Position, bitCount, inputIndex int, gateType GateType.LogicGateType) Structs.Body {
	var body = Structs.MakeBody()
	for i := 0; i < bitCount; i++ {
		color := "cfcfcf"
		if i == 0 {
			color = "afafaf"
		}
		gate := Structs.MakeLogicGate(color, Structs.Position{position.X + i, position.Y, position.Z}, gateType)
		gate.Attributes.SetInt("BitIndex", i)
		gate.Attributes.SetInt("InputIndex", inputIndex)
		body.AddBlock(
			gate,
		)
	}
	return body
}

func GenerateBinaryNumber(position Structs.Position, bitCount, inputIndex int, value uint64) Structs.Body {
	var body = Structs.MakeBody()
	// Dummy gate used to invert NOR gates
	var dummyGate = Structs.MakeLogicGate("000000", Structs.Position{X: position.X, Y: position.Y, Z: position.Z + 1}, GateType.AndLogicGate)
	body.AddBlock(dummyGate)
	for i := 0; i < bitCount; i++ {
		color := "efefef"
		if i == 0 {
			color = "1f1f1f"
		}
		gate := Structs.MakeLogicGate(color, Structs.Position{X: position.X + i, Y: position.Y, Z: position.Z}, GateType.NorLogicGate)
		gate.Attributes.SetInt("BitIndex", i)
		gate.Attributes.SetInt("InputIndex", inputIndex)
		if value&(1<<i) > 0 {
			dummyGate.AddConnection(gate)
		}
		body.AddBlock(gate)
	}
	return body
}

func GenerateBinaryNumberWithOverrideDummyGate(position Structs.Position, bitCount, inputIndex int, value uint64, dummyGate Structs.Block) Structs.Body {
	var body = Structs.MakeBody()
	body.AddBlock(dummyGate)
	for i := 0; i < bitCount; i++ {
		color := "efefef"
		if i == 0 {
			color = "dfdfdf"
		}
		gate := Structs.MakeLogicGate(color, Structs.Position{X: position.X + i, Y: position.Y, Z: position.Z}, GateType.NorLogicGate)
		gate.Attributes.SetInt("BitIndex", i)
		gate.Attributes.SetInt("InputIndex", inputIndex)
		if value&(1<<i) > 0 {
			dummyGate.AddConnection(gate)
		}
		body.AddBlock(gate)
	}
	return body
}
