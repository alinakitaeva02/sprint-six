package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("incorrect data format")
	}
	isMorse := true
	for _, r := range input {
		if r != '.' && r != '-' && r != ' ' {
			isMorse = false
			break
		}
	}
	if isMorse {
		text := morse.ToText(input)
		return text, nil
	}

	hasLetter := strings.IndexFunc(input, unicode.IsLetter) != -1
	if !hasLetter {
		return "", errors.New("incorrect data format")
	}
	morseCode := morse.ToMorse(input)
	return morseCode, nil
}
