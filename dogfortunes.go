//usr/bin/env go run $0 $@; exit

package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
)

func check(e error) {

	if (e != nil) {
		panic(e)
	}
}

func randomNum(min, max int) int {

    return rand.Intn(max - min) + min
}

func countLines(filename string) (int) {

	file, _ := os.Open(filename)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
	    lineCount++
	}
	return lineCount
}

func getFortune(filename string, randomLineNumber int) (string) {

	file, _ := os.Open(filename)
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	lineCount := 1
	for scanner.Scan() {
		if (lineCount == randomLineNumber) {
			return scanner.Text()
		}
	    lineCount++
	}
	return ""
}

func main() {

	lineCount := countLines("fortunes.txt")	

	rand.Seed(time.Now().UTC().UnixNano())
	
	randomLineNumber := randomNum(1, lineCount)
	fortune := getFortune("fortunes.txt", randomLineNumber)
	fmt.Printf("\n\t%s\n\n", fortune)	
}