/**
 * 225. 用队列实现栈
 * https://leetcode-cn.com/problems/implement-stack-using-queues/
 **/

package problems

type MyStack struct {
	List []int
	Size int
}

/** Initialize your data structure here. */
func StackConstructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.List = append(this.List, x)
	this.Size++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	popValue := this.List[this.Size-1]
	this.List = this.List[:this.Size-1]
	this.Size--
	return popValue
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.List[this.Size-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.List) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
