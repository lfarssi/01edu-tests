package main

import (
    "fmt"
    "math/rand"
    "time"
)

func generateUniqueRandomNumbers(n int) []int {
    // Create a map to store unique numbers
    arr := make(map[int]bool)
    rand.Seed(time.Now().UnixNano()) // Seed random number generator

    // Create a slice to hold the final result
    var result []int

    // Keep generating random numbers until we have 'n' unique numbers
    for len(arr) < n {
        num := rand.Int() // Generate a random number
        if !arr[num] {    // Check if the number is unique
            arr[num] = true              // Add to the map
            result = append(result, num) // Add to the result slice
        }
    }

    return result
}

func main() {
    // Generate 500 unique random numbers
    numbers := generateUniqueRandomNumbers(100)

    // Print the result
    for _, num := range numbers {
        fmt.Print(num, " ")
    }
}