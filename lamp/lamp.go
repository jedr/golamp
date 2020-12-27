package lamp

import (
	"errors"
	"fmt"
)

// LightUp lights up the lamp in color of your choice.
func LightUp(color string) (string, error) {
	if color == "" {
		return "", errors.New("empty color")
	}
	message := fmt.Sprintf("%v light is up!", color)
	return message, nil
}
