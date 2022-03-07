package functions

import (
	"math/rand"
	"strconv"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

/*
GenerateRandomInputString takes boolean flag and generates random correct strings
if the parameter is true and random incorrect strings if the flag is false.

Estimate to complete: 40m.
Completed in: 43m.
*/
func GenerateRandomInputString(isCorrect bool) string {
	count := rand.Intn(100)
	input := ""
	for i := 0; i < count; i++ {
		input += strconv.Itoa(rand.Intn(100))
		input += "-"
		randomString := ""
		for randomString == "" {
			randomString = RandomString(rand.Intn(20))
		}
		input += randomString

		if i+1 != count {
			input += "-"
		}
	}

	if isCorrect == false {
		input += RandomString(rand.Intn(20)) + "-" + RandomString(rand.Intn(20)) + "-" + RandomString(rand.Intn(20))
	}

	return input
}
