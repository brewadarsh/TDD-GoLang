package stack

// The interface for Stack Data Structure.
type Stack[T any] struct {
	values []T
}

// Push a value into the Stack in FIFO fashion.
func (stack *Stack[T]) Push(value T) {
	stack.values = append(stack.values, value)
}

// Pop a value from the Stack in FIFO fashion.
func (stack *Stack[T]) Pop() (T, bool) {
	if stack.IsEmpty() {
		var zero T
		return zero, false
	}
	lastValue := stack.values[len(stack.values)-1]
	return lastValue, true
}

// Check if the stack is empty or not.
func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.values) == 0
}
