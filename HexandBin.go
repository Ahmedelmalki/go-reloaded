package main

import (
	"fmt"
	"os"
	"strconv"
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
	splitarray := strings.Split(string(file), " ")
	splitarray = HexandBin(splitarray)
	splitarray = formspaces(splitarray)
	output := strings.Join(splitarray, " ")
	err = os.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error")
		return
	}
}
func HexandBin(s []string) []string {
	arr := s
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == "(hex)" {
			bt, _ := strconv.ParseInt(arr[i-1], 16, 64)
			arr[i-1] = strconv.Itoa(int(bt))
			arr[i] = ""
		}
		if arr[i] == "(bin)" {
			bt, _ := strconv.ParseInt(arr[i-1], 2, 64)
			arr[i-1] = strconv.Itoa(int(bt))
			arr[i] = ""
		}
	}
	return arr
}

func formspaces(arr []string) []string {
	result := []string{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			result = append(result, arr[i])
		}
	}
	return result
}
