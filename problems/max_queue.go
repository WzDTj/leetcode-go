/**
 * 面试题59 - II. 队列的最大值
 * https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/
 **/

package problems

// import (
//         "fmt"
// )

type MaxQueue struct {
	Max   []int
	Queue []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.Queue) == 0 {
		return -1
	}

	return this.Max[0]
}

func (this *MaxQueue) Push_back(value int) {
	// fmt.Println("--Start--")
	for i := 0; i < len(this.Max); i++ {
		if this.Max[i] <= value {
			// fmt.Println("before:", this.Max)
			this.Max = append(this.Max[:i], this.Max[i+1:]...)
			i--
			// fmt.Println("after:", this.Max)
		}
	}
	// fmt.Println("--End--")
	this.Max = append(this.Max, value)
	this.Queue = append(this.Queue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Queue) == 0 {
		return -1
	}

	value := this.Queue[0]

	if this.Max[0] == value {
		this.Max = this.Max[1:]
	}

	this.Queue = this.Queue[1:]

	return value
}
