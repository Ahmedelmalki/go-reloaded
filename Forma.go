package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error")
	}

	fileName := os.Args[1]
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error")
		return
	}
	splitarray := strings.Split(string(file), "\n")
	splitarray = strings.Split(string(file), " ")
	splitarray = forma(splitarray)
	fmt.Println(splitarray[3])
	test := strings.Join(splitarray, " ")
	//splitarray = strings.Split(test, " ")
	//fmt.Println(test)
	

	// for i := 0; i < len(slice); i++  {
	// fmt.Print(slice[i])
	// }


	output := strings.Join(splitarray, " ")
	err = os.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error")
		return
	}
}

func forma(s []string) {
	slice := []string{}
	word := ""
	for _,char := range test {
		if char != ' ' {
			word = word + string(char)
		} else {
			slice = append(slice , word, " ")
			word = ""
		}
	}
	if word != "" {
		slice = append(slice , word, " ")
			word = ""
	}
	for i := 0; i < len(slice); i++  {
		fmt.Print(slice[i])
		}
		fmt.Println()

	for i := 0; i < len(slice)-1; i++ {
		if (slice[i] == " " )&& (slice[i+1] == " " || slice[i+1] == "" ) {
			slice[i] = ""
		}else {
			continue
		}
	}
}