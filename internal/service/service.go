package service

import (
	"errors"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var errEmptyStr = errors.New("empty string")

func isMorseStr(str string) bool {
	var hasLetter bool
	for _, c := range str {
		if c != '.' && c != '-' && c != ' ' {
			hasLetter = true
			break
		}
	}

	return !hasLetter
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
