package errors

import (
	"fmt"
	"testing"
)

func TestErrTypeMismatch(t *testing.T) {
	expectedErrorMessage := "type mismatch: expected %s, got %s"
	if ErrTypeMismatch != expectedErrorMessage {
		t.Errorf("ErrTypeMismatch constant is incorrect, got: %s, want: %s", ErrTypeMismatch, expectedErrorMessage)
	}

	testExpectedType := "string"
	testGotType := "int"
	expectedFormattedError := "type mismatch: expected string, got int"
	formattedError := fmt.Errorf(ErrTypeMismatch, testExpectedType, testGotType).Error()

	if formattedError != expectedFormattedError {
		t.Errorf("Formatted error message is incorrect, got: %s, want: %s", formattedError, expectedFormattedError)
	}
}
