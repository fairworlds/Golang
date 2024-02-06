package shapes

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface { // Создаем интерфейс Shape с единственным.
	// методом для вычисления площади.
	Area() float64
}

type Circle struct { // Структура для представления круга.
	Radius float64
}

func (c Circle) Area() float64 { // Метод для вычисления площади круга.
	// (реализация интерфейса Shape).
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct { // Структура для представления прямоугольника
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 { // Метод для вычисления площади
	// прямоугольника (реализация интерфейса Shape)
	return r.Width * r.Height
}

type Triangle struct { // Структура для представления треугольника
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	// Метод для вычисления площади треугольника
	// (реализация интерфейса Shape)
	return 0.5 * t.Base * t.Height
}

func calculateArea(s any) (float64, error) {
	// Функция calculateArea ожидает на входе объект
	// типа Shape и возвращает его площадь.
	// Если переданный объект не реализует интерфейс
	// Shape, функция возвращает ошибку
	if shape, ok := s.(Shape); ok {
		return shape.Area(), nil
	}
	return 0, errors.New("переданный объект не реализует интерфейс Shape")
}

func Lol() {
	// Создаем объекты разных типов
	// (круг, прямоугольник, треугольник)
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}
	triangle := Triangle{Base: 8, Height: 6}

	// Вызываем функцию calculateArea для каждого
	// объекта и обрабатываем возможные ошибки
	area1, err1 := calculateArea(circle)
	area2, err2 := calculateArea(rectangle)
	area3, err3 := calculateArea(triangle)

	// Выводим результаты
	if err1 == nil {
		fmt.Printf("Круг: радиус %.2f\n", circle.Radius)
		fmt.Printf("Площадь: %.2f\n", area1)
	} else {
		fmt.Println("Ошибка:", err1)
	}

	if err2 == nil {
		fmt.Printf("Прямоугольник: ширина %.2f, высота %.2f\n", rectangle.Width, rectangle.Height)
		fmt.Printf("Площадь: %.2f\n", area2)
	} else {
		fmt.Println("Ошибка:", err2)
	}

	if err3 == nil {
		fmt.Printf("Треугольник: основание %.2f, высота %.2f\n", triangle.Base, triangle.Height)
		fmt.Printf("Площадь: %.2f\n", area3)
	} else {
		fmt.Println("Ошибка:", err3)
	}
}
