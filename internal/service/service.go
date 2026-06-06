package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var errEmptyStr = errors.New("empty string")

func isMorseStr(str string) bool {
	return strings.ToUpper(str) == strings.ToLower(str)
}

func HandleStr(str string) (string, error) {
	if str == "" {
		return "", errEmptyStr
	}

	if isMorseStr(str) {
		return morse.ToText(str), nil
	}

	return morse.ToMorse(str), nil
}
