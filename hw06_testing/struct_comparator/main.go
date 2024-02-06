package comparator

const (
	year FieldComapre = iota
	size
	rate
)

type FieldComapre uint8

type Comparator struct {
	Type FieldComapre
}

func NewComparator(t FieldComapre) *Comparator {
	return &Comparator{
		Type: t,
	}
}

func (c Comparator) Compare(bookOne, bookTwo *Book) bool {
	switch c.Type {
	case year:
		return bookOne.year > bookTwo.year
	case size:
		return bookOne.size > bookTwo.size
	case rate:
		return bookOne.rate > bookTwo.rate
	default:
		return false
	}
}

type Book struct {
	id     uint
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (b *Book) SetRate(newrate float32) {
	b.rate = newrate
}

func (b *Book) SetSize(newsize int) {
	b.size = newsize
}

func (b *Book) SetYear(newyear int) {
	b.year = newyear
}

func (b *Book) SetTitle(newtitle string) {
	b.title = newtitle
}

func (b *Book) SetAuthor(newauthor string) {
	b.author = newauthor
}

func (b *Book) SetID(newid uint) {
	b.id = newid
}

func (b *Book) Rate() float32 {
	return b.rate
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) ID() uint {
	return b.id
}

func (s FieldComapre) String() string {
	switch s {
	case size:
		return "size"
	case rate:
		return "rate"
	case year:
		return "year"
	}
	return "unknown"
}
