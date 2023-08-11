package Fun

import (
	"ScrapBlueprint/Structs"
	"fmt"
	"github.com/nfnt/resize"
	"image/png"
	"os"
)

const (
	pixMax uint64 = 65535
)

func rgbToHex(r, g, b uint64) string {
	return fmt.Sprintf("%02x%02x%02x", r, g, b)
}

func GeneratePixelArt(imageFile, shapeIdToUse string, resizeX, resizeY uint) Structs.Body {
	fl, err := os.Open(imageFile)
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(fl)
	if err != nil {
		panic(err)
	}
	img = resize.Resize(resizeX, resizeY, img, resize.Lanczos3)
	var body = Structs.MakeBody()
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			block := Structs.MakeBlock(shapeIdToUse, rgbToHex(uint64(r)*255/pixMax, uint64(g)*255/pixMax, uint64(b)*255/pixMax), Structs.Position{X: x, Y: y})
			body.AddBlock(block)
		}
	}
	return body
}
