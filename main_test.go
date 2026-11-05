package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerateRandomElements проверяет генерацию заданного числа положительных чисел
func TestGenerateRandomElements(t *testing.T) {
	arrZero := generateRandomElements(0)
	arrNeg := generateRandomElements(-2)
	arrPos := generateRandomElements(4)

	assert.Len(t, arrZero, 0)
	assert.Len(t, arrNeg, 0)
	assert.Len(t, arrPos, 4)

	for _, num := range arrPos {
 	   assert.Positive(t, num)
	}
}

// TestMaximum проверяет поиск максимума
func TestMaximum(t *testing.T) {
	arrNill := make([]int, 0)
	testNum := 3892
	bigInit := []int{testNum, 1, 2}
	bigEnd := []int{1, 2, testNum}
	bidMid := []int{1, testNum, 2}
	equalNum := []int{testNum, testNum, testNum}
	arrOne := []int{testNum}
	negNum := []int{-2, -4}

	varNill := maximum(arrNill)
	varOne := maximum(arrOne)
	varInit := maximum(bigInit)
	varMid := maximum(bidMid)
	varEnd := maximum(bigEnd)
	varEqual := maximum(equalNum)
	varNeg := maximum(negNum)

	assert.Equal(t, -1, varNill)
	assert.Equal(t, testNum, varOne)
	assert.Equal(t, testNum, varInit)
	assert.Equal(t, testNum, varMid)
	assert.Equal(t, testNum, varEnd)
	assert.Equal(t, testNum, varEqual)
	assert.Equal(t, -1, varNeg)
}

// TestMaxChunks проверяет поиск максимума с использованием горутин
func TestMaxChunks(t *testing.T) {
	arrNill := make([]int, 0)
	testNum := 3892
	negNum := []int{-2, -4}
	bigInit := []int{testNum, 1, 2, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8}
	bigEnd := []int{4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 1, 2, testNum}
	bidMid := []int{4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 1, testNum, 2, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8}
	equalNum := []int{testNum, testNum, testNum}
	arrOne := []int{testNum}
	eightNums := []int{1, 2, 3, 4, 5, 6, 7, testNum}
	sevenNums := []int{1, 2, testNum, 4, 5, 6, 7}
	nineNums := []int{1, testNum, 2, 3, 4, 5, 6, 7, 8}

	varNill := maxChunks(arrNill)
	varOne := maxChunks(arrOne)
	varInit := maxChunks(bigInit)
	varMid := maxChunks(bidMid)
	varEnd := maxChunks(bigEnd)
	varEqual := maxChunks(equalNum)
	varEight := maxChunks(eightNums)
	varSeven := maxChunks(sevenNums)
	varNine := maxChunks(nineNums)
	varNeg := maxChunks(negNum)

	assert.Equal(t, -1, varNill)
	assert.Equal(t, testNum, varOne)
	assert.Equal(t, testNum, varInit)
	assert.Equal(t, testNum, varMid)
	assert.Equal(t, testNum, varEnd)
	assert.Equal(t, testNum, varEqual)
	assert.Equal(t, testNum, varEight)
	assert.Equal(t, testNum, varSeven)
	assert.Equal(t, testNum, varNine)
	assert.Equal(t, -1, varNeg)
}
