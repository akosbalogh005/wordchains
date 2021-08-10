#!/usr/bin/python

import time
import re
import unittest

def _get_next_word_list(dictionary : list, word : str) -> list:
    """
    Getting next potential items match to the word. Using regexp. 
    The rule consists of len(word) parts with OR relation. In each part 
    one different letter is wildcard (inregexp: \\w) the others should match.
    
    Example: word = 'abc' then regexp like: '^[\\w][b][c]$|[a][\\w][c]$|[a][b][\\w]$' 
    It is something like: '?bc' or 'a?c' or 'ab?'
    """
    result = []
    # create regexp rule
    rule = ""
    word_len = len(word)
    for i in range(0, word_len):
        r = "^"
        for j in range(0, word_len):
            if i == j:
                r += "\\w{1}"
            else:
                r += f"[{word[j]}]"
        r += "$"
        if i == 0:
            rule = r
        else:
            rule = rule + "|" + r
    reg = re.compile(rule)
    result = [w for w in dictionary if reg.search(w) and word != w ]
    #print("_get_next_word_list", rule, word, result)  
    return result

def _wordchain_inner(dictionary : list, start : str, end : str, current_list: list, best_list: list = []):
    
    #print("_wordchain_inner", dictionary, start, end, current_list, best_list)
    # if we have already proper chain than skip current if its the size greater -> cannot be the shortest, the best   
    if len(best_list) > 0:
        if len(current_list) > len(best_list):
            return

    has = False
    # prcessing all the next possibilities
    for next_word in _get_next_word_list(dictionary, start):
        if next_word in set(current_list):
            continue
        #print("_wordchain_inner next word ", start, "->", next_word)
        has = True
        # take care of the reference! cannot reassign value for current_list and for best_list (passing variable by reference)
        current_list.append(next_word)
        # reach the end, hurray
        if next_word == end:
            if best_list:
                if len(current_list) > len(best_list):
                    # better (shorter) list has already found
                    current_list.remove(next_word)
                    return
            best_list.clear()
            for v in current_list:
                best_list.append(v)
            #print("_wordchain_inner next new best list ", best_list)
            current_list.remove(next_word)
            return
        _wordchain_inner(dictionary, next_word, end, current_list, best_list)
        current_list.remove(next_word)

def wordchain(dictionary : list, start : str, end : str) -> list:
    """
    wordchain creator
    
    Try to make a wordchain starting from 'start' and ending with 'end'. 
    Successive entries in the chain must all be real words, and each 
    can differ from the previous word by just one letter
    http://codekata.com/kata/kata19-word-chains/
    
    Parameters
    ----------
    dictionary
        list of the valid words 
    start
        the starting word of the chain
    start
        ending word of the chain
    return
        the shortest wordchain list, empty if wordchain cannot be created  
        
    """
    assert len(start) == len(end), "The lenght of the starting and ending word should be the same"
    for d in dictionary:
        assert type(d) == str
    
    best_list = []
    _wordchain_inner(dictionary, start, end, [start], best_list)
    
    return best_list

if __name__ == '__main__':
    print (wordchain(["lead", "iead", "fead", "load", "goad", "gold"], "lead", "gold"))
    print(wordchain(["1", "2", "3", "4", "5"], "1", "3"))



class TestWordChain(unittest.TestCase):
    def test_get_next_word_list(self):
        self.assertEqual(_get_next_word_list(["aaa", "aba", "abb", "cbb", "bbbxxx"], "aaa" ), ["aba"])
        self.assertEqual(_get_next_word_list(["111", "112", "121", "222", "333"], "111" ), ["112", "121"])

    def test_wordchain(self):
        self.assertEqual(wordchain(["aaa", "aba", "abb", "bbb"], "aaa", "bbb"), ["aaa", "aba", "abb", "bbb"])
        self.assertEqual(wordchain([], "1", "2"), [])
        self.assertEqual(wordchain(["aaa", "aba", "aba", "bbb"], "aaa", "bbb"), [])
        self.assertEqual(wordchain(["lead", "load", "goad", "gold"], "lead", "gold"), ["lead", "load", "goad", "gold"])
        self.assertEqual(wordchain(["lead", "iead", "fead", "load", "goad", "gold"], "lead", "gold"), ["lead", "load", "goad", "gold"])
        self.assertEqual(wordchain(["1", "2", "3", "4", "5"], "1", "3"), ["1", "3"])
        self.assertEqual(wordchain(["111", "112", "555", "121", "123", "120", "150", "151", "155"], "111", "555"), ["111", "151", "155", "555"])
        self.assertEqual(wordchain(["1111", "1121", "1221", "1231", "1233", "3234", "3233", "1223", "1123", "1133", "1233","1233", "1133", "1113"], "1111", "3233"), ['1111', '1113', '1133', '1233', '3233'])
        
