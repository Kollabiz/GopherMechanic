package Structs

// Body describes a physic body. Physic bodies are generated while building on bearings, pistons or suspension springs
type Body struct {
	// All the blocks that make this Body up
	Children []Block `json:"childs"` // I have no slightest idea why Axolot Games named this "childs" instead of "children" in json
}

// MakeBody creates an empty body
func MakeBody() Body {
	return Body{}
}

// AddBlock adds given block to the body
func (body *Body) AddBlock(block Block) {
	body.Children = append(body.Children, block)
}

// AddBlocks adds a given slice of blocks to the body (blocks that already exist within body are added twice)
func (body *Body) AddBlocks(blocks []Block) {
	body.Children = append(body.Children, blocks...)
}

// RemoveBlock removes given block from the body
//
// If there was a block to remove, true is returned, otherwise function returns false
func (body *Body) RemoveBlock(block Block) bool {
	for i := 0; i < len(body.Children); i++ {
		if body.Children[i].Controller.Id == block.Controller.Id {
			body.Children = append(body.Children[:i], body.Children[i+1:]...)
			return true
		}
	}
	return false
}

// RemoveBlocks removes a given slice of blocks from the body
//
// If any block was removed true is returned, otherwise function returns false
func (body *Body) RemoveBlocks(blocks []Block) bool {
	removedAny := false
	for i := 0; i < len(blocks); i++ {
		if body.RemoveBlock(blocks[i]) {
			removedAny = true
		}
	}

	return removedAny
}
