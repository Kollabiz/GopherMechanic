package Logic

import (
	"ScrapBlueprint/Constants"
	"ScrapBlueprint/Constants/GateType"
	"ScrapBlueprint/Structs"
	"math"
)

func GenerateMultiplexor(indexBitWidth int, position Structs.Position) Structs.Body {
	inputNumber := GenerateBinaryNumberInput(Structs.Position{X: position.X, Y: position.Y + 1, Z: position.Z}, indexBitWidth, 0)
	inputInverter := GenerateBinaryNumberGates(position, indexBitWidth, -1, GateType.NorLogicGate)
	for i := 0; i < indexBitWidth; i++ {
		inputNumber.Children[i].AddConnection(inputInverter.Children[i])
	}
	var muxGatesCount = int(math.Pow(2, float64(indexBitWidth)))
	var muxGates Structs.Body
	for i := 0; i < muxGatesCount; i++ {
		gate := Structs.MakeLogicGate(Constants.DefaultColors["Logic Gate"], Structs.Position{
			X: position.X,
			Y: position.Y,
			Z: position.Z + i + 1,
		}, GateType.AndLogicGate)
		gate.Attributes.SetInt("BitIndex", i)
		gate.Attributes.SetInt("InputIndex", 1)
		muxGates.AddBlock(gate)
	}
	stride := muxGatesCount / 2
	for i := 0; i < indexBitWidth; i++ {
		osc := 0
		for j := 0; j < muxGatesCount; j += stride {
			var inputGate Structs.Block
			if osc%2 != 0 {
				inputGate = inputInverter.Children[i]
			} else {
				inputGate = inputNumber.Children[i]
			}
			for k := j; k < j+stride; k++ {
				inputGate.AddConnection(muxGates.Children[k])
			}
			osc++
		}
		stride /= 2
	}
	inputNumber.Join(&inputInverter)
	inputNumber.Join(&muxGates)
	return inputNumber
}
