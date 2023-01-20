package iteration

// const n = 5

// Repeat return the character for n times
func Repeat(str string, n int) (res string) {
	// var res string
	for i := 0; i < n; i++ {
		res += str
	}
	return
}
