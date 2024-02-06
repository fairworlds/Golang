package comparator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		book1   Book
		book2   Book
		compare FieldComapre
		want    bool
	}{
		{"Good", Book{1, "War and peace", "Tolstoi", 1879, 1000, 8}, Book{2, "Python", "Lics", 2023, 300, 5}, 2, true},
		{"Good1", Book{1, "War and peace", "Tolstoi", 1879, 1000, 8}, Book{2, "Python", "Lics", 2023, 300, 5}, 1, true},
		{"Good2", Book{1, "War and peace", "Tolstoi", 1879, 1000, 8}, Book{2, "Python", "Lics", 2023, 300, 5}, 3, false},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := NewComparator(tC.compare)
			result := got.Compare
			assert.Equal(t, result(&tC.book1, &tC.book2), tC.want)
		})
	}
}
