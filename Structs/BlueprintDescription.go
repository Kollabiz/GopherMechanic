package Structs

import uuid "github.com/satori/go.uuid"

const (
	emptyDescription = "#{STEAM_WORKSHOP_NO_DESCRIPTION}"
)

type BlueprintDescription struct {
	Description string `json:"description"`
	LocalId     string `json:"localId"`
	Name        string `json:"name"`
	Type        string `json:"type"`    // This field should always be "Blueprint"
	Version     int    `json:"version"` // This field is always set to 0
}

func MakeBlueprintDescription(name, description string) BlueprintDescription {
	blueprintUuid := uuid.NewV4()

	return BlueprintDescription{
		Description: description,
		LocalId:     blueprintUuid.String(),
		Name:        name,
		Type:        "Blueprint",
		Version:     0,
	}
}

func MakeBlueprintWithoutDescription(name string) BlueprintDescription {
	blueprintUuid := uuid.NewV4()

	return BlueprintDescription{
		Description: emptyDescription,
		LocalId:     blueprintUuid.String(),
		Name:        name,
		Type:        "Blueprint",
		Version:     0,
	}
}
