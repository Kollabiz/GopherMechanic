package Structs

// ControllerId is identifier for ingoing connections. Used only during serialisation
type ControllerId struct {
	Id int `json:"id"`
}

func GetId(block Block) ControllerId {
	return ControllerId{block.Controller.Id}
}
