package iteration

func Repeat(char string, qtdRepeat int) string {
	repeatedChar := ""
	for i := 0; i < qtdRepeat; i++ {
		repeatedChar += char
	}
	return repeatedChar
}
