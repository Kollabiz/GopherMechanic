package Constants

var currentId = 0

func NextID() int {
	currentId++
	return currentId
}
