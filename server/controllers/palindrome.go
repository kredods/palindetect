package controllers

import(
	"regexp"
	"strings"
)

func IsPalindrome(value string) bool{
	var re = regexp.MustCompile(`[^A-Za-z0-9]`)
	cleanString := strings.ToLower(re.ReplaceAllString(value,``))

	for i := 0; i < len(cleanString)/2 +1 ; i++ {
		if(cleanString[i]!= cleanString[len(cleanString)-i-1]){
			return false
		}
	}
	return true
}