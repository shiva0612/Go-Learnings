package main

import "fmt"

func after() {

	a := []int{1, 2, 3}
	b := []int64{1, 2, 3}
	c := []float32{1, 2, 3}

	aa := addnumber2(a) //aa is of type int
	bb := addnumber2(b) //bb is of type int64
	cc := addnumber2(c) //cc is of type float32
	fmt.Println(aa, bb, cc)

}

// without generics
func addnumber1(in []any) any {
	var sum float32
	for _, v := range in {
		sum += v.(float32) //here u will have to always typecase and add - thats why use generics
	}
	return sum
}

func addnumber2[T int | int64 | float32](in []T) T {
	var sum T
	for _, v := range in {
		sum += v
	}
	return sum
}

type Number interface {
	int | int64 | float32
}

func addnumber3[T Number](in []T) T {
	var sum T
	for _, v := range in {
		sum += v
	}
	return sum
}

func multipleGenerics[Number int32 | int64, Decimal float32 | float64](number Number, decimal Decimal) (Number, Decimal) {
	return number + 1, decimal + 1
}
