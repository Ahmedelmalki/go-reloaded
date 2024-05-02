package goreloaded

import (
	"fmt"
	"os"
	"strconv"
)

func HexandBin(s []string) []string {
	for i := 0; i <= len(s)-1; i++ {
		if s[0] == "(hex)" || s[0] == "(bin)" {
			s[0] = ""
			fmt.Println("invalid input")
			os.Exit(0)
		}
		if s[i] == "(hex)" && i != 0 {
			bt, err := strconv.ParseInt(s[i-1], 16, 64)
			if err != nil {
				fmt.Println("invalid input")
				os.Exit(0)
			}
			s[i-1] = strconv.Itoa(int(bt))
			s[i] = ""
			s = append(s[:i], s[i+1:]...)
			i--
		}
		if s[i] == "(bin)" && i != 0 {
			bt, err := strconv.ParseInt(s[i-1], 2, 64)
			if err != nil {
				fmt.Println("invalid input")
				os.Exit(0)
			}
			s[i-1] = strconv.Itoa(int(bt))
			s[i] = ""
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}
