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
		input := strings.Split(string(file), "\n")
		var line []string
		output := ""
		for i := 0; i < len(input); i++ {
			line = strings.Split(input[i], " ")
			line = goreloaded.UpLowandCap(line)
			line = goreloaded.HexandBin(line)
			line = goreloaded.Formspaces(line)
			line = goreloaded.AtoAN(line)
			temp := strings.Join(line, " ")
			temp = goreloaded.FixPunc(temp)
			temp = goreloaded.SingleQuote(temp)
			temp = strings.Trim(temp, " ")
			output += temp + "\n"
		}
		err = os.WriteFile(os.Args[2], []byte(output), 0o644)
		if err != nil {
			fmt.Printf("Error")
			return
		}
	} else if len(os.Args) != 3 {
		fmt.Println("error: u did not use 3 args")
	}
}
