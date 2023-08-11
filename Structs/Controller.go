package Structs

// Controller describes block "controller". Actually it contains data about block's connections
type Controller struct {
	// Is Block active
	Active bool `json:"active"`
	// Ingoing connections to the block
	Controllers []ControllerId `json:"controllers"`
	// Identifier for outgoing connections
	Id int `json:"id"`
	// Block mode. It describes various information, like logic gate type, light power and so on
	Mode int `json:"mode"`
}
