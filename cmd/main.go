package cmd

import (
	"github.com/oklookat/kp2imdb/util"
)

func NewStack(maxSize int) Stack {
	return Stack{
		MaxSize: maxSize,
	}
}

type Stack struct {
	MaxSize      int
	Stack        []string
	AlwaysBottom []string
}

func (s *Stack) Add(str string) int {
	s.Stack = append(s.Stack, str)
	return len(s.Stack) - 1
}

func (s *Stack) AddAlwaysBottom(str string) int {
	s.AlwaysBottom = append(s.AlwaysBottom, str)
	return len(s.AlwaysBottom) - 1
}

func (s *Stack) Render() {
	util.ClearTerminal()
	s.clearMax(s.Stack)
	s.clearMax(s.AlwaysBottom)
	for _, v := range s.Stack {
		println(v)
	}
	for _, v := range s.AlwaysBottom {
		println(v)
	}
}

func (s *Stack) clearMax(st []string) {
	for len(st) >= s.MaxSize {
		st = removeFromSlice(st, len(st)-1)
	}
}

func removeFromSlice[T any](slice []T, idx int) []T {
	if idx > len(slice)-1 || idx < 0 {
		return slice
	}
	return append(slice[:idx], slice[idx+1:]...)
}
