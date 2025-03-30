package main

func ReversePrefix(word string, ch byte) string {
	var subStr string
	var result string
	var idx int
	var found bool

	// Get sub string until ch first occurance
	for i, c := range word {
		subStr += string(c)

		if byte(c) == ch {
			idx = i
			found = true
			break
		}
	}

	if !found {
		return word
	}

	// Reverse the sub string
	for i := len(subStr) - 1; i >= 0; i-- {
		result += string(subStr[i])
	}

	return result + word[idx+1:]
}
