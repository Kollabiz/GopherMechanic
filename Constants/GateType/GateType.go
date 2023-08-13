package GateType

// LogicGateType describes a logic gate type (wow). Use constants in Constants package to specify a logic gate type
type LogicGateType int

const (
	AndLogicGate  LogicGateType = 0
	OrLogicGate   LogicGateType = 1
	XorLogicGate  LogicGateType = 2
	NandLogicGate LogicGateType = 3
	NorLogicGate  LogicGateType = 4
	NXorLogicGate LogicGateType = 5
)
