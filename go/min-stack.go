package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := Constructor()
	(&stack).Push(-2)
	(&stack).Push(0)
	(&stack).Push(-3)
	fmt.Println((&stack).GetMin())
	(&stack).Pop()
	fmt.Println((&stack).Top())
	fmt.Println((&stack).GetMin())
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type MinStack struct {
	s1, s2 *Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		s1: NewStack(),
		s2: NewStack(),
	}
}

func (this *MinStack) Push(x int) {
	this.s1.Push(x)
	if this.s2.Empty() || x <= this.s2.Peak().(int) {
		this.s2.Push(x)
	}
}

func (this *MinStack) Pop() {
	x := this.s1.Pop()
	if x == this.s2.Peak().(int) {
		this.s2.Pop()
	}

}

func (this *MinStack) Top() int {
	return this.s1.Peak().(int)

}

func (this *MinStack) GetMin() int {
	return this.s2.Peak().(int)

}
