package main


import "fmt"


// takes a slice of integers and returns the sum of all the elements in the slice.
func sumSlice(arr []int) int {
	sum := 0
	// for index, value := range collection
	for _, n := range arr {
		sum += n
	}
	return sum
}

// takes a slice of integers and returns the minimum value in the slice
func minValue(arr []int) int {
	min := arr[0]
	for _, n := range arr {
		if n < min {
			min = n
		}
	}
	return min
}

// takes a slice of integers and returns a new slice with all the even numbers from the original slice
func minValueFixed(nums []int) int {
	if (len(nums) == 0) {
		panic("empty slice")
	}
	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}
	return min
}

func evenNumber(nums []int) []int {
	evenNum := []int{}

	for  _, n := range nums {
		if n%2 == 0 {
			evenNum = append(evenNum, n)
		}
	}
	return evenNum
}

// func main(){
// 	arr := []int{1,2,3,-1,-4,4}
// 	//var number_of_element int
// 	//fmt.Print("Enter the number of elements: ")
//     //fmt.Scan(&number_of_element)
// 	//fmt.Println("You entered: ", number_of_element)

// 	//// var sum int
//     //sum:=0
// 	//for i:=0; i<number_of_element; i++ {
// 	//	fmt.Printf("enter the %vth element: ", i)
// 	//	var temp int
// 	//	fmt.Scan(&temp)
// 	//	arr = append(arr, temp)
// 	//	sum += temp
// 	//}
// 	//fmt.Println("Sum of array is: ",sum)
// 	sum := 0
// 	sum = sumSlice(arr);
// 	fmt.Println("sum of slice is:", sum)
// 	var min = minValue(arr)
// 	fmt.Println("min of slice is:", min)

// 	min2 := minValueFixed(arr)
// 	fmt.Println("min of slice fixed is:", min2)

// 	evenNum := evenNumber(arr)
// 	fmt.Print("list even number: ")
// 	for _, n := range evenNum{
// 		fmt.Print(n, " ")
// 	}
// }
