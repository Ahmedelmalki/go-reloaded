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
		output = fixSingleQuotes(output)
		fmt.Println(output)

	} else if len(os.Args) != 3 {
		fmt.Println("error: u did not use 3 args")
	}
}

func fixSingleQuotes(s string) string {
    rs := strings.Fields(s)
    for i := 0; i < len(rs); i++ {
        if rs[i] == "'" {
            if i > 0 && i < len(rs)-1 {
                // If ' is not the first or last word
                rs[i-1] = strings.TrimRight(rs[i-1], ",.:;!?") + "'"
                rs[i+1] = "'" + strings.TrimLeft(rs[i+1], ",.:;!?")
                rs = append(rs[:i], rs[i+1:]...)
            } else {
                // If ' is the first or last word, remove it
                rs = append(rs[:i], rs[i+1:]...)
            }
        }
    }
    return strings.Join(rs, " ")
}