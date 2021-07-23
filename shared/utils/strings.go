package utils

import (
	"strconv"
	"strings"
)

// IsEmptyString func ti check empty string
func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// ConvertArrStringToInt to convert arr of string to arr of int
func ConvertArrStringToInt(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
