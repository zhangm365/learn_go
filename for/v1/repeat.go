package iteration

const n = 5

// Repeat return the character for 5 times
func Repeat(str string) (res string) {
	// var res string
	for i := 0; i < n; i++ {
		res += str
	}
	return
}
