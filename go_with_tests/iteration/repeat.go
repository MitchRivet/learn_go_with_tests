package iteration

// can use std lib- strings.Repeat
func Repeat(c string, repeat int) string {

	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += c

	}

	return repeated
}
