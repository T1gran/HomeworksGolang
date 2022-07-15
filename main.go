package main

import (
	"strconv" 
)

func ReturnInt()  int {
	a := 1 
	return a
}

func ReturnFloat()  float32 {
	b := 1.1 
	return float32(b)
}

func ReturnIntArray() [3]int  {
	c:=[3]int{1,3,4}
	return c
}

func ReturnIntSlice() []int {
	d:=[]int{1,2,3}
	return d
}

func IntSliceToString(sl []int) string {
	e := strconv.Itoa(sl[0])+strconv.Itoa(sl[1])+strconv.Itoa(sl[2])
	return e
}

func MergeSlices(sl1 []float32, sl2 []int32) []int {
	var sl3 []int
	for _,value := range sl1 {
		sl3=append(sl3, int(value))
	}

	for _,value := range sl2 {
		sl3=append(sl3, int(value))
	}
	return sl3
}

func GetMapValuesSortedByKey(m1 map [int] string) []string {
	a := [] string {}
	b := [] int {}
	for key := range m1 {
		b = append(b, key)
	}
	i := len(b)
	for i>0 {
		i=0
		for index,value := range b {
			if (index < len(m1)-1 && value > b[index+1]) {
				c:= value
				b[index] = b[index+1]
				b[index+1] = c
				i++
			}
		}
	}
	
	for index := range b {
		a = append(a, m1[b[index]])
	}
	return a
}

func main() {
	
}