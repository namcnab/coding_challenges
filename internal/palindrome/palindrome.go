package palindrome

import (
	"regexp"
	"strings"
)

type Text struct {
	Name string `json:"name"`
}

func IsPalindromeStr(text string) bool {

	text = strings.ToLower(text) // Convert all letters to lowercase
	text = ReplacePunctuation(text, "")
	leftPointer := 0
	rightPointer := len(text) - 1

	if len(text) == 1 {
		return true
	}

	for leftPointer < rightPointer {
		if text[leftPointer] != text[rightPointer] {
			return false
		}
		leftPointer++
		rightPointer--
	}

	return true

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

func ReplacePunctuation(text string, replacement string) string {
	reg := regexp.MustCompile(`[[:punct:] | \s]`)
	return reg.ReplaceAllString(text, replacement)
}
