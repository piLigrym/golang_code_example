package functions

import (
	"fmt"
	"regexp"
)

/*
TestValidity takes the string as an input, and returns boolean flag true if the given string complies with the format,
or false if the string does not comply.

Expected format:
Sequence of numbers followed by dash followed by text, eg: 23-ab-48-caba-56-haha
-The numbers can be assumed to be unsigned integers
-The strings can be assumed to be ASCII sequences of arbitrary length larger than 0 (empty strings not allowed).

Estimate to complete: 30m.
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
