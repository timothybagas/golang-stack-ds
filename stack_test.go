package stack

import "testing"

func TestStack_Push(ts *testing.T) {
	ts.Run("push one value", func(t *testing.T) {
		var (
			st    = NewStack()
			value = 3
		)
		st.Push(value)

		if ret, ok := st.Pop().(int); !ok || ret != value {
			t.Errorf(
				"st.Pop() is not returning the correct value, expected %v but got %v",
				value,
				ret,
			)
		}
	})
	ts.Run("push more than one value", func(t *testing.T) {
		var (
			st     = NewStack()
			values = []int{3, 1, 2}
		)
		st.Push(values[0], values[1], values[2])

		for i := len(values) - 1; i >= 0; i-- {
			if ret, ok := st.Pop().(int); !ok || ret != values[i] {
				t.Errorf(
					"st.Pop() is not returning the correct value, expected %v but got %v",
					values[i],
					ret,
				)
			}
		}
	})
}

func TestStack_Pop(t *testing.T) {
	var (
		st = NewStack()
	)
	if ret := st.Pop(); ret != nil {
		t.Errorf(
			"st.Pop() is not returning the correct value, expected %v but got %v",
			nil,
			ret,
		)
	}
}

func TestStack_Peek(t *testing.T) {
	var (
		st     = NewStack()
		values = []int{3, 1}
	)
	for _, value := range values {
		st.Push(value)
		if ret, ok := st.Peek().(int); !ok || ret != value {
			t.Errorf(
				"st.Peek() is not returning the correct value, expected %v but got %v",
				value,
				ret,
			)
		}
	}
}

func TestStack_Length(t *testing.T) {
	var (
		st     = NewStack()
		values = []int{3, 1, 2}
	)
	st.Push(values[0], values[1], values[2])

	if length := st.Length(); length != len(values) {
		t.Errorf(
			"st.Peek() is not returning the correct value, expected %v but got %v",
			len(values),
			length,
		)
	}
}
