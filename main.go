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
	r := []rune(line)
	r = append(r, ' ')

	for i := 0; i <= len(r)-1; i++ {

		// - 12345
		if unicode.IsDigit(r[i]) {

			if i != 0 {
				// - qwe\\5
				// если i == digit и перед ней есть \\, то i записываем как digit
				if r[i-1] == 92 && r[i-2] == 92 {
					arr += string(r[i])

					// - qwe\45
					//	иначе если i == digit и перед ним \ и после него digit, то записываем i как string
				} else if r[i-1] == 92 && unicode.IsDigit(r[i+1]) {
					arr += "_" + string(r[i]) + "_"

					// - qwe\4\5
					// иначе если i == digit и перед ним \ и после него не digit или конец строки, то записываем i как string
				} else if r[i-1] == 92 && !unicode.IsDigit(r[i+1]) {
					arr += "_" + string(r[i]) + "_"

					// в остальных случаях 	i записываем как digit
				} else {
					arr += string(r[i])
				}

				// если i = 0 то i записываем как digit
			} else {
				arr += string(r[i])
			}
		}

		// 	- qwerty
		if unicode.IsLetter(r[i]) {
			arr += "_" + string(r[i]) + "_"
		}

		// 	- \\
		if r[i] == 92 && r[i-1] == 92 {
			arr += "_" + string(r[i]) + "_"
		}
	}

	l := strings.Split(arr, "_")

	for i := 2; i <= len(l); i = i + 2 {
		number, _ := strconv.Atoi(l[i])
		if number == 0 {
			number = 1
		}
		val := strings.Repeat(l[i-1], number)
		value += val
	}

	if value == "" {
		fmt.Println("Incorrect string")
	} else {
		fmt.Println(value)
	}
}

func main() {

	// "a4bc2d5e" => "aaaabccddddde"
	// "abcd" => "abcd"
	// "45"` => "" (некорректная строка)
	// "qwe\4\5" => "qwe45" (*)
	// "qwe\45" => "qwe44444" (*)
	// "qwe\\5" => "qwe\\\\\" (*)

	var text string

	for text != "-" { // break the loop if text == "-"
		fmt.Print("Enter your string: ")
		fmt.Fscan(os.Stdin, &text)
		unPackStr(text)
	}
}
