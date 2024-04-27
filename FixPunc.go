// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	if len(os.Args) == 3 {
// 		fileName := os.Args[1]
// 		file, err := os.ReadFile(fileName)
// 		if err != nil {
// 			fmt.Printf("Error")
// 			return
// 		}

// 		splitarray := strings.Split(string(file), "\n")
// 		splitarray = strings.Split(string(file), " ")
// 		//splitarray = Formspaces(splitarray)
// 		output := strings.Join(splitarray, " ")
// //		 output = Formspaces(output)
// 		output = FixPunc(output)
// 		// err = os.WriteFile(os.Args[2], []byte(output), 0644)
// 		// if err != nil {
// 		// 	fmt.Printf("Error")
// 		// 	return
// 		fmt.Printf(output)
// 	} else if len(os.Args) != 3 {
// 		fmt.Printf("error: u did not use 3 args")
// 	}
// }

// // to fix the punctuations
// func FixPunc(s string) string {
// 	rs := ""
// 	maxlen := len(s)
// 	for i := 0; i < maxlen; i++ {
// 		char := rune(s[i])
// 		if char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';' {
// 			if rune(rs[len(rs)-1]) == ' ' {
// 				rs = rs[:len(rs)-1]
// 				rs += string(char)
// 			} else {
// 				rs += string(char)
// 			}
// 			if i < maxlen-1 {
// 				if rune(s[i+1]) != ' ' {
// 					rs += " "
// 				}
// 			}
// 		} else {
// 			rs += string(char)
// 		}

// 	}
// 	return rs
// }
