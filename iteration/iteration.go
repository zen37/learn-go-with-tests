package iteration

func main() {

}

func repeat(character string, times int) string {
	var s string
	for i := 1; i <= times; i++ {
		s += character
	}
	return s
}
