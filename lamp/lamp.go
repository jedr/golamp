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
	message := fmt.Sprintf("%v light is on", color)
	return message, nil
}

func LightDown(color string) (string, error) {
	if color == "" {
		return "", errors.New("empty color")
	}
	message := fmt.Sprintf("%v light is off", color)
	return message, nil
}
