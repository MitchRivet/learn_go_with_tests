package iteration

const repeatCount int = 5

func Repeat(c string) string {

	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += c

	}

	return repeated
}
