package goreloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UpLowandCap(arr []string) []string {
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == "(up," {
			if i != len(arr)-1 {
				arr[i+1] = arr[i+1][:len(arr[i+1])-1]
				indix, _ := strconv.Atoi(arr[i+1])
				for j := 1; j <= indix && i-j >= 0; j++ {
					arr[i-j] = strings.ToUpper(arr[i-j])
				}
				arr[i] = ""
				arr[i+1] = ""
				arr = Formspaces(arr)
				i--
				continue
			} else {
				fmt.Println("index out of range")
				os.Exit(1)
			}
		} else if arr[i] == "(low," && arr[0] != "(low," {
			if i != len(arr)-1 {
				arr[i+1] = arr[i+1][:len(arr[i+1])-1]
				indix, _ := strconv.Atoi(arr[i+1])
				for j := 1; j <= indix && i-j >= 0; j++ {
					arr[i-j] = strings.ToLower(arr[i-j])
				}
				arr[i] = ""
				arr[i+1] = ""
				arr = Formspaces(arr)
				i--
				continue
			} else {
				fmt.Println("index out of range")
				os.Exit(1)
			}
		} else if arr[i] == "(cap," {
			if i != len(arr)-1 {
				arr[i+1] = arr[i+1][:len(arr[i+1])-1]
				indix, _ := strconv.Atoi(arr[i+1])
				for j := 1; j <= indix && i-j > 0; j++ {
					arr[i-j] = strings.ToUpper(arr[i-j][:1]) + strings.ToLower(arr[i-j][1:])
				}
				arr[i] = ""
				arr[i+1] = ""
				arr = Formspaces(arr)
				i--
				continue
			} else {
				fmt.Println("index out of range")
				os.Exit(1)
			}
		} else if arr[i] == "(up)" {
			if i > 0 {
				arr[i-1] = strings.ToUpper(arr[i-1])
				arr[i] = ""
				arr = Formspaces(arr)
				i--
				continue
			}
		} else if arr[i] == "(low)" {
			if i > 0 {
				arr[i-1] = strings.ToLower(arr[i-1])
				arr[i] = ""
				arr = Formspaces(arr)
				i--
				continue
			} else {
				fmt.Println(" index out of range ")
				os.Exit(1)
			}
		} else if arr[i] == "(cap)" {
			if i != 0 {
				arr[i-1] = strings.ToUpper(arr[i-1][:1]) + strings.ToLower(arr[i-1][1:])
				arr[i] = ""
				arr = Formspaces(arr)
				i--
				continue
			} else {
				fmt.Println(" index out of range ")
				os.Exit(1)
			}
		}
	}
	return arr
}
