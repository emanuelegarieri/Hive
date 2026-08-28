package io

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func Scanner() (string, bool) {
	// We manage "gracefully the ctrl+d"
	if !scanner.Scan() {
		return "", false
	}
	input := strings.TrimSpace(scanner.Text())
	return input, true
}

func ReadWords() []string {
	var words []string
	// Open file
	file, err := os.Open("wordle-words.txt")
	if err != nil {
		return nil
	}
	defer file.Close()
	// Scan
	fileScanner := bufio.NewScanner(file)
	// Create a list of words
	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}
	return words
}

func WriteCSV(user string, word string, attempts int, result string) {
	file, err := os.OpenFile("stats.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil {
		return
	}
	defer file.Close()
	// Writer
	writer := csv.NewWriter(file)
	// What to write
	writer.Write([]string{user,	word, strconv.Itoa(attempts), result})
	// Clear cache
	writer.Flush()
}

func ReadCSV(user string) (int, int, float32) {
	var played int
	var won int
	var totalAttempts int

	file, err := os.Open("stats.csv")
	if err != nil {
		return 0, 0, 0
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, 0, 0
	}

	for i := range records {
		if records[i][0] == user {
			played++
			attempts, err := strconv.Atoi(records[i][2])
			// Always gracefully
			if err != nil {
				continue
			}
			totalAttempts += attempts
			if records[i][3] == "win" {
				won++
			}
		}
	}

	// Average score
	if played == 0 {
		return 0, 0, 0
	}
	avg := float32(totalAttempts) / float32(played)
	return played, won, avg
}