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
		// here where we will call our funcs :)
		splitarray := strings.Split(string(file), "\n")
		splitarray = strings.Split(string(file), " ")
		output := strings.Join(splitarray, " ")
		err = os.WriteFile(os.Args[2], []byte(output), 0o644)
		if err != nil {
			fmt.Printf("Error")
			return
			// fmt.Printf(output)
		}

	} else if len(os.Args) != 2 {
		fmt.Printf("error: u did not use 3 args")
	}
}
