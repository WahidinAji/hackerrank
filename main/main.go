package main

import "fmt"

func main() {
	// fmt.Println(angryProfessor(3, []int32{-1, -3, 4, 2}))
	// fmt.Println(angryProfessor(2, []int32{0, -1, 2, 1}))

	// //beatiful Days
	// fmt.Println(beautifulDays(22, 2, 2))
	// fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
	fmt.Println(equalizeArray([]int32{4, 4, 2, 1, 4, 4}))
}

// dup := func(arr []int32) int32 {
// 	var temp = make(map[int32]bool, 0)
// 	for i := 0; i < len(arr); i++ {
// 		if temp[arr[i]] == true {
// 			return arr[i]
// 		} else {
// 			temp[arr[i]] = true
// 		}
// 	}
// 	return -1
// }
// fmt.Println(dup(arr))
//https://www.hackerrank.com/challenges/equality-in-a-array/problem?isFullScreen=true
func equalizeArray(arr []int32) int32 {
	// var count int32
	var min []int32
	var max int32
	var temp = make(map[int32]bool, 0)
	for i := 0; i < len(arr); i++ {
		if temp[arr[i]] == true {
			continue
		} else {
			temp[arr[i]] = true
		}
	}
	fmt.Println(min)
	max = int32(len(temp))
	return max

	// dup := func(arr []int32) int32 {
	// 	var temp = make(map[int32]bool, 0)
	// 	for i := 0; i < len(arr); i++ {
	// 		if temp[arr[i]] == true {
	// 			return arr[i]
	// 		} else {
	// 			temp[arr[i]] = true
	// 		}
	// 	}
	// 	return -1
	// }
	// fmt.Println(dup(arr))
	// return 1
}

//https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var count int32
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}

func angryProfessor(k int32, a []int32) string {
	var onTime []int32

	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			onTime = append(onTime, a[i])
		}
	}
	if len(onTime) < int(k) {
		return "YES"
	}
	return "NO"
}

func beautifulDays(i int32, j int32, k int32) int32 {
	a := fmt.Sprint(i)
	fmt.Println(len(a))
	return 1
}
