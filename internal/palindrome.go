package palindrome

import "strings"

type Text struct {
	Name string `json:"name"`
}

func IsPalindromeStr(text string) bool {
	// Convert all letters to lowercase
	text = strings.ToLower(text)

	// Convert text to byte slice
	byteSlice := []byte(text)
	bkwdText := ""

	for i := len(byteSlice) - 1; i > -1; i-- {
		bkwdText += string(byteSlice[i])
	}

	if bkwdText == text {
		return true
	} else {
		return false
	}
}

func longestPalindromeStr(text string) {

}
