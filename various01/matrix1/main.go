package main

import (
    "fmt"
    //"bufio"
    //"os"
)

func main() {

matrix1 := []int{1,2,3}
matrix2 := [][]int{{1,10,100},{2,20,200}}

fmt.Printf ("Matrix1: %v\n", matrix1)
fmt.Printf ("Matrix2: %v\n", matrix2)

fmt.Printf ("Matrix1 length: %d\n", len(matrix1))
fmt.Printf ("Matrix2 length: %d\n", len(matrix2))

fmt.Printf ("Matrix1 element 0: %d\n", matrix1[0])
fmt.Printf ("Matrix2 element 0: %d\n", matrix2[0])
fmt.Printf ("Matrix2 element 0,0: %d\n", matrix2[0][0])


for i := range matrix2 {
    fmt.Printf ("Matrix2 row %d: %v \n", i, matrix2[i])

    for j := range matrix2[i] {
        fmt.Printf ("Matrix1 row %d, column %d: %v \n", i, j, matrix2[i][j])
    }
}


matrixa := [2][2]int{{1,1},{0,1}}

for i := range matrixa {
    fmt.Printf ("Matrixa row %d: %v \n", i, matrixa[i])

    for j := range matrixa[i] {
        fmt.Printf ("Matrixa row %d, column %d: %v \n", i, j, matrixa[i][j])
    }
}

var adj := [][]int{}

}

