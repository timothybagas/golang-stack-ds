package stack

import "sync"

type Stack struct {
	head   *node
	length int
	sync.RWMutex
}

func NewStack(vs ...interface{}) *Stack {
	st := new(Stack)
	st.Push(vs...)
	return st
}

func (st *Stack) Push(vs ...interface{}) {
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

func (st *Stack) Pop() interface{} {
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

func (st *Stack) Peek() interface{} {
	st.RLock()
	defer st.RUnlock()

	if st.head == nil {
		panic("trying to pop from empty stack")
	}
	return st.head.val
}

func (st *Stack) Length() int {
	st.RLock()
	defer st.RUnlock()
	return st.length
}

func (st *Stack) Clear() {
	st.Lock()
	defer st.Unlock()
	st.head = nil
	st.length = 0
}

func (st *Stack) IsEmpty() bool {
	st.RLock()
	defer st.RUnlock()
	return st.head == nil && st.length == 0
}

func (st *Stack) Copy() *Stack {
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
