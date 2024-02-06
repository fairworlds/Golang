package chessboard

import "errors"

const (
	black = "#"
	white = " "
)

func CreateString(countColumn int, firstStep int) string {
	var readyString string
	for i := 0; i < countColumn; i++ {
		if len(readyString)%2 == 0 && firstStep%2 == 0 {
			readyString += black
			continue
		} else if len(readyString)%2 != 0 && firstStep%2 == 0 {
			readyString += white
			continue
		}
		if len(readyString)%2 == 0 && firstStep%2 != 0 {
			readyString += white
			continue
		} else if len(readyString)%2 != 0 && firstStep%2 != 0 {
			readyString += black
			continue
		}
	}
	return readyString
}

func ChessBoard(row, column int) (string, error) {
	var board string
	if column <= 0 || row <= 0 {
		return "", errors.New(
			"значение не может быть меньше или равно нулю")
	}
	for i := 0; row > i; i++ {
		board += "\n" + CreateString(column, i)
	}
	return board, nil
}
