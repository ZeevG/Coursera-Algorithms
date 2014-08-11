package main

import (
    "fmt"
)

func merge(left, right []int) []int {

    // Make a new slice with a length of 0 and
    // a capacity of len(right)+len(left)
    sorted := make([]int, 0, len(right)+len(left))

    for len(left) != 0 && len(right) != 0 {
        if left[0] < right[0] {
            // Is there a better way to remove the first element?
            sorted = append(sorted, left[0])
            left = left[1:]
        } else {
            sorted = append(sorted, right[0])
            right = right[1:]
        }
    }

    for len(left) != 0 {
        sorted = append(sorted, left[0])
        left = left[1:]
    }

    for len(right) != 0 {
        sorted = append(sorted, right[0])
        right = right[1:]
    }


    return sorted
}

func mergeSort(unsorted []int) []int {

    if len(unsorted) == 1 {
        return unsorted
    }


    middle := len(unsorted)/2.0
    left := unsorted[:middle]
    right := unsorted[middle:]

    right = mergeSort(right)
    left = mergeSort(left)

    return merge(right, left)

}


func main() {

    unsorted := []int {
        34,
        23,
        423,
        356,
        43,
        4,
        2,
        32,
        41,
        2154,
        5,
        63,
        43,
        123,
        42,
        34,
        52,
        35,
    }

    fmt.Printf("%v\n",unsorted)

    sorted := mergeSort(unsorted)

    fmt.Printf("==========\n%v\n", sorted)
    
}