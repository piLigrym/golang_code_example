package functions

import "testing"

func TestGenerateRandomInputStringCorrect(t *testing.T) {
	input := GenerateRandomInputString(true)
	if TestValidity(input) != true {
		t.Errorf("String not correct %s", input)
	}
}

func TestGenerateRandomInputStringNotCorrect(t *testing.T) {
	input := GenerateRandomInputString(false)
	if TestValidity(input) != false {
		t.Errorf("String correct %s", input)
	}
}
