package main

import (
	"fmt"
	// "github.com/xudipta/prep/problems"
	"github.com/xudipta/prep/utils"
)

func main() {
	fmt.Println("Welcome to the Prep Playground!")
	
	// Example usage of utils.Sum
	nums := []int{1, 2, 3, 4, 5}
	sum := utils.Sum(nums)
	fmt.Printf("The sum of %v is %d\n", nums, sum)
}
