package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
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
	splitarray = fixPunc(strings.Join(splitarray, " "))
	output := strings.Join(splitarray, " ")
	err = os.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error")
		return
	//fmt.Printf(output)
	}
	
		
	} else if len(os.Args) <= 2 {
		fmt.Printf("error: u did not use 3 args")
	} else {
		fmt.Printf("error: u did not use 3 args")
	}
}
func hasPunctuation(s string) bool {
	for _, r := range s {
		if unicode.IsPunct(r) {
			return true
		}
	}
	return false
}

func fixPunc(arr string) []string {
	result := []string{}
	for i:=1 ; i <len(arr) ; i++ {
		// if (arr[i+1] == "." || arr[i+1] == "." || arr[i+1] == "," || arr[i+1] == "!" || arr[i+1] == "?") && arr[i] == " "{
		// 	arr[i] = ""
		//fmt.Println("------------------------")
		//
		if unicode.IsPunct(rune(arr[i])) && arr[i]!= '\'' {
			result = append(result, string(arr[i+1]))
			fmt.Println(arr)
	}
}	
return result

}