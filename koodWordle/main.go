package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"koodWordle/game"
	"koodWordle/io"
)

func seeStats(user string) {
	fmt.Print("Do you want to see your stats? (yes/no): ")

	for {
		reply, ok := io.Scanner()
		if !ok {
			return
		}

		if reply == "yes" {
			played, won, avg := io.ReadCSV(user)

			fmt.Println("Stats for " + user + ":")
			fmt.Printf("Games played: %d\n", played)
			fmt.Printf("Games won: %d\n", won)
			fmt.Printf("Average attempts per game: %.2f\n", avg)
			break

		} else if reply == "no" {
			break
		}
	}

	fmt.Println("Press Enter to exit...")

	for {
		exit, ok := io.Scanner()

		if !ok || exit == "" {
			return
		}
	}
}

func main() {
	// Basics
	remainingLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	attempts := 6
	attemptUsed := 0

	// If no number, bye
	if len(os.Args) != 2 {
		fmt.Println("Please provide a number as command line argument")
		return
	}

	// Take the number of the test word
	wordTest, err := strconv.Atoi(os.Args[1])

	// If what I got is not a number, bye
	if err != nil {
		fmt.Println("Invalid command-line argument. Please launch with a valid number.")
		return
	}

	// Take the words
	words := io.ReadWords()

	if len(words) == 0 {
		return
	}

	// Ask for name and manage ctrl+d
	fmt.Print("Enter your username: ")
	user, ok := io.Scanner()

	if !ok {
		return
	}

	// The number is a number, but the word does not exist
	if wordTest < 0 || wordTest >= len(words) {
		fmt.Println("Invalid word number.")
		fmt.Println("Press Enter to exit...")
		io.Scanner()
		return
	}

	// Take the match
	word := words[wordTest]

	// Greet the MTF
	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")

	// Until the user don't finish the attempts
	for attemptUsed < attempts {
		fmt.Print("Enter your guess:  ")

		guess, ok := io.Scanner()

		if !ok {
			return
		}

		// Checky Checky time

		// Must be 5 letters
		if len([]rune(guess)) != 5 {
			fmt.Println("Your guess must be exactly 5 letters long.")
			continue
		}

		// Only lowercase letters
		valid := true

		for _, r := range guess {
			if r < 'a' || r > 'z' {
				valid = false
				break
			}
		}

		if !valid {
			fmt.Println("Your guess must only contain lowercase letters.")
			continue
		}

		// Check if the word exists in the list
		found := false

		for i := range words {
			if words[i] == guess {
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Word not in list. Please enter a valid word.")
			continue
		}

		// If the word passed all the security checks:
		// Count it as one attempt
		attemptUsed++

		// We play and we want back -->
		feedback, newLetters := game.CheckGuess(
			word,
			guess,
			remainingLetters,
		)

		// Remaining letters
		remainingLetters = newLetters

		// If the word finally match, win!
		if strings.ToLower(guess) == strings.ToLower(word) {
			fmt.Println("Congratulations! You've guessed the word correctly.")

			// We record it as a winner
			io.WriteCSV(user, word, attemptUsed, "win")
			break
		}

		// And the feedback
		fmt.Println("Feedback: " + feedback)

		// Otherwise we print back the remaining letters in a nice way
		fmt.Println("Remaining letters: " +	strings.Join(strings.Split(remainingLetters, ""), " ") + " ")

		// And the remaining attempts
		fmt.Println(
			"Attempts remaining:  " +
				strconv.Itoa(attempts-attemptUsed),
		)

		// If a looser
		if attemptUsed == attempts {
			// We say what was the word
			fmt.Println("Game over. The correct word was: " + word)

			// And we record it as a looser
			io.WriteCSV(user, word, attemptUsed, "loss")
			break
		}
	}

	// At the end we ask if the user wants to see the stats
	seeStats(user)
}