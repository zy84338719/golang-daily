package main

import "fmt"

func validMountainArray(A []int) bool {
	if len(A)<3 {
		return false
	}
	i:=1
	for ;i<len(A);i++ {
		if A[i]==A[i]{
			return false
		}
		if A[i-1]<A[i]{
			break
		}
	}
	for ;i<len(A);i++{
		if  A[i-1]-A[i]<=0{
			return false
		}
	}
	return true
}

func main() {
	var A []int = []int{3,5,5}
	fmt.Println(validMountainArray(A))
}