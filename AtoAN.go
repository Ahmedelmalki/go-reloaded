package goreloaded

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
