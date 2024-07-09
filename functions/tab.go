package ascii

// Tab handles the control character /t which print double space
func Tab(s string) string {
	stringLength := len(s)
	result := ""

	for i := 0; i < stringLength; i++ {
		if s[i] == '\\' && s[i+1] == 't' {
			result = result + "  "
			i++
		} else {
			result += string(s[i])
		}
	}
	return result
}
