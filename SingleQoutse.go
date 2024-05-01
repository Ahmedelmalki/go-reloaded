package goreloaded

func SingleQuote(s string) string {
	rs := []rune(s)
	isfirstQuote := true
	for i := 0; i < len(rs); i++ {
		if isfirstQuote && i > 0 && (i < len(rs)-1) && (rs[i] == '\'' && (rs[i-1] == ' ' || rs[i+1] == ' ')) {
			if rs[i+1] == ' ' {
				rs = append(rs[:i+1], rs[i+2:]...)
				isfirstQuote = false
				continue
			} else {
				isfirstQuote = false
				continue
			}
		} else if !isfirstQuote && i != 0 && i != len(rs)-1 && rs[i] == '\'' && (rs[i-1] == ' ' || rs[i+1] == ' ') {
			if rs[i-1] == ' ' {
				rs = append(rs[:i-1], rs[i:]...)
				isfirstQuote = true
				continue
			} else {
				isfirstQuote = true
				continue
			}
		} else if i == 0 && (rs[i] == '\'') && isfirstQuote && i != len(rs)-1 {
			if rs[i+1] == ' ' {
				rs = append(rs[:i], rs[i+1:]...)
				isfirstQuote = false
				continue
			} else {
				isfirstQuote = false
				continue
			}
		} else if rs[i] == '\'' && !isfirstQuote && i == len(rs)-1 && i != 0 {
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
