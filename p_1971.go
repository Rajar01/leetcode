package main

// Stack

type stack struct {
	top    *node
	length int
}

type node struct {
	value interface{}
	prev  *node
}

func NewStack() *stack {
	return &stack{nil, 0}
}

func (s *stack) Push(value interface{}) {
	n := &node{value: value, prev: s.top}
	s.top = n
	s.length++
}

func (s *stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--

	return n.value
}

func (s *stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}

	return s.top.value
}

func (s *stack) Len() int {
	return s.length
}

// End of stack

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	if n == 1 && source == 0 && destination == 0 {
		return true
	}

	if n == 1 && source != 0 && destination != 0 {
		return false
	}

	// Graph implementation using adjacent list

	al := map[int][]int{}

	for i := 1; i <= n; i++ {
		al[i] = []int{}
	}

	for _, v := range edges {
		al[v[0]] = append(al[v[0]], v[1])
		al[v[1]] = append(al[v[1]], v[0])
	}

	// DFS implementation

	s := NewStack()
	s.Push(source)

	visited := make([]bool, n)

	for s.Len() > 0 {
		cv := s.Pop()

		if cv, ok := cv.(int); ok {
			if !visited[cv] {
				visited[cv] = true

				for _, v := range al[cv] {
					if v == destination {
						return true
					}

					if !visited[v] {
						s.Push(v)
					}
				}
			}
		}
	}

	return false
}
