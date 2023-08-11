package Serialization

import (
	"ScrapBlueprint/Structs"
	"encoding/json"
	"os"
)

func SaveBlueprint(baseDirectory string, blueprint *Structs.Blueprint, blueprintMeta Structs.BlueprintDescription, putThumbnailPlaceholder bool) error {
	if _, err := os.Stat(baseDirectory); err != nil {
		err := os.MkdirAll(baseDirectory, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if putThumbnailPlaceholder {
		if _, err := os.Stat(baseDirectory + "/icon.png"); err != nil {
			_, err := os.Create(baseDirectory + "/icon.png")
			if err != nil {
				return err
			}
		}
		err := CopyFile("Assets/no_thumb.png", baseDirectory+"/icon.png")
		if err != nil {
			return err
		}
	}
	bpFile, err := os.OpenFile(baseDirectory+"/blueprint.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer bpFile.Close()
	if err != nil {
		return err
	}
	metaFile, err := os.OpenFile(baseDirectory+"/description.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer metaFile.Close()
	if err != nil {
		return err
	}
	marshaledBp, err := json.Marshal(blueprint)
	if err != nil {
		return err
	}
	_, err = bpFile.Write(marshaledBp)
	if err != nil {
		return err
	}
	marshaledMeta, err := json.Marshal(blueprintMeta)
	if err != nil {
		return err
	}
	_, err = metaFile.Write(marshaledMeta)
	if err != nil {
		return err
	}
	return nil
}

func SaveBlueprintWithThumbnail(baseDirectory, thumbPath string, blueprint *Structs.Blueprint, blueprintMeta Structs.BlueprintDescription) error {
	if _, err := os.Stat(baseDirectory); err != nil {
		err := os.MkdirAll(baseDirectory, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(baseDirectory + "/icon.png"); err != nil {
		_, err := os.Create(baseDirectory + "/icon.png")
		if err != nil {
			return err
		}
	}
	err := CopyFile(thumbPath, baseDirectory+"/icon.png")
	if err != nil {
		return err
	}
	return SaveBlueprint(baseDirectory, blueprint, blueprintMeta, false)
}
