package main

import "fmt"

func before() {

	a := []int{1, 2, 3}
	b := []int64{1, 2, 3}
	c := []float32{1, 2, 3}

	fmt.Println(addint(a))
	fmt.Println(addint64(b))
	fmt.Println(addfloat32(c))

}

func addint(in []int) int {
	var sum int
	for _, v := range in {
		sum += v
	}
	return sum
}

func addint64(in []int64) int64 {
	var sum int64
	for _, v := range in {
		sum += v
	}
	return sum
}

func addfloat32(in []float32) float32 {
	var sum float32
	for _, v := range in {
		sum += v
	}
	return sum
}
