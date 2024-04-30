package goreloaded

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
