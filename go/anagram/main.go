package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "documenting"

	fmt.Printf("All two worded anagrams of the word %q are:\n %v\n", word, TwoWordedAnagrams(word))
}

func TwoWordedAnagrams(word string) []string {
	return makeTwoWordedAnagramsOf(word)
}

func makeTwoWordedAnagramsOf(word string) (allAnagrams []string) {
	allLettersFromWord := splitWordIntoArrayOfLetters(word)
	allAnagrams = joinEachLetterWithEachOfTheOthers(allLettersFromWord)

	return allAnagrams
}

func splitWordIntoArrayOfLetters(word string) []string {
	allLetters := strings.Split(word, "")
	return allLetters
}

func joinEachLetterWithEachOfTheOthers(allLettersFromWord []string) (allAnagrams []string) {
	for thisPosition, thisLetter := range allLettersFromWord {
		allOthers := makeListWithAlButOneIn(thisPosition, allLettersFromWord)
		allAnagrams = allPairsOf(thisLetter, allOthers)
	}
	return
}

func makeListWithAlButOneIn(positionOfThisLetter int, listOfAllLetters []string) (allOthers []string) {
	allOthers = append(allOthers, listOfAllLetters[:positionOfThisLetter]...)
	allOthers = append(allOthers, listOfAllLetters[positionOfThisLetter+1:]...)
	return
}

func allPairsOf(thisLetter string, allOthers []string) (allAnagrams []string) {
	for _, otherLetter := range allOthers {
		anagram := fmt.Sprintf("%s%s", thisLetter, otherLetter)
		allAnagrams = append(allAnagrams, anagram)
	}
	return
}
