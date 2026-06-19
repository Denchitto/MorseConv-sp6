package service

import (
	"MorseConv-sp6/pkg/morse"
	"strings"
)

func MorseConv(data string) string {

	runesRemover := func(r rune) rune {
		switch r {
		case ' ', '-', '.':
			return -1
		}
		return r
	}
	blankCheck := strings.Map(runesRemover, data)
	if blankCheck == "" {
		return morse.ToText(data)
	}
	return morse.ToMorse(data)
}
