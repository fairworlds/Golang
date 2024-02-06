package main

import (
	"errors"
	"math"
	"testing"
)

func TestCalculateArea(t *testing.T) {
	// Тестовые случаи
	testCases := []struct {
		shape Shape
		area  float64
		err   error
	}{
		{Circle{Radius: 5}, math.Pi * 25, nil},
		{Rectangle{Width: 10, Height: 5}, 50, nil},
		{Triangle{Base: 8, Height: 6}, 24, nil},
		{nil, 0, errors.New("переданный объект не реализует интерфейс Shape")},
	}

	// Проверяем каждый тестовый случай
	for _, tc := range testCases {
		area, err := calculateArea(tc.shape)

		// Проверяем площадь
		if area != tc.area {
			t.Errorf("Неправильная площадь. Ожидалось %.2f, получено %.2f", tc.area, area)
		}

		// Проверяем ошибку
		if (err == nil && tc.err != nil) || (err != nil && tc.err == nil) || (err != nil && tc.err != nil && err.Error() != tc.err.Error()) {
			t.Errorf("Неправильная ошибка. Ожидалось '%v', получено '%v'", tc.err, err)
		}
	}
}
