package main

import (
	"fmt"
	"math"
)
func main(){
	fmt.Println(mySqrt(8))
	fmt.Println(isPerfectSquare(808201))
	fmt.Print(899*899)
}
func mySqrt(x int) int {

	sqrt := (math.Sqrt(float64(x)))
	return int(math.Floor(sqrt))
}


func isPerfectSquare(num int) bool {
	
    for i:=1; i<(2147483647); i++{
		if i*i == num {
		return true
		}
		if i*i >num{
			break
		}
	}
	return false
}



func twoSum(nums []int, target int) []int {
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++ {
			if (nums[i]+nums[j]== target){
				return []int{i,j}
			}
		}
	}
	return []int{}

}