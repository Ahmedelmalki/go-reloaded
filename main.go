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
	splitarray = changer(splitarray)

	output := strings.Join(splitarray, " ")
	err = os.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error")
		return
	}
	// fmt.Print(output)
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

//checking if the input is valid or not
func isHex(s string) bool {
	_, err := strconv.ParseInt(s, 16, 64)
	return err == nil
}

func isBin(s string) bool {
	_, err := strconv.ParseInt(s, 2, 64)
	return err == nil
}

// UP and low
func changer(arr []string) []string {
	//for basic case
	for i := 0; i < len(arr); i++ {
		if arr[i] == "(hex)" {
			if isHex(arr[i-1]) {
				bt, _ := strconv.ParseInt(arr[i-1], 16, 64)
				arr[i-1] = strconv.Itoa(int(bt))
				arr[i] = ""
			} else {
				fmt.Println("Error: invalid (hex) input")
				return arr
			}
		}
		if arr[i] == "(bin)" {
			if isBin(arr[i-1]) {
				bt, _ := strconv.ParseInt(arr[i-1], 2, 64)
				arr[i-1] = strconv.Itoa(int(bt))
				arr[i] = ""
			} else {
				fmt.Println("Error: invalid (bin) input")
				return arr
			}
		}
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
				arr[i-j] = strings.Title(arr[i-j])
			}
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

// for handling punctuations
func fixpunc(s string) string {
	arr := []rune(s)
	for i := 0; i< len(arr) ; i++ {
		if arr[i] == ' ' && (arr[i+1] == ',' || arr[i+1] == '.' || arr[i+1] == ' ' || arr[i+1] == '?' || arr[i+1] == '!' || arr[i+1] == ':' || arr[i+1] == ';') {
			arr[i] = ''
		}
	}
	return string(arr)
}
