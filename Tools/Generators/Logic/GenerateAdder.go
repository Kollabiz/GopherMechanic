package Logic

import (
	"ScrapBlueprint/Constants"
	"ScrapBlueprint/Constants/GateType"
	"ScrapBlueprint/Structs"
)

func GenerateAdderArray(bitWidth int, position Structs.Position) Structs.Body {
	adder := Structs.MakeBody()
	adder.AddBlocks([]Structs.Block{
		Structs.MakeLogicGate(
			"efefef",
			Structs.Position{
				X: position.X,
				Y: position.Y,
				Z: position.Z + 1,
			},
			GateType.OrLogicGate,
		),
		Structs.MakeLogicGate(
			"afafaf",
			Structs.Position{
				X: position.X,
				Y: position.Y + 1,
				Z: position.Z + 1,
			},
			GateType.OrLogicGate,
		),
		Structs.MakeLogicGate(
			Constants.DefaultColors["Logic Gate"],
			position,
			GateType.AndLogicGate,
		),
		Structs.MakeLogicGate(
			Constants.DefaultColors["Logic Gate"],
			Structs.Position{
				X: position.X,
				Y: position.Y + 1,
				Z: position.Z,
			},
			GateType.XorLogicGate,
		),
	})
	adder.Children[0].AddConnection(adder.Children[2])
	adder.Children[0].AddConnection(adder.Children[3])
	adder.Children[1].AddConnection(adder.Children[2])
	adder.Children[1].AddConnection(adder.Children[3])
	adder.Children[0].Attributes.SetInt("InputIndex", 0)
	adder.Children[0].Attributes.SetInt("BitIndex", 0)
	adder.Children[1].Attributes.SetInt("InputIndex", 1)
	adder.Children[1].Attributes.SetInt("BitIndex", 0)
	adder.Children[3].Attributes.SetInt("InputIndex", 2)
	adder.Children[3].Attributes.SetInt("BitIndex", 0)
	adderIndex := 4
	for i := 1; i < bitWidth; i++ {
		gates := []Structs.Block{
			//Inputs
			Structs.MakeLogicGate( // First operand
				"efefef",
				Structs.Position{
					X: position.X + i,
					Y: position.Y,
					Z: position.Z,
				},
				GateType.OrLogicGate,
			),
			Structs.MakeLogicGate( // Second operand
				"afafaf",
				Structs.Position{
					X: position.X + i,
					Y: position.Y,
					Z: position.Z + 1,
				},
				GateType.OrLogicGate,
			),
			// Inner logic
			Structs.MakeLogicGate( // In1, In2 -> X1
				Constants.DefaultColors["Logic Gate"],
				Structs.Position{
					X: position.X + i,
					Y: position.Y + 1,
					Z: position.Z,
				},
				GateType.XorLogicGate,
			),
			Structs.MakeLogicGate( // CIn, X1 -> A1
				Constants.DefaultColors["Logic Gate"],
				Structs.Position{
					X: position.X + i,
					Y: position.Y + 1,
					Z: position.Z + 1,
				},
				GateType.AndLogicGate,
			),
			Structs.MakeLogicGate( // In1, In2 -> C
				Constants.DefaultColors["Logic Gate"],
				Structs.Position{
					X: position.X + i,
					Y: position.Y + 2,
					Z: position.Z,
				},
				GateType.AndLogicGate,
			),
			Structs.MakeLogicGate( // C, A1 -> COut
				Constants.DefaultColors["Logic Gate"],
				Structs.Position{
					X: position.X + i,
					Y: position.Y + 2,
					Z: position.Z + 1,
				},
				GateType.OrLogicGate,
			),
			Structs.MakeLogicGate( // X1, CIn -> Out
				"5f5f5f",
				Structs.Position{
					X: position.X + i,
					Y: position.Y + 2,
					Z: position.Z + 2,
				},
				GateType.XorLogicGate,
			),
		}
		// Adding attributes
		gates[0].Attributes.SetInt("InputIndex", 0)
		gates[0].Attributes.SetInt("BitIndex", i)
		gates[1].Attributes.SetInt("InputIndex", 1)
		gates[1].Attributes.SetInt("BitIndex", i)
		gates[6].Attributes.SetInt("InputIndex", 2)
		gates[6].Attributes.SetInt("BitIndex", i)
		// Connecting gates
		gates[0].AddConnection(gates[2])
		gates[0].AddConnection(gates[4])
		gates[1].AddConnection(gates[2])
		gates[1].AddConnection(gates[4])
		gates[2].AddConnection(gates[3])
		gates[2].AddConnection(gates[6])
		gates[3].AddConnection(gates[5])
		gates[4].AddConnection(gates[5])
		// Connecting with previous adders
		adder.Children[adderIndex-2].AddConnection(gates[3])
		adder.Children[adderIndex-2].AddConnection(gates[6])
		adderIndex += 7
		// Adding to body
		adder.AddBlocks(gates)
	}
	return adder
}
