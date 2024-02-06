package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	sortedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Тестовые случаи для успешного поиска
	index, err := binarySearch(0, len(sortedSlice)-1, 3, sortedSlice)
	assert.NoError(t, err, "Ожидался успех при поиске элемента")
	assert.Equal(t, 2, index, "Неправильный индекс найденного элемента")

	index, err = binarySearch(0, len(sortedSlice)-1, 7, sortedSlice)
	assert.NoError(t, err, "Ожидался успех при поиске элемента")
	assert.Equal(t, 6, index, "Неправильный индекс найденного элемента")

	// Тестовый случай для неудачного поиска
	index, err = binarySearch(0, len(sortedSlice)-1, 10, sortedSlice)
	assert.Error(t, err, "Ожидалась ошибка при неудачном поиске")
	assert.Equal(t, 0, index, "Неправильный индекс при неудачном поиске")
}
