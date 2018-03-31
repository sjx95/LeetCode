package main

import "fmt"

func main() {
	a := Constructor([]int{1, 2, 4, 3, 4, 5, 6})
	fmt.Println(a.SumRange(2, 2))
	a.Update(1, 2)
	fmt.Println(a.SumRange(0, 4))
}

type NumArray struct {
	lch, rch *NumArray
	l, r int
	sum int
}


func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		var e NumArray
		e.l, e.r = -1, -1
		return e
	}
	var layer []NumArray
	for i, n := range nums {
		var e NumArray
		e.l, e.r = i, i
		e.sum = n
		layer = append(layer, e)
	}
	for len(layer) > 1 {
		i := 0
		var upper_layer []NumArray
		for i < len(layer) {
			var e NumArray
			e.lch = &layer[i]
			e.l, e.r = e.lch.l, e.lch.r
			e.sum = e.lch.sum
			i++
			if i < len(layer) {
				e.rch = &layer[i]
				e.r = e.rch.r
				e.sum += e.rch.sum
				i++
			}
			upper_layer = append(upper_layer, e)
		}
		layer = upper_layer
	}
	return layer[0]
}


func (this *NumArray) Update(i int, val int)  {
	if this.l == i && this.r == i {
		this.sum = val
	} else {
		if i <= this.lch.r {
			this.lch.Update(i, val)
		} else {
			this.rch.Update(i, val)
		}
		if this.rch == nil {
			this.sum = this.lch.sum
		} else {
			this.sum = this.lch.sum + this.rch.sum
		}
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i <= this.l && j >= this.r {
		return this.sum
	}
	if j < this.l || i > this.r {
		return 0
	}
	if this.rch != nil {
		return this.lch.SumRange(i, j) + this.rch.SumRange(i, j)
	} else {
		return this.lch.SumRange(i, j)
	}
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
