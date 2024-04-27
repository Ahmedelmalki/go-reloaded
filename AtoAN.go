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
// 		// here where we will call our funcs :)
// 		splitarray := strings.Split(string(file), "\n")
// 		splitarray = strings.Split(string(file), " ")
// 		splitarray = AtoAN(splitarray)
// 		output := strings.Join(splitarray, " ")
// 		// err = os.WriteFile(os.Args[2], []byte(output), 0o644)
// 		// if err != nil {
// 		// 	fmt.Printf("Error")
// 		// 	return
// 			fmt.Println(output)
// 		}

// 	// } else if len(os.Args) != 3 {
// 	// 	fmt.Println("error: u did not use 3 args")
// 	// }
// }
// func isVowel(s string) bool{
	
// 		char := s[0]
// 		if char == 'a' || char == 'e' || char == 'u' || char == 'o' || char == 'i' || char == 'A' || char == 'E' || char == 'U' || char == 'O' || char == 'I' {
// 			return true
// 		} 
	
// 	return false
// }

// func AtoAN(s []string) []string {
// 	for i:=0 ; i<len(s) ; i++ {
// 		if (s[i]=="A" || s[i]=="a") {
// 			if i!=len(s)-1 &&  isVowel(s[i+1]){
// 				s[i]= s[i]+"n"
// 			}
// 		}
// 		if s[i]=="An" || s[i]=="an" {
// 			if i!=len(s)-1 &&  !isVowel(s[i+1]){
// 				s[i]= string(s[i][0])
// 			}
// 		}
// 	}
// 	return s
// }
