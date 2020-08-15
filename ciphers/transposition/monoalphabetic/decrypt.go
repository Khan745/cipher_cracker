package monoalphabetic

import (
	"strings"
)

func transformDecrypt(char rune, alphabtNormal string, alphabetSecret string) string {
	index := strings.Index(alphabetSecret, string(char))
	decryptedCharacter := alphabtNormal[index]
	return string(decryptedCharacter)
}

func Decrypt(text string, alphabtNormal string, alphabetSecret string) string {
	newText := ""
	for _, char := range text {
		newText += transformDecrypt(char, alphabtNormal, alphabetSecret)
	}
	return newText
}
