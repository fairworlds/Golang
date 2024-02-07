package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func countWorld(row string) map[string]int { // прием строки в качестве аргумента
	wordDict := make(map[string]int)
	words := strings.Fields(row)

	for _, text := range words {
		_, ok := wordDict[strings.ToLower(text)]
		if !ok {
			wordDict[strings.ToLower(text)] = 1
		} else {
			wordDict[strings.ToLower(text)]++
		}
	}
	return wordDict
}

func WordSplit(row string) ([]string, error) {
	text := []string{}

	if len(strings.TrimSpace(row)) == 0 {
		return []string{""}, errors.New("передали пустую строку")
	}
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_!.' ''\n'", r)
	}
	words := strings.FieldsFunc(row, splitFunc)
	text = append(text, words...)
	return text, nil
}

func main() {
	fmt.Println("Введите строку:")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	splitWords, err := WordSplit(text)
	if err != nil {
		panic(err)
	}
	fmt.Println("Разделенные слова:", splitWords)
	wordCount := countWorld(text)
	fmt.Println("Количество повторений слов:", wordCount)
}
