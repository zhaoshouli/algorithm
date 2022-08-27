package main

import "fmt"

type MovingAverage struct {
	index   int   // 当前环形数组的位置
	count   int   // 数组大小
	sum     int   // 数据总量
	buckets []int // 环形数组
}

func Constructor(size int) MovingAverage {
	return MovingAverage{index: size, buckets: make([]int, size)}
}

func (mv *MovingAverage)Next(val int) float64 {
	mv.sum += val
	mv.index = (mv.index + 1) % len(mv.buckets)
	fmt.Println("--", mv.index)
	if mv.count < len(mv.buckets) {
		mv.count ++
		mv.buckets[mv.index] = val
	} else {
		mv.sum -= mv.buckets[mv.index]
		mv.buckets[mv.index] = val
	}
	return float64(mv.sum) / float64(mv.count)
}

func main()  {
	con := Constructor(3)
	fmt.Println(con.Next(1))
	fmt.Println(con.Next(10))
	fmt.Println(con.Next(3))
	fmt.Println(con.Next(1))
	fmt.Println(con.Next(10))
}
