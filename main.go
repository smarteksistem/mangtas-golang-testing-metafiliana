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

func (p PairList) Len() int {
	return len(p)
}
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

	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("You can enter the words or sentences and after you finish inserting words or sentence, please press [enter] key then type ':exit' to process.")
	fmt.Println("Enter Lines:")
	for scn.Scan() {
		line := scn.Text()
		if line == ":exit" {
			break
		}

		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		words := strings.FieldsFunc(line, f)

		lines = append(lines, words...)
	}

	wordsCount := make(map[string]int)
	for _, line := range lines {
		loweredWord := strings.ToLower(line)
		if _, ok := wordsCount[loweredWord]; !ok {
			wordsCount[loweredWord] = 1
		} else {
			wordsCount[loweredWord] = wordsCount[loweredWord] + 1
		}
	}

	p := make(PairList, len(wordsCount))

	i := 0
	for k, v := range wordsCount {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	for i, k := range p {
		if i >= 10 {
			break
		}

		fmt.Printf("%v\t%v\n", k.Key, k.Value)
	}
}
