package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"goreloaded"
)

func main() {
	if len(os.Args) == 3 {
		fileName := os.Args[1]
		file, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error")
			return
		}
		input := regexp.MustCompile("\n").ReplaceAllLiteralString(string(file), " \n ")
		splitarray := strings.Split(input, " ")
		splitarray = goreloaded.Formspaces(splitarray)
		splitarray = goreloaded.UpLowandCap(splitarray)
		splitarray = goreloaded.HexandBin(splitarray)
		splitarray = goreloaded.AtoAN(splitarray)
		output := ""
		for i := 0; i < len(splitarray); i++ {
			if splitarray[i] != "\n" {
				output += splitarray[i] + " "
			} else {
				output += splitarray[i]
			}
		}
		output = goreloaded.FixPunc(output)
		output = goreloaded.SingleQuote(output)
		err = os.WriteFile(os.Args[2], []byte(output), 0o644)
		if err != nil {
			fmt.Printf("Error")
			return
		}

	} else if len(os.Args) != 3 {
		fmt.Println("error: u did not use 3 args")
	}
}
