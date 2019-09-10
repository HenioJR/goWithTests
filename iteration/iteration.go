package iteration

func Repeat(char string) string {
	repeatedChar := ""
	for i := 0; i < 5; i++ {
		repeatedChar += char
	}
	return repeatedChar
}
