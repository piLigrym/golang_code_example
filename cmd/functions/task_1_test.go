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
