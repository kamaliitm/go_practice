package main

type MinStackNode struct {
	val  int
	min  int
	next *MinStackNode
}

type MinStack struct {
	head *MinStackNode
}

func Constructor() MinStack {
	return MinStack{
		head: nil,
	}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = &MinStackNode{
			val: val,
			min: val,
		}
		return
	}

	min := val
	if min > this.head.min {
		min = this.head.min
	}

	this.head = &MinStackNode{
		val:  val,
		min:  min,
		next: this.head,
	}
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}
