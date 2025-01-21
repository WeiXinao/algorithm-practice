package main

import (
	"container/heap"
	"fmt"
)

type BigHeap []int

func (bh BigHeap) Len() int {
	return len(bh)
}

func (bh BigHeap) Less(i int, j int) bool {
	return bh[i] > bh[j]
}

func (bh BigHeap) Swap(i int, j int) {
	bh[i], bh[j] = bh[j], bh[i]
}

func (bh *BigHeap) Push(x any) {
	*bh = append(*bh, x.(int))
}

func (bh *BigHeap) Pop() any {
	n := len(*bh)
	x := (*bh)[n-1]
	*bh = (*bh)[:n-1]
	return x
}

type SmallHeap []int

func (sh SmallHeap) Len() int {
	return len(sh)
}

func (bh SmallHeap) Less(i int, j int) bool {
	return bh[i] < bh[j]
}

func (bh SmallHeap) Swap(i int, j int) {
	bh[i], bh[j] = bh[j], bh[i]
}

func (bh *SmallHeap) Push(x any) {
	*bh = append(*bh, x.(int))
}

func (bh *SmallHeap) Pop() any {
	n := len(*bh)
	x := (*bh)[n-1]
	*bh = (*bh)[:n-1]
	return x
}

type MedianFinder struct {
	A *SmallHeap // 小顶堆，存储较大的一半
	B *BigHeap   // 大顶堆，存储较小的一半
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	a := &SmallHeap{}
	heap.Init(a)
	b := &BigHeap{}
	heap.Init(b)
	return MedianFinder{a, b}
}

func (this *MedianFinder) AddNum(num int) {
	if this.A.Len() == this.B.Len() {
		heap.Push(this.B, num)
		bigger := heap.Pop(this.B)
		heap.Push(this.A, bigger)
	}
	if this.A.Len() != this.B.Len() {
		heap.Push(this.A, num)
		smaller := heap.Pop(this.A)
		heap.Push(this.B, smaller)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.A.Len() == 0 && this.B.Len() == 0 {
		return 0
	}
	var res float64
	if this.A.Len() == this.B.Len() {
		res = (float64([]int(*this.A)[0]) + float64([]int(*this.B)[0])) / 2
	} else {
		res = float64([]int(*this.A)[0])
	}
	return res
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
}
