package main

import (
    "fmt"
)

func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}

func main() {
    var num int
    var nums []int

    fmt.Println("Enter a number to find:")
    fmt.Scanln(&num)

    for {
        var n int
        fmt.Println("Enter a number to add (or enter 'q' to quit):")
        _, err := fmt.Scanln(&n)
        if err != nil {
            break
        }
        nums = append(nums, n)
    }

    find(num, nums...)
}
