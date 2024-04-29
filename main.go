package main

import (
	"fmt"
	"os"
	"strconv"
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
		// splitarray := strings.Split(string(file), "\n")
		splitarray := strings.Split(string(file), " ")
		splitarray1 := Formspaces(splitarray)
		splitarray2 := upLowandCap(splitarray1)
		splitarray3 := HexandBin(splitarray2)
		splitarray4 := AtoAN(splitarray3)
		output := strings.Join(splitarray4, " ")
		output1 := FixPunc(output)
		fmt.Println(output1)
		output2 := SingleQoutse(output1)
		err = os.WriteFile(os.Args[2], []byte(output2), 0o644)
		fmt.Println(output2)
		if err != nil {
			fmt.Printf("Error")
			return
		}

	} else if len(os.Args) != 3 {
		fmt.Println("error: u did not use 3 args")
	}
}

//---------------------------------------------------------------------------------------------

func upLowandCap(arr []string) []string {
	for i := 0; i <= len(arr)-1; i++ {
		// for changging more than one string
		if arr[i] == "(up," {
			arr[i+1] = arr[i+1][:len(arr[i+1])-1]
			if i != len(arr)-1 {
				indix, _ := strconv.Atoi(arr[i+1])
				fmt.Println(arr, indix)
				for j := 1; j <= indix && i-j >= 0; j++ {
					fmt.Println(j)
					arr[i-j] = strings.ToUpper(arr[i-j])
				}
				arr[i] = ""
				arr[i+1] = ""
			} else {
				fmt.Println("index out of range")
				os.Exit(1)
			}
		}
		if arr[i] == "(low," && arr[0] != "(low," {
			arr[i+1] = arr[i+1][:len(arr[i+1])-1]
			if i != len(arr)-1 {
				indix, _ := strconv.Atoi(arr[i+1])
				for j := 1; j <= indix && i-j >= 0; j++ {
					arr[i-j] = strings.ToLower(arr[i-j])
				}
				arr[i] = ""
				arr[i+1] = ""
			} else {
				fmt.Println("index out of range")
				os.Exit(1)
			}
		}
		if arr[i] == "(cap," {
			arr[i+1] = arr[i+1][:len(arr[i+1])-1]
			if i != len(arr)-1 {
				indix, _ := strconv.Atoi(arr[i+1])
				for j := 1; j <= indix && i-j > 0; j++ {
					arr[i-j] = strings.ToUpper(arr[i-j][:1]) + strings.ToLower(arr[i-j][1:])
				}
				arr[i] = ""
				arr[i+1] = ""
			} else {
				fmt.Println("index out of range")
				os.Exit(1)
			}
		}
		if arr[i] == "(up)" && arr[0] != "(up)" {
			arr[i-1] = strings.ToUpper(arr[i-1])
			arr[i] = ""
		}
		if arr[i] == "(low)" && arr[0] != "(low)" {
			arr[i-1] = strings.ToLower(arr[i-1])
			arr[i] = ""
		}
		if arr[i] == "(cap)" {
			if arr[0] != "(cap)" {
				arr[i-1] = strings.ToUpper(arr[i-1][:1]) + strings.ToLower(arr[i-1][1:])
				arr[i] = ""
			} else {
				fmt.Println(" index out of range ")
				os.Exit(1)
			}
		}
	}
	return arr
}

//____________________________________________________________________________________//

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

//________________________________________________________________________________//

// to fix the punctuations
func FixPunc(s string) string {
	rs := ""
	maxlen := len(s)
	for i := 0; i < maxlen; i++ {
		char := rune(s[i])
		if char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';' {
			if rune(rs[len(rs)-1]) == ' ' {
				rs = rs[:len(rs)-1]
				rs += string(char)
			} else {
				rs += string(char)
			}
			if i < maxlen-1 {
				if rune(s[i+1]) != ' ' {
					rs += " "
				}
			}
		} else {
			rs += string(char)
		}

	}
	return rs
}

//___________________________________________________________________________________________________________//

func isVowel(s string) bool {
	char := s[0]
	if char == 'a' || char == 'e' || char == 'u' || char == 'o' || char == 'i' || char == 'A' || char == 'E' || char == 'U' || char == 'O' || char == 'I' {
		return true
	}

	return false
}

func AtoAN(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "A" || s[i] == "a" {
			if i != len(s)-1 && isVowel(s[i+1]) {
				s[i] = s[i] + "n"
			}
		}
		if s[i] == "An" || s[i] == "an" {
			if i != len(s)-1 && !isVowel(s[i+1]) {
				s[i] = string(s[i][0])
			}
		}
	}
	return s
}

//______________________________________________________________________________________________________//

func Formspaces(arr []string) []string {
	result := []string{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			result = append(result, arr[i])
		}
	}
	return result
}

//--------------------------------------------------------------------

func SingleQoutse(s string) string {
	rs := []rune(s)
	b := true
	for i := 0; i < len(rs); i++ {
		if b && i > 0 && (i < len(rs)-1 ) && (rs[i] == '\'' && rs[i-1] == ' ') {
			if rs[i+1] == ' ' {
				fmt.Println("2222")
				rs = append(rs[:i+1], rs[i+2:]...)
				b = false
				continue
			}
		} else if !b && rs[i] == '\'' && rs[i-1] == ' ' && i <= len(rs)-1 {
			if rs[i-1] == ' ' {
				fmt.Println("3333333333")
				rs = append(rs[:i-1], rs[i:]...)
				b = true
				continue
			}
		}else if !b && rs[0] == '\'' && rs[1] == ' ' {
			fmt.Println("111111")
			rs = append(rs[:1], rs[2:]...)
			continue
		}
	}
	return string(rs)
}
