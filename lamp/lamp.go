package lamp

import (
	"fmt"
)

// LightUp lights up the lamp in color of your choice.
func LightUp(color string) string {
	message := fmt.Sprintf("%v light is up!", color)
	return message
}
