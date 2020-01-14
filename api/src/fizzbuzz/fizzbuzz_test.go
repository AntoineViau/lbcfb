package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	output, err := FizzBuzz("fizz", "buzz", 3, 5, 5)
	expected := []string{"1", "2", "fizz", "4", "buzz"}
	if err != nil {
		t.Errorf("err should be nil")
	}
	if len(output) != len(expected) {
		t.Errorf("Invalid output length, expected %d, got %d", len(expected), len(output))
	}
	for i := 0; i < len(output); i++ {
		if output[i] != expected[i] {
			t.Errorf("Invalid string for index %d, expected %s, given %s", i, expected[i], output[i])
		}
	}
}

func TestInvalidLimit(t *testing.T) {
	output, err := FizzBuzz("fizz", "buzz", 3, 5, -5)
	if err == nil {
		t.Errorf("err should not be nil")
	}
	if output != nil {
		t.Errorf("output should not be nil")
	}
}
func TestInvalidInt1(t *testing.T) {
	output, err := FizzBuzz("fizz", "buzz", 0, 5, 5)
	if err == nil {
		t.Errorf("err should not be nil")
	}
	if output != nil {
		t.Errorf("output should not be nil")
	}
}

func TestInvalidInt2(t *testing.T) {
	output, err := FizzBuzz("fizz", "buzz", 3, 0, 5)
	if err == nil {
		t.Errorf("err should not be nil")
	}
	if output != nil {
		t.Errorf("output should not be nil")
	}
}
