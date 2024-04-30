package goreloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UpLowandCap(arr []string) []string {
	for i := 0; i <= len(arr)-1; i++ {
		fmt.Println(arr[i])
		// for changging more than one string
		if arr[i] == "(up," {
			if i != len(arr)-1 {
				arr[i+1] = arr[i+1][:len(arr[i+1])-1]
				indix, _ := strconv.Atoi(arr[i+1])
				fmt.Println(arr, indix)
				for j := 1; j <= indix && i-j >= 0; j++ {
					fmt.Println(j)
					arr[i-j] = strings.ToUpper(arr[i-j])
				}
				arr[i] = ""
				arr[i+1] = ""
				arr = Formspaces(arr)
				continue
			} else {
				fmt.Println("index out of range")
				os.Exit(1)
			}
		}
		if arr[i] == "(low," && arr[0] != "(low," {
			if i != len(arr)-1 {
				arr[i+1] = arr[i+1][:len(arr[i+1])-1]
				indix, _ := strconv.Atoi(arr[i+1])
				for j := 1; j <= indix && i-j >= 0; j++ {
					arr[i-j] = strings.ToLower(arr[i-j])
				}
				arr[i] = ""
				arr[i+1] = ""
				arr = Formspaces(arr)
				continue
			} else {
				fmt.Println("index out of range")
				os.Exit(1)
			}
		}
		if arr[i] == "(cap," {
			if i != len(arr)-1 {
				arr[i+1] = arr[i+1][:len(arr[i+1])-1]
				indix, _ := strconv.Atoi(arr[i+1])
				for j := 1; j <= indix && i-j > 0; j++ {
					arr[i-j] = strings.ToUpper(arr[i-j][:1]) + strings.ToLower(arr[i-j][1:])
				}
				arr[i] = ""
				arr[i+1] = ""
				arr = Formspaces(arr)
				continue
			} else {
				fmt.Println("index out of range")
				os.Exit(1)
			}
		}
		if arr[i] == "(up)" {
			if i > 0 {
				arr[i-1] = strings.ToUpper(arr[i-1])
				arr[i] = ""
				arr = Formspaces(arr)
				continue
			}
		}
		if arr[i] == "(low)" {
			print("yes")
			if i > 0 {
				arr[i-1] = strings.ToLower(arr[i-1])
				arr[i] = ""
				arr = Formspaces(arr)
				continue
			} else {
				fmt.Println(" index out of range ")
				os.Exit(1)
			}
		}
		if arr[i] == "(cap)" {
			if i != 0 {
				arr[i-1] = strings.ToUpper(arr[i-1][:1]) + strings.ToLower(arr[i-1][1:])
				arr[i] = ""
				arr = Formspaces(arr)
				continue
			} else {
				fmt.Println(" index out of range ")
				os.Exit(1)
			}
		}
	}
	return arr
}
