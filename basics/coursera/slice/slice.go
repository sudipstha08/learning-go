// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
// The program should be written as a loop. Before entering the loop, the program should create an empty
// integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an
// integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints
// the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers
// which the user decides to enter. The program should only quit (exiting the loop) when the user enters the
// character ‘X’ instead of an integer

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type UserPrompt int

const (
	ListSize            int    = 3
	MaxUserPromptLength int    = 40
	QuitCode            string = "X"
)

const (
	ReadStringValue UserPrompt = iota
	ShowStringValue
	RegexFound
	RegexNotFound
)

func GetUserPromptPadded(up UserPrompt, str string, strFmt string) string {
	var paddedStr = str + strings.Repeat(" ", (MaxUserPromptLength-len(str))) + strFmt
	if !(up == ReadStringValue) {
		paddedStr = paddedStr + "\n"
	}
	return paddedStr
}

func GetUserPrompt(up UserPrompt) string {
	var strReturnVal string
	switch up {
	case ReadStringValue:
		strReturnVal = GetUserPromptPadded(up, "Please enter an integer ('X' to Exit):", "")
	case RegexFound:
		strReturnVal = GetUserPromptPadded(up, "Valid Input!", "")
	case RegexNotFound:
		strReturnVal = GetUserPromptPadded(up, "Invalid Input!", "")
	}
	return strReturnVal
}

func main() {
	var intSlice = make([]int, ListSize)

	re := regexp.MustCompile(`^\d+$`)
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Printf(GetUserPrompt(ReadStringValue))
	for scanner.Scan() {
		var strVal string = scanner.Text()
		if strVal == QuitCode {
			os.Exit(0)
		}
		match := re.Match([]byte(strVal))
		intVal, err := strconv.Atoi(strVal)
		if match == false || err != nil {
			println(GetUserPrompt(RegexNotFound))
			fmt.Printf(GetUserPrompt(ReadStringValue))
			continue
		}
		intSlice = append(intSlice, intVal)
		sort.Sort(sort.IntSlice(intSlice))
		var result []int = intSlice[ListSize:len(intSlice)]
		fmt.Println(result)
		fmt.Printf(GetUserPrompt(ReadStringValue))
	}
}