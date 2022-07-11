package main

func a(arr []int32) int32 {
	var max int32
    var count int32
    for i := 0; i < len(arr); i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    for i := 0; i < len(arr); i++ {
        if arr[i] != max {
            count++
        }
    }
    return count
}