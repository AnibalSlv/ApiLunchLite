package utils

import (
	"unicode"
)

func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	// Convertimos a runas para no romper caracteres especiales (como Ñ o tildes)
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
