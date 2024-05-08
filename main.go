package main

import (
	"log"
	"strings"
)

var parsingTable = map[string]map[string]string{
	"S": {
		"a": "aQ",
	},
	"Q": {
		"a": "a",
		"b": "bAB",
		"c": "cAB",
	},
	"A": {
		"a": "aF",
	},
	"F": {
		"a": "J",
		"b": "J",
		"c": "J",
		"$": "eps",
	},
	"J": {
		"c": "CaF",
		"b": "bF",
	},
	"C": {
		"a": "abG",
		"c": "cE",
	},
	"G": {
		"a": "AB",
		"c": "cA",
	},
	"E": {
		"a": "C",
	},
	"B": {
		"a": "CAa",
	},
}

func parse(input string) {
	stack := &Stack{}
	flag := false
	stack.Push("$")
	usage := "S"
	for _, ctr := range input {
		stack.Push(string(ctr))
		for ind, symb := range usage {
			top := stack.Peek()
			if symb >= 97 && symb <= 122 {
				if ind+1 < len(usage) && rune(usage[ind+1]) <= 97 {
					usage = usage[1:]
					flag = true
					continue
				} else if ind+1 < len(usage) && rune(usage[ind+1]) >= 97 {
					usage = usage[1:]
					flag = true
					if top == string(symb) {
						continue
					} else {
						log.Panic("Incorrect rule")
					}
				}
			}

			rule, ok := parsingTable[string(symb)][top]
			if !ok {
				log.Panic("Incorrect rule")
			} else {
				if flag {
					usage = strings.Replace(usage, string(usage[0]), rule, 1)
					flag = false
				} else {
					usage = strings.Replace(usage, string(usage[ind]), rule, 1)
				}
				stack.Pop()
			}
		}
	}
	log.Println("Correct")
}

func main() {
	input := "abca"
	//input1 := "aa"
	parse(input)
	//parse(input1)
}
