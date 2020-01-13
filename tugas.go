package tugas9

import (
	"fmt"
	"strings"
)

func Validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, fmt.Errorf("cannot be empty")
	}
	return true, nil
}
