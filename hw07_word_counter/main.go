package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func countWorld(row []string) map[string]int {
	wordDict := make(map[string]int)
	for _, text := range row {
		wordDict[strings.ToLower(text)]++
	}
	return wordDict
}

func WordSplit(row string) ([]string, error) {
	text := []string{}

	if len(strings.TrimSpace(row)) == 0 {
		return nil, errors.New("empty string")
	}

	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_!.' ''\n'", r)
	}

	words := strings.FieldsFunc(row, splitFunc)
	text = append(text, words...)
	return text, nil
}

func main() {
	fmt.Println("Введите текст")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	splitworld, err := WordSplit(text)
	if err != nil {
		panic(err)
	}
	fmt.Println(countWorld(splitworld))
}
