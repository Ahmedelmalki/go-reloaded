package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 3 {
		fileName := os.Args[1]
		file, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error")
			return
		}
		
		splitarray := strings.Split(string(file), "\n")
		splitarray = strings.Split(string(file), " ")
		splitarray = fixPunc(splitarray)
		output := strings.Join(splitarray, " ")
		// err = os.WriteFile(os.Args[2], []byte(output), 0644)
		// if err != nil {
		// 	fmt.Printf("Error")
		// 	return
		fmt.Printf(output)
	} else if len(os.Args) != 3 {
		fmt.Printf("error: u did not use 3 args")
	}
}


