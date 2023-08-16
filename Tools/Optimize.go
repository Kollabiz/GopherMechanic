package Tools

import (
	"ScrapBlueprint/Structs"
	"math/rand"
)

// A typical quick sort by distance to corner with minimal X Y Z coordinates
func sortByDistanceToCorner(blocks []Structs.Block) []Structs.Block {
	if len(blocks) == 2 {
		if blocks[1].Pos.X < blocks[0].Pos.X || blocks[1].Pos.Y < blocks[1].Pos.Y || blocks[1].Pos.Z < blocks[1].Pos.Z {
			return []Structs.Block{blocks[1], blocks[0]}
		}
		return blocks
	} else if len(blocks) < 2 {
		return blocks
	}
	baseElem := blocks[rand.Int()%len(blocks)]
	var lower []Structs.Block
	var greater []Structs.Block
	for i := 0; i < len(blocks); i++ {
		if blocks[i].Controller.GetId() == baseElem.Controller.GetId() {
			continue
		}
		if blocks[i].Pos.X < baseElem.Pos.X || blocks[i].Pos.Y < baseElem.Pos.Y || blocks[i].Pos.Z < baseElem.Pos.Z {
			lower = append(lower, blocks[i])
		} else {
			greater = append(greater, blocks[i])
		}
	}
	lowerSorted := sortByDistanceToCorner(lower)
	greaterSorted := sortByDistanceToCorner(greater)
	return append(append(lowerSorted, baseElem), greaterSorted...)
}

// getVolumeBlocks returns all the blocks that are inside the specified volume
func getVolumeBlocks(volumeStart Structs.Position, volumeSize Structs.Position, blocks []Structs.Block) []Structs.Block {
	var b []Structs.Block
	for i := 0; i < len(blocks); i++ {
		block := blocks[i]
		if block.Pos.X < volumeStart.X+volumeSize.X &&
			block.Pos.Y < volumeStart.Y+volumeSize.Y &&
			block.Pos.Z < volumeStart.Z+volumeSize.Z &&
			block.Pos.X >= volumeStart.X && block.Pos.Y >= volumeStart.Y && block.Pos.Z >= volumeStart.Z {
			b = append(b, block)
		}
	}
	return b
}

// allBlocksAreJoinable returns true if all passed blocks can be joined into a single volume
func allBlocksAreJoinable(blocks []Structs.Block) bool {
	for i := 1; i < len(blocks); i++ {
		if blocks[i].ShapeId != blocks[i-1].ShapeId ||
			blocks[i].Color != blocks[i-1].Color ||
			blocks[i].ZAxis != blocks[i-1].ZAxis ||
			blocks[i].XAxis != blocks[i-1].XAxis {
			return false
		}
	}
	return true
}

// checkVolume checks the volume integrity
func checkVolume(volumeSize Structs.Position, blocks []Structs.Block) bool {
	return len(blocks) == volumeSize.X*volumeSize.Y*volumeSize.Z && allBlocksAreJoinable(blocks)
}

func removeBlockFromArray(arr []Structs.Block, elem Structs.Block) []Structs.Block {
	for i := 0; i < len(arr); i++ {
		if arr[i].Controller.GetId() == elem.Controller.GetId() {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

// Optimize optimizes body by joining blocks into volumes if possible
// You should NEVER use this function on bodies that contain parts alongside with blocks, because this may break your blueprint
//
// Note: Only blocks with same ShapeId, color and rotation are joined
func Optimize(body *Structs.Body) {
	sorted := sortByDistanceToCorner(body.Children)
	var optimized []Structs.Block
	for len(sorted) > 0 {
		startBlock := sorted[0]
		bounds := Structs.Position{1, 1, 1}
		volume := []Structs.Block{startBlock}
		// X axis
		for checkVolume(bounds, volume) {
			bounds.X++
			volume = getVolumeBlocks(startBlock.Pos, bounds, sorted)
		}
		if bounds.X > 1 {
			bounds.X--
		}
		volume = getVolumeBlocks(startBlock.Pos, bounds, sorted)
		// Z axis
		for checkVolume(bounds, volume) {
			bounds.Z++
			volume = getVolumeBlocks(startBlock.Pos, bounds, sorted)
		}
		if bounds.Z > 1 {
			bounds.Z--
		}
		volume = getVolumeBlocks(startBlock.Pos, bounds, sorted)
		// Y axis
		for checkVolume(bounds, volume) {
			bounds.Y++
			volume = getVolumeBlocks(startBlock.Pos, bounds, sorted)
		}
		if bounds.Y > 1 {
			bounds.Y--
		}
		volume = getVolumeBlocks(startBlock.Pos, bounds, sorted)
		startBlock.Bounds = bounds
		optimized = append(optimized, startBlock)
		for i := 0; i < len(volume); i++ {
			sorted = removeBlockFromArray(sorted, volume[i])
		}
	}
	body.Children = optimized
}
