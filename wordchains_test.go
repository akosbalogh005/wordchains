package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextWordList(t *testing.T) {
	assert.Equal(t, getNextWordList(&[]string{"aaa", "aba", "abb", "cbb", "bbbxxx"}, "aaa"), []string{"aba"})
	assert.Equal(t, getNextWordList(&[]string{"111", "112", "121", "222", "333"}, "111"), []string{"112", "121"})
}

func wordChainTestHelper(t *testing.T, dictionary []string, start string, end string, result []string) {
	ret, err := WordChain(dictionary, start, end)
	assert.Equal(t, err, nil)
	assert.Equal(t, result, ret)
}

func TestWordChain(t *testing.T) {
	wordChainTestHelper(t, []string{"aaa", "aba", "abb", "bbb"}, "aaa", "bbb", []string{"aaa", "aba", "abb", "bbb"})
	wordChainTestHelper(t, []string{}, "1", "2", []string{})
	wordChainTestHelper(t, []string{"aaa", "aba", "aba", "bbb"}, "aaa", "bbb", []string{})
	wordChainTestHelper(t, []string{"lead", "load", "goad", "gold"}, "lead", "gold", []string{"lead", "load", "goad", "gold"})
	wordChainTestHelper(t, []string{"lead", "iead", "fead", "load", "goad", "gold"}, "lead", "gold", []string{"lead", "load", "goad", "gold"})
	wordChainTestHelper(t, []string{"1", "2", "3", "4", "5"}, "1", "3", []string{"1", "3"})
	wordChainTestHelper(t, []string{"111", "112", "555", "121", "123", "120", "150", "151", "155"}, "111", "555", []string{"111", "151", "155", "555"})
	wordChainTestHelper(t, []string{"1111", "1121", "1221", "1231", "1233", "3234", "3233", "1223", "1123", "1133", "1233", "1233", "1133", "1113"}, "1111", "3233", []string{"1111", "1113", "1133", "1233", "3233"})
}
