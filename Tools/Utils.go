package Tools

import "fmt"

func ColorToHex(r, g, b uint8) string {
	return fmt.Sprintf("%02x%02x%02x", r, g, b)
}
