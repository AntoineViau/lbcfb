package fizzbuzz

import (
	"errors"
	"strconv"
)

// FizzBuzz is the core business job:
// counting from 1 to limit, it displays string1 and/or string2
// when the counter can be divided by int1 and/or int2.
func FizzBuzz(string1 string, string2 string, int1 int, int2 int, limit int) ([]string, error) {
	if limit < 0 {
		return nil, errors.New("limit must greater or equal to 0")
	}
	if int1 == 0 {
		return nil, errors.New("int1 can't be 0")
	}
	if int2 == 0 {
		return nil, errors.New("int2 can't be 0")
	}
	var output []string
	for i := 1; i <= limit; i++ {
		str := ""
		if i%int1 == 0 {
			str += string1
		}
		if i%int2 == 0 {
			str += string2
		}
		if str == "" {
			str = strconv.Itoa(i)
		}
		output = append(output, str)
	}
	return output, nil
}
