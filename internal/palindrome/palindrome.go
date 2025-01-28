package palindrome

import (
	"strings"
)

type Text struct {
	Name string `json:"name"`
}

func IsPalindromeStr(text string) bool {
	// Convert all letters to lowercase
	text = strings.ToLower(text)

	if len(text) == 1 {
		return true
	}

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

func LongestPalindromeStr(text string) string {
	startInd := 0
	endInd := len(text)
	tempText := text[startInd:endInd]
	isPal := IsPalindromeStr(tempText)

	for !isPal && (endInd-startInd) >= 3 {
		startInd++
		endInd--
		tempText = text[startInd:endInd]
		isPal = IsPalindromeStr(tempText)
	}

	if isPal && len(tempText) >= 3 {
		return tempText
	} else {
		return ""
	}
}
