package Structs

import "ScrapBlueprint/Constants"

// Block describes scrap mechanic tile
type Block struct {
	// Bounds are used to make block volumes easily
	Bounds Position `json:"bounds"`
	// Block color in HEX
	Color string `json:"color"`
	// Block controller (useful only for blocks with connection nodes)
	Controller *Controller `json:"controller"`
	// Block position
	Pos Position `json:"pos"`
	// Block ShapeID, a.k.a. what block this is
	ShapeId string `json:"shapeId"`
	// Block X-axis rotation in 1/4 turns
	XAxis int `json:"xaxis"`
	// Block Z-axis rotation in 1/4 turns
	ZAxis int `json:"zaxis"`
}

// MakeBlock creates a block with given shapeId, color and position
func MakeBlock(shapeId, color string, position Position) Block {
	return Block{
		Bounds: Position{1, 1, 1},
		Color:  color,
		Controller: &Controller{
			Active:      false,
			Controllers: nil,
			Id:          Constants.NextID(),
			Mode:        0,
		},
		Pos:     position,
		ShapeId: shapeId,
		XAxis:   0,
		ZAxis:   0,
	}
}

// MakeBlockWithRotation creates a block with given shapeId, color, position and rotation, specified by two integers for x and z axes respectively. Rotation is specified in quarters of turn (90 degrees step)
func MakeBlockWithRotation(shapeId, color string, position Position, xRotation, zRotation int) Block {
	return Block{
		Bounds: Position{1, 1, 1},
		Color:  color,
		Controller: &Controller{
			Active:      false,
			Controllers: nil,
			Id:          Constants.NextID(),
			Mode:        0,
		},
		Pos:     position,
		ShapeId: shapeId,
		XAxis:   xRotation,
		ZAxis:   zRotation,
	}
}

// MakeLogicGate creates a logic gate with given color, position and gate type
func MakeLogicGate(color string, position Position, gateType Constants.LogicGateType) Block {
	return Block{
		Bounds: Position{1, 1, 1},
		Color:  color,
		Controller: &Controller{
			Active:      false,
			Controllers: nil,
			Id:          Constants.NextID(),
			Mode:        int(gateType),
		},
		Pos:     position,
		ShapeId: Constants.Blocks["Logic Gate"],
		XAxis:   0,
		ZAxis:   0,
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
	return Block{
		Bounds: Position{1, 1, 1},
		Color:  color,
		Controller: &Controller{
			Active:      false,
			Controllers: nil,
			Id:          Constants.NextID(),
			Mode:        mode,
		},
		Pos:     position,
		ShapeId: shapeId,
		XAxis:   0,
		ZAxis:   0,
	}
}

// MakeBlockVolume creates a volume of blocks with specified bounds
func MakeBlockVolume(shapeId, color string, position Position, bounds Position) Block {
	return Block{
		Bounds: bounds,
		Color:  color,
		Controller: &Controller{
			Active:      false,
			Controllers: nil,
			Id:          Constants.NextID(),
			Mode:        0,
		},
		Pos:     position,
		ShapeId: shapeId,
		XAxis:   0,
		ZAxis:   0,
	}
}

// AddConnection adds a connection between current block and given block
func (block *Block) AddConnection(child Block) {
	block.Controller.Controllers = append(block.Controller.Controllers, GetId(child))
}

// AddConnectionByID adds a connection between current block and given controller identifier
//
// You should be careful while using this method, you may break the blueprint if you specify non-existing controller id
func (block *Block) AddConnectionByID(id int) {
	block.Controller.Controllers = append(block.Controller.Controllers, ControllerId{id})
}

// RemoveConnectionByID removes a connection between current block and given controller identifier
//
// If there was a connection with given ID true is returned, otherwise function returns false
func (block *Block) RemoveConnectionByID(id int) bool {
	var index int
	found := false
	for index = 0; index < len(block.Controller.Controllers); index++ {
		if block.Controller.Controllers[index].Id == id {
			found = true
			break
		}
	}
	if !found {
		return false
	}

	block.Controller.Controllers = append(block.Controller.Controllers[:index], block.Controller.Controllers[index+1:]...)
	return true
}

// RemoveConnection removes connection between current block and given block
//
// If there was a connection with given block true is returned, otherwise function returns false
func (block *Block) RemoveConnection(child Block) bool {
	return block.RemoveConnectionByID(child.Controller.Id)
}
