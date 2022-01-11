package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

// override swap function in sort package
// for Pairlist interface
func (p PairList) Len() int {
	return len(p)
}

// override swap function in sort package
// for Pairlist interface
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// reconfigure less function to 'more' function
// so the sorting is descending instead ascending
func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func main() {
	var lines []string

	// initiate scanner
	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("You can enter the words or sentences and after you finish inserting words or sentence, please press [enter] key then type ':exit' to process.")
	fmt.Println("Enter Lines:")

	// do loop text scan
	for scn.Scan() {
		// scan text in line
		line := scn.Text()

		// check string to break the scan loop
		if line == ":exit" {
			break
		}

		// trim and separate text to slice of words
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		words := strings.FieldsFunc(line, f)

		lines = append(lines, words...)
	}

	// map the word - counter pair
	wordsCount := make(map[string]int)
	for _, line := range lines {
		loweredWord := strings.ToLower(line)
		if _, ok := wordsCount[loweredWord]; !ok {
			wordsCount[loweredWord] = 1
		} else {
			wordsCount[loweredWord] = wordsCount[loweredWord] + 1
		}
	}

	// sorting function
	pairList := make(PairList, len(wordsCount))

	// insert mapped word - counter
	// to [int]Pair{key,value} slice
	i := 0
	for k, v := range wordsCount {
		pairList[i] = Pair{k, v}
		i++
	}

	// use sort package sorting
	// by counter descending
	sort.Sort(pairList)

	// loop the sorted slice
	// get the top ten from list
	for i, k := range pairList {
		if i >= 10 {
			break
		}

		fmt.Printf("%v\t%v\n", k.Key, k.Value)
	}
}
