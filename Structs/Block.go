package Structs

import (
	"ScrapBlueprint/Constants"
	"ScrapBlueprint/Constants/GateType"
	"ScrapBlueprint/Constants/PartTier"
	"fmt"
)

// Block describes scrap mechanic tile
type Block struct {
	// Bounds are used to make block volumes easily
	Bounds Position `json:"bounds"`
	// Block color in HEX
	Color string `json:"color"`
	// Block controller (useful only for blocks with connection nodes)
	Controller Controller `json:"controller"`
	// Block position
	Pos Position `json:"pos"`
	// Block ShapeID, a.k.a. what block this is
	ShapeId string `json:"shapeId"`
	// Block X-axis rotation in 1/4 turns
	XAxis int `json:"xaxis"`
	// Block Z-axis rotation in 1/4 turns
	ZAxis int `json:"zaxis"`
	// Attributes of block used only during generation. These are not saved into the blueprint
	Attributes *CustomAttributes
}

// MakeBlock creates a block with given shapeId, color and position
func MakeBlock(shapeId, color string, position Position) Block {
	controller :=
		Controller(&BasicController{
			Active:      false,
			Controllers: nil,
			Id:          Constants.NextID(),
			Mode:        0,
		})
	return Block{
		Bounds:     Position{1, 1, 1},
		Color:      color,
		Controller: controller,
		Pos:        position,
		ShapeId:    shapeId,
		XAxis:      0,
		ZAxis:      0,
		Attributes: NewCustomAttributes(),
	}
}

// MakeBlockWithRotation creates a block with given shapeId, color, position and rotation, specified by two integers for x and z axes respectively. Rotation is specified in quarters of turn (90 degrees step)
func MakeBlockWithRotation(shapeId, color string, position Position, xRotation, zRotation int) Block {
	controller :=
		Controller(&BasicController{
			Active:      false,
			Controllers: nil,
			Id:          Constants.NextID(),
			Mode:        0,
		})
	return Block{
		Bounds:     Position{1, 1, 1},
		Color:      color,
		Controller: controller,
		Pos:        position,
		ShapeId:    shapeId,
		XAxis:      xRotation,
		ZAxis:      zRotation,
		Attributes: NewCustomAttributes(),
	}
}

// MakeLogicGate creates a logic gate with given color, position and gate type
func MakeLogicGate(color string, position Position, gateType GateType.LogicGateType) Block {
	controller :=
		Controller(&BasicController{
			Active:      false,
			Controllers: nil,
			Id:          Constants.NextID(),
			Mode:        int(gateType),
		})
	return Block{
		Bounds:     Position{1, 1, 1},
		Color:      color,
		Controller: controller,
		Pos:        position,
		ShapeId:    Constants.Blocks["Logic Gate"],
		XAxis:      0,
		ZAxis:      0,
		Attributes: NewCustomAttributes(),
	}
}

func MakeSensor(color string, position Position, sensorTier PartTier.PartTier, buttonMode bool, colorMode bool, detectColor string, sensorRange int, audioEnabled bool) Block {
	controller :=
		Controller(&SensorController{
			AudioEnabled: audioEnabled,
			ButtonMode:   buttonMode,
			Color:        detectColor,
			ColorMode:    colorMode,
			Active:       false,
			Controllers:  nil,
			Id:           Constants.NextID(),
			Mode:         0,
			Range:        sensorRange,
		})
	return Block{
		Bounds:     Position{1, 1, 1},
		Color:      color,
		Controller: controller,
		Pos:        position,
		ShapeId:    Constants.Blocks[fmt.Sprintf("Sensor %d", int(sensorTier))],
		XAxis:      0,
		ZAxis:      0,
		Attributes: NewCustomAttributes(),
	}
}

/*
	MakeBlockWithMode creates a block with given shapeId, color, position and controller mode

For example, you can create an AND logic gate by calling this function like this:

MakeBlockWithMode(Constants.Blocks["Logic Gate"], "DF7F01", Position{0, 0, 0}, 0)

# But you should use MakeLogicGate instead to create logic gates

(Logic gate mode note: 0 - AND, 1 - OR, 2 - XOR, 3 - NAND, 4 - NOR, 5 - NXOR)
*/
func MakeBlockWithMode(shapeId, color string, position Position, mode int) Block {
	controller :=
		Controller(&BasicController{
			Active:      false,
			Controllers: nil,
			Id:          Constants.NextID(),
			Mode:        mode,
		})
	return Block{
		Bounds:     Position{1, 1, 1},
		Color:      color,
		Controller: controller,
		Pos:        position,
		ShapeId:    shapeId,
		XAxis:      0,
		ZAxis:      0,
		Attributes: NewCustomAttributes(),
	}
}

// MakeBlockVolume creates a volume of blocks with specified bounds
func MakeBlockVolume(shapeId, color string, position Position, bounds Position) Block {
	controller :=
		Controller(&BasicController{
			Active:      false,
			Controllers: nil,
			Id:          Constants.NextID(),
			Mode:        0,
		})
	return Block{
		Bounds:     bounds,
		Color:      color,
		Controller: controller,
		Pos:        position,
		ShapeId:    shapeId,
		XAxis:      0,
		ZAxis:      0,
		Attributes: NewCustomAttributes(),
	}
}

// AddConnection adds a connection between current block and given block
func (block *Block) AddConnection(child Block) {
	block.Controller.AddConnection(child.Controller.GetId())
}

// AddConnectionByID adds a connection between current block and given controller identifier
//
// You should be careful while using this method, you may break the blueprint if you specify non-existing controller id
func (block *Block) AddConnectionByID(id int) {
	block.Controller.AddConnection(id)
}

// RemoveConnectionByID removes a connection between current block and given controller identifier
//
// If there was a connection with given ID true is returned, otherwise function returns false
func (block *Block) RemoveConnectionByID(id int) bool {
	var index int
	found := false
	connections := block.Controller.GetConnections()
	for index = 0; index < len(connections); index++ {
		if connections[index].Id == id {
			found = true
			break
		}
	}
	if !found {
		return false
	}

	block.Controller.RemoveConnection(connections[index].Id)
	return true
}

// RemoveConnection removes connection between current block and given block
//
// If there was a connection with given block true is returned, otherwise function returns false
func (block *Block) RemoveConnection(child Block) bool {
	return block.RemoveConnectionByID(child.Controller.GetId())
}
