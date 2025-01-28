package iteration

import (
	"strings"
)

func Repeat(character string, times int) (repeated string) {
	repeated = ""
	for i := 0; i < times; i++ {
		repeated += character
	}
	return
}

func RepeatWithStringsBuilder(character string, times int) (repeated string) {
	var builder strings.Builder
	for i := 0; i < times; i++ {
		builder.WriteString(character)
	}
	repeated = builder.String()
	return
}
