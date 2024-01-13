package main

import (
	"fmt"
)

type CompareMode int

const ( // задаем константу для режима сравнения
	Year CompareMode = iota
	Size
	Rate
)

type Book struct { // определяем структуру
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

// Методы для установки и получения полей структуры

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) GetRate() float64 {
	return b.rate
}

// Структура с методом для сравнения книг

type BookComparator struct {
	mode CompareMode // константа режима сравнения
}

func NewBookComparator(mode CompareMode) *BookComparator {
	return &BookComparator{
		mode: mode,
	}
}

func (bc *BookComparator) Compare(book1, book2 Book) bool {
	switch bc.mode {
	case Year:
		return book1.year > book2.year
	case Size:
		return book1.size > book2.size
	case Rate:
		return book1.rate > book2.rate
	default:
		return false
	}
}

func main() {
	book1 := Book{
		id:     1,
		title:  "Book 1",
		author: "Author 1",
		year:   2022,
		size:   200,
		rate:   4.5,
	}

	book2 := Book{
		id:     2,
		title:  "Book 2",
		author: "Author 2",
		year:   2020,
		size:   150,
		rate:   4.2,
	}

	comparator := NewBookComparator(Year)
	fmt.Println(comparator.Compare(book1, book2))
	// Выводит true, так как год у book1 больше года у book2
}
