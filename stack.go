package main

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	} else {
		item := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return item
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peek() string {
	return s.items[len(s.items)-1]
}
