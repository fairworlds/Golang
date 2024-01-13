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
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float64
}

// Методы для установки и получения полей структуры

func (b *Book) SetID(id int) {
	b.ID = id
}

func (b *Book) GetID() int {
	return b.ID
}

func (b *Book) SetTitle(title string) {
	b.Title = title
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) SetAuthor(author string) {
	b.Author = author
}

func (b *Book) GetAuthor() string {
	return b.Author
}

func (b *Book) SetYear(year int) {
	b.Year = year
}

func (b *Book) GetYear() int {
	return b.Year
}

func (b *Book) SetSize(size int) {
	b.Size = size
}

func (b *Book) GetSize() int {
	return b.Size
}

func (b *Book) SetRate(rate float64) {
	b.Rate = rate
}

func (b *Book) GetRate() float64 {
	return b.Rate
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
		return book1.Year > book2.Year
	case Size:
		return book1.Size > book2.Size
	case Rate:
		return book1.Rate > book2.Rate
	default:
		return false
	}
}

func main() {
	book1 := Book{
		ID:     1,
		Title:  "Book 1",
		Author: "Author 1",
		Year:   2022,
		Size:   200,
		Rate:   4.5,
	}

	book2 := Book{
		ID:     2,
		Title:  "Book 2",
		Author: "Author 2",
		Year:   2020,
		Size:   150,
		Rate:   4.2,
	}

	comparator := NewBookComparator(Year)
	fmt.Println(comparator.Compare(book1, book2))
	// Выводит true, так как год у book1 больше года у book2
}
