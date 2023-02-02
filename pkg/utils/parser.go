package utils

import (
	"golang.org/x/text/encoding/charmap"
)

func ParseText(bytes []byte) (string, error) {
	if dst, err := charmap.Windows1252.NewDecoder().Bytes(bytes); err != nil {
		return "", err
	} else {
		str, err := charmap.Windows1252.NewEncoder().String(string(dst))
		if err != nil {
			return "", err
		}
		return str, nil
	}
}
