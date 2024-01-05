package iteration

const repeatFactor = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatFactor; i++ {
		repeated += character
	}
	return repeated
}
