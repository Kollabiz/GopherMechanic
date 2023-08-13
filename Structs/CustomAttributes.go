package Structs

import (
	"ScrapBlueprint/Constants/AttribType"
	"fmt"
	"reflect"
)

// CustomAttributes contains attributes created during generation. These attributes are NOT saved to blueprint and used only during generation process
type CustomAttributes struct {
	// A map of attributes. They are extracted with reflection
	attributes map[string]any
}

func NewCustomAttributes() *CustomAttributes {
	attrib := new(CustomAttributes)
	attrib.attributes = make(map[string]any)
	return attrib
}

//		###########
//		# Getters #
//		###########

// GetInt return a value of an integer attribute, if it exists
func (attrib *CustomAttributes) GetInt(attribName string) int {
	attribInterface, ok := attrib.attributes[attribName]
	if !ok {
		panic(fmt.Sprintf("undefined attribute: %s", attribName))
	}

	attribVal := reflect.ValueOf(attribInterface)
	attribType := reflect.TypeOf(attribInterface)
	if attribType.Kind() == reflect.Int {
		return int(attribVal.Int())
	}
	panic(fmt.Sprintf("attribute %s is not an int", attribName))
}

// GetUInt returns a value of an unsigned integer attribute if it exists
func (attrib *CustomAttributes) GetUInt(attribName string) uint {
	attribInterface, ok := attrib.attributes[attribName]
	if !ok {
		panic(fmt.Sprintf("undefined attribute: %s", attribName))
	}

	attribVal := reflect.ValueOf(attribInterface)
	attribType := reflect.TypeOf(attribInterface)
	if attribType.Kind() == reflect.Uint {
		return uint(attribVal.Uint())
	}
	panic(fmt.Sprintf("attribute %s is not an uint", attribName))
}

// GetFloat returns a value of a float attribute if it exists
func (attrib *CustomAttributes) GetFloat(attribName string) float64 {
	attribInterface, ok := attrib.attributes[attribName]
	if !ok {
		panic(fmt.Sprintf("undefined attribute: %s", attribName))
	}

	attribVal := reflect.ValueOf(attribInterface)
	attribType := reflect.TypeOf(attribInterface)
	if attribType.Kind() == reflect.Float64 {
		return attribVal.Float()
	}
	panic(fmt.Sprintf("attribute %s is not a float64", attribName))
}

// GetString returns a value of a string attribute if it exists
func (attrib *CustomAttributes) GetString(attribName string) string {
	attribInterface, ok := attrib.attributes[attribName]
	if !ok {
		panic(fmt.Sprintf("undefined attribute: %s", attribName))
	}

	attribVal := reflect.ValueOf(attribInterface)
	attribType := reflect.TypeOf(attribInterface)
	if attribType.Kind() == reflect.String {
		return attribVal.String()
	}
	panic(fmt.Sprintf("attribute %s is not a string", attribName))
}

// GetBool returns a value of a boolean attribute if it exists
func (attrib *CustomAttributes) GetBool(attribName string) bool {
	attribInterface, ok := attrib.attributes[attribName]
	if !ok {
		panic(fmt.Sprintf("undefined attribute: %s", attribName))
	}

	attribVal := reflect.ValueOf(attribInterface)
	attribType := reflect.TypeOf(attribInterface)
	if attribType.Kind() == reflect.Bool {
		return attribVal.Bool()
	}
	panic(fmt.Sprintf("attribute %s is not a bool", attribName))
}

// AttributeExists checks if attribute with given name exists
func (attrib *CustomAttributes) AttributeExists(attribName string) bool {
	_, ok := attrib.attributes[attribName]
	return ok
}

// IsOfType checks if attribute with given name is of certain type. If there is no such attribute it panics
func (attrib *CustomAttributes) IsOfType(attribName string, attribType AttribType.AttribType) bool {
	attribVal, ok := attrib.attributes[attribName]
	if !ok {
		panic(fmt.Sprintf("undefined attribute: %s", attribName))
	}

	return AttribType.AttribTypeFromKind(reflect.TypeOf(attribVal).Kind()) == attribType
}

//		###########
//		# Setters #
//		###########

// SetInt sets a value of specified integer attribute or creates new one
func (attrib *CustomAttributes) SetInt(attribName string, value int) {
	attribInterface, ok := attrib.attributes[attribName]
	if !ok {
		attrib.attributes[attribName] = value
		return
	}

	attribType := reflect.TypeOf(attribInterface)
	if attribType.Kind() == reflect.Int {
		attrib.attributes[attribName] = value
	} else {
		panic(fmt.Sprintf("attribute %s is not an int", attribName))
	}
}

// SetUInt sets a value of specified unsigned integer attribute or creates a new one
func (attrib *CustomAttributes) SetUInt(attribName string, value uint) {
	attribInterface, ok := attrib.attributes[attribName]
	if !ok {
		attrib.attributes[attribName] = value
		return
	}

	attribType := reflect.TypeOf(attribInterface)
	if attribType.Kind() == reflect.Uint {
		attrib.attributes[attribName] = value
	} else {
		panic(fmt.Sprintf("attribute %s is not an uint", attribName))
	}
}

// SetFloat sets a value of specified float attribute or creates a new one
func (attrib *CustomAttributes) SetFloat(attribName string, value float64) {
	attribInterface, ok := attrib.attributes[attribName]
	if !ok {
		attrib.attributes[attribName] = value
		return
	}

	attribType := reflect.TypeOf(attribInterface)
	if attribType.Kind() == reflect.Float64 {
		attrib.attributes[attribName] = value
	} else {
		panic(fmt.Sprintf("attribute %s is not a float64", attribName))
	}
}

// SetString sets a value of specified string attribute or creates a new one
func (attrib *CustomAttributes) SetString(attribName, value string) {
	attribInterface, ok := attrib.attributes[attribName]
	if !ok {
		attrib.attributes[attribName] = value
		return
	}

	attribType := reflect.TypeOf(attribInterface)
	if attribType.Kind() == reflect.String {
		attrib.attributes[attribName] = value
	} else {
		panic(fmt.Sprintf("attribute %s is not a string", attribName))
	}
}

// SetBool sets a value of specified boolean attribute or creates new one
func (attrib *CustomAttributes) SetBool(attribName string, value bool) {
	attribInterface, ok := attrib.attributes[attribName]
	if !ok {
		attrib.attributes[attribName] = value
		return
	}

	attribType := reflect.TypeOf(attribInterface)
	if attribType.Kind() == reflect.Bool {
		attrib.attributes[attribName] = value
	} else {
		panic(fmt.Sprintf("attribute %s is not a bool", attribName))
	}
}

// RemoveAttribute removes attribute with given name
func (attrib *CustomAttributes) RemoveAttribute(attribName string) {
	attrib.attributes[attribName] = nil
}
