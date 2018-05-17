package gods

type Stack []interface{}

func NewStack() *Stack {
	return make(Stack, 0)
}

func (s *Stack) push(element interface{}) *Stack {
	append(s, element)
	return s
}

func (s *Stack) peek() interface{} {
	size := len(s)
	return s[size-1]
}

func (s *Stack) hasNext() bool {
	size := len(s)
	if size != 0 {
		return true
	}
	return false
}

func (s *Stack) pop() interface{} {
	size := len(s)
	value := s[size-1]
	newStack := NewStack()
	append(newStack, s[0:size-1])
	s = newStack
	return value
}
