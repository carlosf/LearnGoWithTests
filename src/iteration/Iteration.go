package iteration

func Repeat(character string, numOfTimes int) string {
	var repeated string
	for i := 0; i < numOfTimes; i++ {
		repeated += character
	}
	return repeated
}
