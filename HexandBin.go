package goreloaded

import (
	"fmt"
	"strconv"
)

func isHex(s string) bool {
	_, err := strconv.ParseUint(s, 16, 64)
	return err == nil
}

func isBin(s string) bool {
	_, err := strconv.ParseUint(s, 2, 64)
	return err == nil
}

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
