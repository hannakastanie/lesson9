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
	testNum := 3892

	assert.Equal(t, 0, maximum([]int{}))
	assert.Equal(t, testNum, maximum([]int{testNum}))
	assert.Equal(t, testNum, maximum([]int{testNum, 1, 2}))
	assert.Equal(t, testNum, maximum([]int{1, testNum, 2}))
	assert.Equal(t, testNum, maximum([]int{1, 2, testNum}))
	assert.Equal(t, testNum, maximum([]int{testNum, testNum, testNum}))
	assert.Equal(t, -2, maximum([]int{-2, -4}))
}

// TestMaxChunks проверяет поиск максимума с использованием горутин
func TestMaxChunks(t *testing.T) {
	testNum := 3892

	assert.Equal(t, 0, maxChunks([]int{}))
	assert.Equal(t, testNum, maxChunks([]int{testNum}))
	assert.Equal(t, testNum, maxChunks([]int{testNum, 1, 2, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8}))
	assert.Equal(t, testNum, maxChunks([]int{4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 1, 2, testNum}))
	assert.Equal(t, testNum, maxChunks([]int{4, 5, 6, 7, 8, 4, 5, 6, testNum, 5, 6, 7, 8, 4, 5, 6, 7, 8}))
	assert.Equal(t, testNum, maxChunks([]int{testNum, testNum, testNum}))
	assert.Equal(t, testNum, maxChunks([]int{1, 2, 3, 4, 5, 6, 7, testNum}))
	assert.Equal(t, testNum, maxChunks([]int{1, 2, testNum, 4, 5, 6, 7}))
	assert.Equal(t, testNum, maxChunks([]int{1, testNum, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, -2, maxChunks([]int{-2, -4}))
}
