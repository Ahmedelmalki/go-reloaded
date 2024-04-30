package goreloaded

func Formspaces(arr []string) []string {
	result := []string{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			result = append(result, arr[i])
		}
	}
	return result
}
