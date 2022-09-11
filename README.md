# Stack Data Structure Implementation in Golang (thread safe)

Sample code:
```
package main

import (
    stack "github.com/timothybagas/golang-stack-ds"
)

func main() {
    // creates an empty stack
    st := stack.NewStack()
    
    // push element into the stack
    st.Push(1)
    st.Push(2)
    
    // pop an element from the stack
    st.Pop()
}
```