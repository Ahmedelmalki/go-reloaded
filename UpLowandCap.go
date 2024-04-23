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
	splitarray = upLowandCap(splitarray)
	output := strings.Join(splitarray, " ")
	err = os.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error")
		return
	}
}

// using trimatoi for converting from string to int
func TrimAtoi(s string) int {
	num := []rune{}
	ss := []rune(s)
	sign := 1
	b := false
	for i := 0; i < len(s); i++ {
		if ss[i] == '-' && b == false {
			sign = sign * -1
		}
		if s[i] <= '9' && ss[i] >= '0' {
			b = true
			num = append(num, ss[i])
		}
	}
	sum := 0
	for i := 0; i < len(num); i++ {
		sum = sum*10 + int(num[i]-'0')
	}
	return sum * sign
}

func upLowandCap(arr []string) []string {
	for i := 0; i <= len(arr)-1; i++ {
		//for changging more than one string
		if arr[i] == "(up," {
			indix := TrimAtoi(arr[i+1])
			for j := 1; j <= indix; j++ {
				arr[i-j] = strings.ToUpper(arr[i-j])
			}
			arr[i] = ""
			arr[i+1] = ""
		}
		if arr[i] == "(low," {
			indix := TrimAtoi(arr[i+1])
			for j := 1; j <= indix; j++ {
				arr[i-j] = strings.ToLower(arr[i-j])
			}
			arr[i] = ""
			arr[i+1] = ""
		}
		if arr[i] == "(cap," {
			indix := TrimAtoi(arr[i+1])
			for j := 1; j <= indix; j++ {
				arr[i-j] = strings.ToUpper(arr[i-j][:1])+strings.ToLower(arr[i-j][1:])
			}
			arr[i] = ""
			arr[i+1] = ""
		}
		if arr[i] == "(up)" {
			arr[i-1] = strings.ToUpper(arr[i-1])
			arr[i] = ""
		}
		if arr[i] == "(low)" {
			arr[i-1] = strings.ToLower(arr[i-1])
			arr[i] = ""
		}
		if arr[i] == "(cap)" {
			arr[i-1] = strings.Title(arr[i-1])
			arr[i] = ""
		}
	}
	return arr
}
