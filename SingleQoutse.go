package goreloaded

import (
	"fmt"
)
func SingleQuote(s string) string {
	rs := []rune(s)
	isfirstQuote := true
	for i := 0; i < len(rs); i++ {
		if isfirstQuote && i > 0 && (i < len(rs)-1) && (rs[i] == '\'' && rs[i-1] == ' ') {
			if rs[i+1] == ' ' {
				fmt.Println("2222")
				rs = append(rs[:i+1], rs[i+2:]...)
				isfirstQuote = false
				continue
			}
		}
		if !isfirstQuote && (rs[i] == '\'' && rs[i-1] == ' ') && (i < len(rs)-1) {
			if rs[i+1] == ' ' {
				fmt.Println("3333333333")
				rs = append(rs[:i-1], rs[i:]...)
				isfirstQuote = true
				continue
			}
		}
		if len(rs) > 1 && (rs[0] == '\'' && rs[1] == ' ') && isfirstQuote {
			fmt.Println("444444444")
			rs = append(rs[:1], rs[2:]...)
			isfirstQuote = false
			continue
		}
		if rs[i] == '\'' && !isfirstQuote && i == len(rs)-1 && i != 0 {
			if rs[i-1] == ' ' {
				rs = append(rs[:i-1], rs[i:]...)
				isfirstQuote = true
				continue
			} else {
				isfirstQuote = true
				continue
			}
		}
	}
	return string(rs)
}