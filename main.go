// это запись разбора домашнего задания модуля 20 от лектора
package main

import "fmt"

const rows = 3
const cols = 3

const oneRows = 3
const oneCols = 5
const twoCols = 4

func determinant(matrix [rows][cols]int64) int64 {
	minorDetOne := matrix[1][1]*matrix[2][2] - matrix[1][2]*matrix[2][1]
	minorDetTwo := matrix[1][0]*matrix[2][2] - matrix[1][2]*matrix[2][0]
	minorDetThree := matrix[1][0]*matrix[2][1] - matrix[1][1]*matrix[2][0]

	determinant := matrix[0][0]*minorDetOne - matrix[0][1]*minorDetTwo + matrix[0][2]*minorDetThree

	return determinant
}

func multiply(m1 [oneRows][oneCols]int64, m2 [oneCols][twoCols]int64) [oneRows][twoCols]int64 {
	var m [oneRows][twoCols]int64
	for i := 0; i < oneRows; i++ {
		for j := 0; j < twoCols; j++ {
			for k := 0; k < oneCols; k++ {
				m[i][j] = m[i][j] + m1[i][k]*m2[k][j]
			}
		}
	}
	return m
}

func main() {
	mm := [rows][cols]int64{
		{1, 0, 10},
		{10, 11, 12},
		{10, 11, 13},
	}
	fmt.Println(determinant(mm))
	m1 := [oneRows][oneCols]int64{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}
	m2 := [oneCols][twoCols]int64{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	fmt.Println(multiply(m1, m2))
}
