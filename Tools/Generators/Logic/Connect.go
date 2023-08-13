package Logic

import (
	"ScrapBlueprint/Structs"
)

func sortByBitIndex(blocks []Structs.Block) []Structs.Block {
	var sorted = make([]Structs.Block, len(blocks))
	for i := 0; i < len(blocks); i++ {
		block := blocks[i]
		if !block.Attributes.AttributeExists("BitIndex") {
			panic("trying to connect invalid numbers (no BitIndex attribute)")
		}
		bitIndex := block.Attributes.GetInt("BitIndex")
		sorted[bitIndex] = block
	}
	return sorted
}

func ConnectBinary(body1, body2 Structs.Body, index1, index2 int) {
	var firstBodyGates []Structs.Block
	var secondBodyGates []Structs.Block
	for i := 0; i < len(body1.Children); i++ {
		block := body1.Children[i]
		if block.Attributes.AttributeExists("InputIndex") && block.Attributes.GetInt("InputIndex") == index1 {
			firstBodyGates = append(firstBodyGates, block)
		}
	}
	firstBodyGates = sortByBitIndex(firstBodyGates)

	for i := 0; i < len(body2.Children); i++ {
		block := body2.Children[i]
		if block.Attributes.AttributeExists("InputIndex") && block.Attributes.GetInt("InputIndex") == index2 {
			secondBodyGates = append(secondBodyGates, block)
		}
	}
	secondBodyGates = sortByBitIndex(secondBodyGates)

	minLength := len(firstBodyGates)
	if len(secondBodyGates) < minLength {
		minLength = len(secondBodyGates)
	}

	for i := 0; i < minLength; i++ {
		firstBodyGates[i].AddConnection(secondBodyGates[i])
	}
}
