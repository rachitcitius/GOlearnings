// UpperCase1stCharInWord
//Function converts 1st character of every word to upper case
//Command Line utility
package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	//fmt.Println(ToUpperChar("go language is so cool & wonderful. I am loving it!"))
	argsWithoutProg := os.Args[1]
	fmt.Println(ToUpperChar(argsWithoutProg))
}

func ToUpperChar(fullString string) string  {
	splitWords := strings.Split(fullString, " ")
	newWords := make([]string, len(splitWords))
	
	for i, eachWord := range splitWords {
		newWords[i] = strings.ToUpper(string(eachWord[0])) + string(eachWord[1: len(eachWord)])
		//i++
	}
	return(strings.Join(newWords, " "))
}
