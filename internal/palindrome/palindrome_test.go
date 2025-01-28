package palindrome

import (
	"fmt"
	"log"
	"testing"
)

// TestLongestPalindromeStr calls greetings.Hello with a name, checking
// for a valid return value.
func TestIsPalindromeStr(t *testing.T) {
	textSlice := []string{"racecar", "jtabatk", "bob", "nneka", "cup", "jueroreuk", "hello", "madam", "world", "level", "example", "rotor", "deified"}
	expectedResults := []bool{true, false, true, false, false, false, false, true, false, true, false, true, true}

	fmt.Println("****** TestIsPalindromeStr  ******")
	for index, value := range textSlice {
		isPal := IsPalindromeStr(value)
		expectedResult := expectedResults[index]

		log.Printf("Text: %s, isPal: %t", value, isPal)

		if isPal != expectedResult {
			t.Fatalf(`IsPalindromeStr("%s") = %t,  %t is expected.`, value, isPal, expectedResult)
		}
	}

	fmt.Println("************************")
}

func TestLongestPalindromeStr(t *testing.T) {
	textSlice := []string{"racecar", "jtabatk", "bob", "nneka", "cup", "jueroreuk"}
	expectedResults := []string{"racecar", "tabat", "bob", "", "", "ueroreu"}

	fmt.Println("****** TestLongestPalindromeStr  ******")
	for index, value := range textSlice {
		longestPal := LongestPalindromeStr(value)
		expectedResult := expectedResults[index]

		log.Printf("Text: %s, longestPal: %s", value, longestPal)

		if longestPal != expectedResult {
			t.Fatalf(`LongestPalindromeStr("%s") = %s,  %s is expected.`, value, longestPal, expectedResult)
		}
	}
	fmt.Println("************************")
}
