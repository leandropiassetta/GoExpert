package portuguese

import (
	"fmt"
	"strings"
)

func TranslateLoveForPortugueseWord(word string) (string, error) {
	if strings.ToUpper(word) != "LOVE" {
		return "", fmt.Errorf("D'ont translate this word: %s", word)
	}

	return "Amor", nil
}
