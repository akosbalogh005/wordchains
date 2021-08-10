package main

import (
	"fmt"
	"regexp"
	"sort"
)

// getNextWordList
//Getting next potential items match to the word. Using regexp.
//The rule consists of len(word) parts with OR relation. In each part
//one different letter is wildcard (inregexp: \\w) the others should match.
//
//Example: word = 'abc' then regexp like: '^[\\w][b][c]$|[a][\\w][c]$|[a][b][\\w]$'
//It is something like: '?bc' or 'a?c' or 'ab?'
func getNextWordList(dictionary *[]string, word string) []string {
	var result []string
	// create regexp rule
	rule := ""
	wordLen := len(word)
	for i := 0; i < wordLen; i++ {
		r := "^"
		for j := 0; j < wordLen; j++ {
			if i == j {
				r += "\\w{1}"
			} else {
				r += fmt.Sprintf("[%c]", word[j])
			}
		}
		r += "$"
		if i == 0 {
			rule = r
		} else {
			rule = rule + "|" + r
		}
	}
	reg := regexp.MustCompile(rule)
	for _, w := range *dictionary {
		if reg.MatchString(w) && word != w {
			result = append(result, w)
		}
	}
	//fmt.Printf("getNextWordList - rule: %v, word:%v, return: %v \n", rule, word, result)
	return result
}

func wordChainInner(dictionary *[]string, start string, end string, currentList *[]string, bestList *[]string) {
	//fmt.Println("wordChainInner", *dictionary, start, end, *currentList, *bestList)
	// if we have already proper chain than skip current if its the size greater -> cannot be the shortest, the best
	if len(*bestList) > 0 && len(*currentList) > len(*bestList) {
		return
	}

	// prcessing all the next possibilities
	for _, nextWord := range getNextWordList(dictionary, start) {
		found := false
		for _, w := range *currentList {
			if w == nextWord {
				found = true
				break
			}
		}
		if found {
			continue
		}
		//fmt.Println("wordChainInner next word : ", start, "->", nextWord)
		*currentList = append(*currentList, nextWord)
		// reach the end, hurray
		if nextWord == end {
			if len(*currentList) > len(*bestList) && len(*bestList) > 0 {
				// better (shorter) list has already found
				*currentList = (*currentList)[:len(*currentList)-1]
				return
			}
			*bestList = (*bestList)[:0]
			*bestList = append(*bestList, *currentList...)
			//fmt.Println("wordChainInner next new best list ", *bestList)
			*currentList = (*currentList)[:len(*currentList)-1]
			return
		}
		wordChainInner(dictionary, nextWord, end, currentList, bestList)
		*currentList = (*currentList)[:len(*currentList)-1]
	}

}

// WordChain wordchain creator
//
//Try to make a wordchain starting from 'start' and ending with 'end'.
//Successive entries in the chain must all be real words, and each
//can differ from the previous word by just one letter
//http://codekata.com/kata/kata19-word-chains/
//
//Parameters
//
//dictionary
//    list of the valid words
//start
//    the starting word of the chain
//start
//    ending word of the chain
//return
//    the shortest wordchain list, empty if wordchain cannot be created
func WordChain(dictionary []string, start string, end string) ([]string, error) {

	if len(start) != len(end) {
		return nil, fmt.Errorf("the lenght of the starting and ending word should be the same")
	}

	bestList := []string{}
	var currentList []string
	dictionarySorted := make([]string, len(dictionary))

	// sorting to be able to search easily
	copy(dictionarySorted, dictionary)
	sort.Strings(sort.StringSlice(dictionarySorted))

	currentList = append(currentList, start)
	wordChainInner(&dictionarySorted, start, end, &currentList, &bestList)
	return bestList, nil
}

func main() {

	ret, err := WordChain([]string{"lead", "iead", "fead", "load", "goad", "gold"}, "lead", "gold")
	fmt.Printf("err:%v. ret: %v", err, ret)

}
