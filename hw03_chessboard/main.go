package main

import "fmt"

func main() {
	razmer := 8

	doska := make([][]string, razmer) //make создает срез поля строчки , при помощи [] [] получается 8х8 где каждый элемент это срез на доске
	for a := 0; a < razmer; a++ {
		doska[a] = make([]string, razmer)
	}

	for a := 0; a < razmer; a++ {
		for b := 0; b < razmer; b++ {
			if (a+b)%2 == 0 {
				doska[a][b] = " "
			} else { //или проьел или решетка согласно условиям задачи
				doska[a][b] = "#"
			}

		}
	}

	for a := 0; a < razmer; a++ { //наполнение клеток проходимся по каждой клетке доски
		for b := 0; b < razmer; b++ {
			fmt.Printf("%s", doska[a][b])
		}
		fmt.Println() //выводим символ новой строки, чтобы создать разделение между строчками доски
	}

}
