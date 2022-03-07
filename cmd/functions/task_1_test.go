package functions

import "testing"

func TestTestValidityValidateStringSuccessful(t *testing.T) {
	result := TestValidity("23-ab-48-caba-56-haha")
	if result != true {
		t.Errorf("String should be valid.")
	}
}

func TestTestValidityValidateStringUnsuccessful(t *testing.T) {
	result := TestValidity("23-ab-48-")
	if result != false {
		t.Errorf("String should be invalid.")
	}
}

func TestAverageNumberCorrect(t *testing.T) {
	result := AverageNumber("23-ab-48-ret")
	expectedValue := 35.5
	if result != expectedValue {
		t.Errorf("Uncorrect number: %f expected %f", result, expectedValue)
	}
}

func TestWholeStoryCorrect(t *testing.T) {
	result := WholeStory("1-hello-2-world")
	expectedValue := "hello world"
	if result != expectedValue {
		t.Errorf("Unccorect story: %s expected %s", result, expectedValue)
	}
}
