package stack

import "sync"

type Stack interface {
	Push(vs ...interface{})
	Pop() interface{}
	Peek() interface{}
	Length() int
	Clear()
	IsEmpty() bool
	Copy() Stack
}

type stack struct {
	head   *node
	length int
	sync.RWMutex
}

func NewStack(vs ...interface{}) Stack {
	st := new(stack)
	st.Push(vs...)
	return st
}

func (st *stack) Push(vs ...interface{}) {
	st.Lock()
	defer st.Unlock()

	for _, v := range vs {
		nd := newNode(nil, v)
		if st.head != nil {
			nd.next = st.head
		}
		st.head = nd
		st.length++
	}
}

func (st *stack) Pop() interface{} {
	st.Lock()
	defer st.Unlock()

	if st.head == nil {
		return nil
	}
	val := st.head.val
	st.head = st.head.next
	st.length--

	return val
}

func (st *stack) Peek() interface{} {
	st.RLock()
	defer st.RUnlock()

	if st.head == nil {
		return nil
	}
	return st.head.val
}

func (st *stack) Length() int {
	st.RLock()
	defer st.RUnlock()
	return st.length
}

func (st *stack) Clear() {
	st.Lock()
	defer st.Unlock()
	st.head = nil
	st.length = 0
}

func (st *stack) IsEmpty() bool {
	st.RLock()
	defer st.RUnlock()
	return st.head == nil && st.length == 0
}

func (st *stack) Copy() Stack {
	st.RLock()
	defer st.RUnlock()

	var (
		values = make([]interface{}, 0)
		ptr    = st.head
	)
	for ptr != nil {
		values = append(values, ptr.val)
		ptr = ptr.next
	}
	var (
		newStack = NewStack()
	)
	for i := len(values) - 1; i >= 0; i-- {
		newStack.Push(values[i])
	}
	return newStack
}
