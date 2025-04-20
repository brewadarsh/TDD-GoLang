package stack

import "testing"

func TestStack(t *testing.T) {
	intStack := new(Stack[int])
	// Check if the stack is empty or not.
	assertCondition(t, true, intStack.IsEmpty())

	// Check if the Pop returns false.
	_, exists := intStack.Pop()
	assertCondition(t, false, exists)

	intStack.Push(20)
	intStack.Push(25)

	// Check if the Pop returns true and the returned value is expected or not.
	lastValue, exists := intStack.Pop()
	assertCondition(t, true, exists)
	assertResponse(t, lastValue, 25)
}

func assertCondition(t *testing.T, expected, got bool) {
	t.Helper()
	if got != expected {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func assertResponse(t *testing.T, expected, got int) {
	t.Helper()
	if got != expected {
		t.Errorf("got %v but expected %v", got, expected)
	}
}
