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
	//splitarray = formspaces(splitarray)
	output := strings.Join(splitarray, " ")
	err = os.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error")
		return
	}
}
// checking whether i-1 is convertible or not
func isHex(s string) bool {
    _, err := strconv.ParseUint(s, 16, 64)
    return err == nil
}

func isBin(s string) bool {
    _, err := strconv.ParseUint(s, 2, 64)
    return err == nil
}
// convirting
func HexandBin(s []string) []string {
	arr := s
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == "(hex)" && isHex(arr[i-1]) {
			bt, _ := strconv.ParseInt(arr[i-1], 16, 64)
			arr[i-1] = strconv.Itoa(int(bt))
			arr[i] = ""
		} else if arr[i] == "(bin)" && isBin(arr[i-1]) {
			bt, _ := strconv.ParseInt(arr[i-1], 2, 64)
			arr[i-1] = strconv.Itoa(int(bt))
			arr[i] = ""
		} else {
			fmt.Print("")
		}
	}
	return arr
}


