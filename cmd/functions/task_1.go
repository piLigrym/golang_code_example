package functions

import (
	"fmt"
	"regexp"
	"strconv"
)

/*
TestValidity takes the string as an input, and returns boolean flag true if the given string complies with the format,
or false if the string does not comply.

Expected format:
Sequence of numbers followed by dash followed by text, eg: 23-ab-48-caba-56-haha
-The numbers can be assumed to be unsigned integers
-The strings can be assumed to be ASCII sequences of arbitrary length larger than 0 (empty strings not allowed).

Estimate to complete: 30m.
Completed in: 37m.
*/
func TestValidity(input string) bool {
	// check input string if match with format.
	result, err := regexp.MatchString(`^(\d+-[a-zA-Z]+-*)+$`, input)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return result
}

/*
AverageNumber takes the string, and returns the average number from all the number.

Estimate to complete: 30m.
Completed in: 16m.
*/
func AverageNumber(input string) float64 {
	r := regexp.MustCompile("[0-9]+")
	numbers := r.FindAllString(input, -1)
	sum := 0.0
	count := len(numbers)
	for i := 0; i < count; i++ {
		number, err := strconv.ParseFloat(numbers[i], 8)
		if err != nil {
			panic("String with number contains not a string.")
		}
		sum += number
	}
	average := sum / float64(count)
	return average
}

/*
WholeStory takes the string, and returns a text that is composed from all the text words separated by spaces,
e.g. story called for the string 1-hello-2-world should return text: "hello world".

Estimate to complete: 30m.
Completed in: 14m.
*/
func WholeStory(input string) string {
	r := regexp.MustCompile("[a-zA-Z]+")
	words := r.FindAllString(input, -1)
	finalString := ""
	count := len(words)
	for i := 0; i < count; i++ {
		if i == count-1 {
			finalString += words[i]
		} else {
			finalString += fmt.Sprintf("%s ", words[i])
		}
	}
	return finalString
}
