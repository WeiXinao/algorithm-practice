package main

import (
	"fmt"
)

type Checkout struct {
	priceQueue []int
	sortQueue  []int
}

func Constructor() Checkout {
	var checkout Checkout
	checkout.priceQueue = make([]int, 0)
	checkout.sortQueue = make([]int, 0)
	return checkout
}

func (this *Checkout) Get_max() int {
	if len(this.sortQueue) <= 0 {
		return -1
	}
	return this.sortQueue[0]
}

func (this *Checkout) Add(value int) {
	this.priceQueue = append(this.priceQueue, value)
	i := 0
	if i == 0 || value > this.sortQueue[i-1] {
		for i = 0; i < len(this.sortQueue) && this.sortQueue[i] >= value; i++ {
		}
		this.sortQueue = this.sortQueue[:i]
	}
	this.sortQueue = append(this.sortQueue, value)
	i++
}

func (this *Checkout) Remove() int {
	if len(this.priceQueue) <= 0 {
		return -1
	}
	rmPrice := this.priceQueue[0]
	this.priceQueue = this.priceQueue[1:]
	if rmPrice == this.sortQueue[0] {
		this.sortQueue = this.sortQueue[1:]
	}
	return rmPrice
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */

func main() {
	obj := Constructor()
	// obj.Add(4)
	// obj.Add(7)
	// param_1 := obj.Get_max()
	// param_2 := obj.Remove()
	// param_3 := obj.Get_max()
	// fmt.Println(param_1, param_2, param_3)

	param_1 := obj.Remove()
	param_2 := obj.Get_max()
	fmt.Println(param_1, param_2)
}
