package main

import "fmt"

func generatePasal(numRows int) [][]int {
    triangle := make([][]int, numRows)

    for i := 0; i < numRows; i++ {
        temp := make([]int, i+1)
        temp[0] = 1
        temp[i] = 1

        for j := 1; j < i; j++ {
            temp[j] = triangle[i-1][j-1] + triangle[i-1][j]
        }

        triangle[i] = temp
    }

    return triangle
}

func main() {
    fmt.Println(generatePasal(5))
    fmt.Println(generatePasal(7))
}
