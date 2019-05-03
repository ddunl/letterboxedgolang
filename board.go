package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
	"bufio"
	"io"
	"time"
)


func main() {
	startTime := time.Now()

	args := os.Args[1:]
	if (len(args) != 1) {
		os.Exit(1)
	}
	
	puzz := Unpack(args[0])

	fname := "modWords2.txt"

	playableWords := getDict(fname, puzz)
	fmt.Printf("Found all valid words... (%d) \n", len(playableWords))
	sols := getSolutions(playableWords)
	fmt.Printf("Found all solutions... (%d) \n", len(sols))

	for i, sol := range sols {
		fmt.Printf("%d. %s \n", i+1, sol)
	}


	totalTime := time.Now().Sub(startTime)

	fmt.Printf("%v \n", totalTime)

}

func noRepeats(str string) bool {
	//sees if a string has repeating characters
	last := ' '
	for _, s := range str {
		if (s == last) {
			return false
		}
		last = s
	}
	return true
}



func Unpack (bstr string) []string {
	//unpacks board from string to string slice
	if len(bstr) != 12 {
		fmt.Printf("Wrong board length")
		return []string{"aaa", "aaa", "aaa", "aaa"}
	}

	return []string{bstr[:3], bstr[3:6], bstr[6:9], bstr[9:12]}

}

func isPlayable(str string, puzz []string) bool {
	//checks if a word is playable on the grid
	hashes := ""

	for _, let := range str {
		for j, side := range puzz {
			if (strings.Contains(side, fmt.Sprintf("%c", let))) {
				hashes += strconv.Itoa(j)
				break
			}
		}
	}

	if (len(hashes) != len(str)) {
		return false
	} else {
		return noRepeats(hashes)
	}

	 
}

func getDict(fname string, puzz []string) []string {
	//gets valid "dictionary" of words
	file, e := os.Open(fname)
	var ret []string
	if e != nil {
		panic(e)
	}

	defer file.Close()

	r := bufio.NewReader(file)
		
	for {

		line, _, e := r.ReadLine()

		if e == io.EOF {
			break
		} else if isPlayable(string(line), puzz) {
			ret = append(ret, string(line))
		}

		
	}
	
	return ret
}

func isSolution (words []string) bool {
	//checks if array of words is a solution
	letters := strings.Join(words, "")
	set := ""
	for _, let := range letters {
		strlet := string(let)
		if ! strings.Contains (set, strlet) {
			set = set + strlet
		}
	}

	if len(set) == 12 {
		return true
	} else {
		return false
	}

}


func getSolutions (words []string) []string {
	//finds solutions based on valid words
	ret := []string{}
	for _, word1 := range words {
		for _, word2 := range words {
			if isSolution([]string{word1, word2}) && word1[len(word1) - 1] == word2[0]{
				
				ret = append(ret, word1 + " " + word2)
			}
		}
	}

	return ret
}
