package main

import (
	"bufio" // пакет для чтения ввода с клавиатуры
	"fmt"
	"os"      // для работы операционной системой
	"strconv" // преобразрование из строки в число
)

func main() {
	reader := bufio.NewReader(os.Stdin) // считываем с клавиатуры и сохраняем в sizereader

	fmt.Print("Введите размер доски: ")
	sizereader, _ := reader.ReadString('\n')
	sizereader = sizereader[:len(sizereader)-1] // удаляем символ новой строки Len возвращает длину
	// строки или кода len(sizereader) возвращает длину переменной sizereader

	sizechess, err := strconv.Atoi(sizereader) // преобразуем данные в целое число при этом
	// atoi - это функция в пакете strconv, которая преобразует строку в целое число (integer).
	// В данном случае strconv.Atoi(sizereader) преобразует значение переменной sizereader из строки в целое число (integer).
	// Если преобразование не удалось, функция возвращает ошибку
	if err != nil {
		fmt.Println("Ошибка при чтении размера доски:", err)
		return
	}

	chess := make([][]string, sizechess)
	for a := 0; a < sizechess; a++ {
		chess[a] = make([]string, sizechess)
	}

	for a := 0; a < sizechess; a++ {
		for b := 0; b < sizechess; b++ {
			if (a+b)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
