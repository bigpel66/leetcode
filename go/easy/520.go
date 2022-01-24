// 520. Detect Captial
//
// https://leetcode.com/problems/detect-capital/

// detectCapitalUse function checks whether the string follows the captial rules or not.
// To return true, one of the rule should be satisfied. Or, return false.
// 1. All capital letters.
// 2. Only the first capital letter.
// 3. All non-capital letters.
// Divide into 2 cases and iterate own way.
// If the length of the word is 1, then it is always true.
func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	if unicode.IsUpper(rune(word[0])) {
		keepUpper := unicode.IsUpper(rune(word[1]))
		for i := 2; i < len(word); i++ {
			if keepUpper != unicode.IsUpper(rune(word[i])) {
				return false
			}
		}
	} else {
		for _, ch := range word {
			if unicode.IsUpper(rune(ch)) {
				return false
			}
		}
	}
	return true
}

