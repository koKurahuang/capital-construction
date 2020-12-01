package leetcode

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var save  = newStack()
	for i, _ := range tokens {
		switch tokens[i] {
		case "/":
			num1 := save.Pop()
			num2 := save.Pop()
			newone := num2 / num1
			save.Push(newone)
		case "*":
			num1 := save.Pop()
			num2 := save.Pop()
			newone := num2 * num1
			save.Push(newone)
		case "+":
			num1 := save.Pop()
			num2 := save.Pop()
			newone := num2 + num1
			save.Push(newone)
		case "-":
			num1 := save.Pop()
			num2 := save.Pop()
			newone := num2 - num1
			save.Push(newone)
		default:
			num, _ := strconv.Atoi(tokens[i])
			save.Push(num)
		}
	}
	return save.Pop()
}

type Stack struct {
	value []int
}

func newStack() *Stack {
	var s Stack
	s.value = make([]int, 0)
	return &s
}

func (s *Stack) Push(v int) {
	s.value = append(s.value, v)
}

func (s *Stack) Pop() int {
	ret := s.value[len(s.value) - 1]
	s.value = s.value[:len(s.value)-1]
	return ret
}



