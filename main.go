package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func unPackStr(line string) {
	var arr, value string
	for _, item := range line {

		if unicode.IsDigit(item) {
			arr += strings.TrimSpace(string(item))
		} else {
			arr += " " + string(item) + " "
		}
	}

	l := strings.Split(arr, " ")

	for i := 2; i <= len(l); i = i + 2 {
		number, _ := strconv.Atoi(l[i])
		if number == 0 {
			number = 1
		}
		val := strings.Repeat(l[i-1], number)
		value += val
	}

	if value == "" {
		fmt.Println("Incorrect String")
	} else {
		fmt.Println(value)
	}
}

func main() {
	var text string
	for text != "-" { // break the loop if text == "-"
		fmt.Print("Enter your string: ")
		fmt.Fscan(os.Stdin, &text)
		unPackStr(text)
	}
}
