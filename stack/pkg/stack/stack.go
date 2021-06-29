package stack

import (
	"fmt"
)

// Stack struct
type Stack struct {
	list []interface{}
}

// New stack
func New() *Stack {
	s := new(Stack)
	s.list = make([]interface{}, 0, 8)

	return s
}

// IsEmpty empty or not
func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}

// Pop last element
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	el := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return el
}

// Push element to last
func (s *Stack) Push(el interface{}) {
	s.list = append(s.list, el)
}

// ToSlice get stack list
func (s *Stack) ToSlice() []interface{} {
	return s.list
}

func (s *Stack) String() string {
	str := ""
	for i := 0; i < len(s.list); i++ {
		str += fmt.Sprint(s.list[i])
	}

	return str
}
