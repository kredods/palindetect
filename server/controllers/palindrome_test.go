package controllers

import (
	"testing"
)

func TestPalindromes(t *testing.T){
	successCases :=[] string{
		"kayak",
		"mom",
		"dad",
		"racecar",
		"A Man, A Plan, A Canal-Panama",
		"Madam In Eden, I'm Adam",
		"Mr. Owl Ate My Metal Worm",
	}

	failureCases := []string{
		"jack",
		"jill",
		"went",
		"up",
		"hill",
	}

	for _, element := range successCases{
		assertPalindrome(element,t)

	}

	for _, element := range failureCases{
		assertNotPalindrome(element, t)
	}

}

func assertNotPalindrome(value string, t *testing.T){
	if(isPalindrome(value)){
		t.Fatal(value + " is a palindrome")
	}
}

func assertPalindrome(value string, t *testing.T){
	if(!isPalindrome(value)){
		t.Fatal( value + " is not a palindrome")
	}
}