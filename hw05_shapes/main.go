package main

import (
	"errors"
	"fmt"
	"math"
)

// Создаем интерфейс Shape с единственным.
// методом для вычисления площади.
type Shape interface {
	Area() float64
}

// Структура для представления круга.
type Circle struct {
	Radius float64
}

// Метод для вычисления площади круга.
// (реализация интерфейса Shape).
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Структура для представления прямоугольника
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод для вычисления площади прямоугольника
// (реализация интерфейса Shape)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Структура для представления треугольника
type Triangle struct {
	Base   float64
	Height float64
}

// Метод для вычисления площади треугольника
// (реализация интерфейса Shape)
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Функция calculateArea ожидает на входе объект
// типа Shape и возвращает его площадь.
// Если переданный объект не реализует интерфейс
// Shape, функция возвращает ошибку.
func calculateArea(s Shape) (float64, error) {
	if area := s.Area(); !math.IsNaN(area) {
		return area, nil
	}
	return 0, errors.New("неверный тип объекта")
}

func main() {
	//Создаем объекты разных типов
	//(круг, прямоугольник, треугольник)
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}
	triangle := Triangle{Base: 8, Height: 6}

	//Вызываем функцию calculateArea для каждого
	//объекта и обрабатываем возможные ошибки
	area1, err1 := calculateArea(circle)
	area2, err2 := calculateArea(rectangle)
	area3, err3 := calculateArea(triangle)

	//Выводим результаты на экран
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
