package main

import (
	"fmt"
	"os"
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
		// here where we will call our funcs :)
		splitarray := strings.Split(string(file), "\n")
		splitarray = strings.Split(string(file), " ")
		fmt.Println(splitarray)
		splitarray = goreloaded.Formspaces(splitarray)
		splitarray = goreloaded.UpLowandCap(splitarray)
		splitarray = goreloaded.HexandBin(splitarray)
		splitarray = goreloaded.AtoAN(splitarray)
		output := strings.Join(splitarray, " ")
		output = goreloaded.FixPunc(output)
		output = goreloaded.SingleQuote(output)
		err = os.WriteFile(os.Args[2], []byte(output), 0o644)
		fmt.Println(output)
		if err != nil {
			fmt.Printf("Error")
			return
		}

	} else if len(os.Args) != 3 {
		fmt.Println("error: u did not use 3 args")
	}
}
