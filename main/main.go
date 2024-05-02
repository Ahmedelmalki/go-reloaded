package main

import (
	"fmt"
	"os"
	"strings"

	"goreloaded"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./[program_name] [input_file] [name_of_output]")
	}

	fileName := os.Args[1]
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error")
		return
	}
	input := strings.Split(string(file), "\n")
	output := processInput(input)
	err = os.WriteFile(os.Args[2], []byte(output), 0o644)
	if err != nil {
		fmt.Printf("Error")
		return
	}
}

func formatText(input string) string {
	var line []string
	line = strings.Split(input, " ")
	line = goreloaded.UpLowandCap(line)
	line = goreloaded.HexandBin(line)
	line = goreloaded.Formspaces(line)
	line = goreloaded.AtoAN(line)
	tmp := strings.Join(line, " ")
	tmp = goreloaded.FixPunc(tmp)
	tmp = goreloaded.SingleQuote(tmp)
	tmp = strings.Trim(tmp, " ")
	return (tmp + "\n")
}

func processInput(input []string) string {
	output := ""
	for i := 0; i < len(input); i++ {
		output += formatText(input[i])
	}
	return output
}
