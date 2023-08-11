package Structs

// Blueprint describes a Scrap mechanic blueprint
type Blueprint struct {
	// All the bodies in this blueprint
	Bodies []Body `json:"bodies"`
	// Blueprint version
	Version int `json:"version"`
}

func NewBlueprint() *Blueprint {
	bp := new(Blueprint)
	bp.Version = 3
	return bp
}

func (bp *Blueprint) AddBody(body Body) {
	bp.Bodies = append(bp.Bodies, body)
}
