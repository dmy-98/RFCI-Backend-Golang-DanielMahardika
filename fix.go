package main

import (
	"fmt"
)

func main() {
	collection := []int{3, 5, 2, -4, 8, 11}
	fmt.Println(foo(7, collection))
}

func foo(input int, numbers []int)  [][]int {
	var result [][]int
	
	for i:=len(numbers)-1; i>0;i--{
		for j:=i-1;j>=0;j--{
			var temp []int
			if numbers[i]+numbers[j]==input{
				temp=append(temp,numbers[i])
				temp=append(temp,numbers[j])
				result=append(result,temp)
			}
		}
	}

	return result
}
