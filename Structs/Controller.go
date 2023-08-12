package Structs

// Controller describes a controller for blocks. It describes block's connections, ID and some additional stuff (range of a light, for example)
type Controller interface {
	// AddConnection adds a connection to given controller id
	AddConnection(id int)
	// RemoveConnection removes a connection to given controller id
	RemoveConnection(id int)
	// HasConnection returns true if there is a connection to given id, otherwise it returns false
	HasConnection(id int) bool
	SetActive(active bool)
	GetActive() bool
	GetId() int
	SetMode(mode int)
	GetMode() int
	// GetConnections returns all the connections of controller
	GetConnections() []ControllerId
}

// BasicController describes a controller for basic interactive blocks
type BasicController struct {
	// Is Block active
	Active bool `json:"active"`
	// Ingoing connections to the block
	Controllers []ControllerId `json:"controllers"`
	// Identifier for outgoing connections
	Id int `json:"id"`
	// Block mode. It describes various information, like logic gate type, light power and so on
	Mode int `json:"mode"`
}

// AddConnection adds a connection to given controller id
func (controller *BasicController) AddConnection(id int) {
	controller.Controllers = append(controller.Controllers, ControllerId{id})
}

// RemoveConnection removes a connection to given controller id
func (controller *BasicController) RemoveConnection(id int) {
	for i := 0; i < len(controller.Controllers); i++ {
		if controller.Controllers[i].Id == id {
			controller.Controllers = append(controller.Controllers[:i], controller.Controllers[i+1:]...)
			return
		}
	}
}

// HasConnection returns true if there is a connection to given id, otherwise it returns false
func (controller *BasicController) HasConnection(id int) bool {
	for i := 0; i < len(controller.Controllers); i++ {
		if controller.Controllers[i].Id == id {
			return true
		}
	}
	return false
}

func (controller *BasicController) SetActive(active bool) {
	controller.Active = active
}

func (controller *BasicController) GetActive() bool {
	return controller.Active
}

func (controller *BasicController) GetId() int {
	return controller.Id
}

func (controller *BasicController) SetMode(mode int) {
	controller.Mode = mode
}

func (controller *BasicController) GetMode() int {
	return controller.Mode
}

// GetConnections returns all the connections of controller
func (controller *BasicController) GetConnections() []ControllerId {
	return controller.Controllers
}

//		#######################
//		# Generic controllers #
//		#######################

type SensorController struct {
	// AudioEnabled determines if audio is enabled for the sensor
	AudioEnabled bool `json:"audioEnabled"`
	// ButtonMode determines if sensor should behave like a button
	ButtonMode bool `json:"buttonMode"`
	// Color determines the color for sensor to detect
	Color string `json:"color"`
	// ColorMode determines if sensor should see only blocks with certain color
	ColorMode bool `json:"colorMode"`
	// Is Block active
	Active bool `json:"active"`
	// Ingoing connections to the block
	Controllers []ControllerId `json:"controllers"`
	// Identifier for outgoing connections
	Id int `json:"id"`
	// Block mode. It describes various information, like logic gate type, light power and so on
	Mode int `json:"mode"`
	// Range determines sensor's range
	Range int `json:"range"`
}

func (controller *SensorController) AddConnection(id int) {
	controller.Controllers = append(controller.Controllers, ControllerId{id})
}

func (controller *SensorController) RemoveConnection(id int) {
	for i := 0; i < len(controller.Controllers); i++ {
		if controller.Controllers[i].Id == id {
			controller.Controllers = append(controller.Controllers[:i], controller.Controllers[i+1:]...)
			return
		}
	}
}

// HasConnection returns true if there is a connection to given id, otherwise it returns false
func (controller *SensorController) HasConnection(id int) bool {
	for i := 0; i < len(controller.Controllers); i++ {
		if controller.Controllers[i].Id == id {
			return true
		}
	}
	return false
}

func (controller *SensorController) SetActive(active bool) {
	controller.Active = active
}

func (controller *SensorController) GetActive() bool {
	return controller.Active
}

func (controller *SensorController) GetId() int {
	return controller.Id
}

func (controller *SensorController) SetMode(mode int) {
	controller.Mode = mode
}

func (controller *SensorController) GetMode() int {
	return controller.Mode
}

// GetConnections returns all the connections of controller
func (controller *SensorController) GetConnections() []ControllerId {
	return controller.Controllers
}

// TimerController describes a controller for timers
type TimerController struct {
	// Is Block active
	Active bool `json:"active"`
	// Ingoing connections to the block
	Controllers []ControllerId `json:"controllers"`
	// Identifier for outgoing connections
	Id int `json:"id"`
	// Block mode. It describes various information, like logic gate type, light power and so on
	Mode int `json:"mode"`
	// Time set in Seconds
	Seconds int `json:"seconds"`
	// Time set in Ticks
	Ticks int `json:"ticks"`
}

// AddConnection adds a connection to given controller id
func (controller *TimerController) AddConnection(id int) {
	controller.Controllers = append(controller.Controllers, ControllerId{id})
}

// RemoveConnection removes a connection to given controller id
func (controller *TimerController) RemoveConnection(id int) {
	for i := 0; i < len(controller.Controllers); i++ {
		if controller.Controllers[i].Id == id {
			controller.Controllers = append(controller.Controllers[:i], controller.Controllers[i+1:]...)
			return
		}
	}
}

// HasConnection returns true if there is a connection to given id, otherwise it returns false
func (controller *TimerController) HasConnection(id int) bool {
	for i := 0; i < len(controller.Controllers); i++ {
		if controller.Controllers[i].Id == id {
			return true
		}
	}
	return false
}

func (controller *TimerController) SetActive(active bool) {
	controller.Active = active
}

func (controller *TimerController) GetActive() bool {
	return controller.Active
}

func (controller *TimerController) GetId() int {
	return controller.Id
}

func (controller *TimerController) SetMode(mode int) {
	controller.Mode = mode
}

func (controller *TimerController) GetMode() int {
	return controller.Mode
}

// GetConnections returns all the connections of controller
func (controller *TimerController) GetConnections() []ControllerId {
	return controller.Controllers
}
