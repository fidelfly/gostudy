package basic

import "fmt"

type SliceContainer struct {
	MySlice       []int
	MySliceObject []IntObject
}

type IntObject struct {
	value int
}

func RunSliceTest() {

	c := SliceContainer{}

	intSlice := []int{1, 2, 3, 4, 5}

	c.MySlice = intSlice

	var intSlice0 []int

	intSlice0 = intSlice

	intSlice[2] = 100

	fmt.Println("Int slice :", c.MySlice[2], " and ", intSlice0[2])

	objSlice := []IntObject{IntObject{1}, IntObject{2}, IntObject{3}}

	c.MySliceObject = objSlice

	objSlice[1].value = 100

	fmt.Println("Object slice : ", c.MySliceObject[1].value)

	intArray := [...]int{1, 2, 3}

	var intArray0 [3]int
	intArray0 = intArray

	intArray[1] = 100

	fmt.Println("Array : ", intArray0[1])

}
