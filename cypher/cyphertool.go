package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main () {
	CypherTool()
}

// Main function which calls the other functions and handle the prints
func CypherTool() { 

	// Print the greeting message
	fmt.Println("\n\nWelcome to the encryption tool!\n---")

	// Store all the input in 3 different variables
	toEncrypt, encoding, text := getInput()

	result := ""

	// The if check which encryption types and text the user selects, and send the inputs to the corresponding functions
	// Example: if it has been selected "1", then the third input which is the message
	if encoding == "1" {
		result = rot13(text)

	} else if encoding == "2" {
		result = reverseAlphabet(text)

	} else if encoding == "3" {
		result = EncryptChinese(text, toEncrypt)
	}

	// Print the final result
	fmt.Println("\nHere is your result:")
	fmt.Println(result)
	fmt.Println()
}

// Function for applying Rot13 encryption, suitable for both encryption and decryption
func rot13(text string) string {

	result := ""

	// For each rune in the text
	for _, r := range text {
		// In case of lowercase letters, is added 13 positions until 'z'
		if r >= 'a' && r <= 'z' {
			r = r + 13
			// If we go past z, go back to the beginning
			if r > 'z' {
				r = r - 26
			}

		// In case of uppercase letters, we do the same
		} else if r >= 'A' && r <= 'Z' {
			r = r + 13
			// If we go past Z, go back to the beginning
			if r > 'Z' {
				r = r - 26
			}
		}

		// Digits are left unchanged here, as they do not match the letter conditions.
		// Their values are concatenated into the result string.
		result = result + string(r)
	}

	// at the end, the result is sent back to the main function CypherTool
	return result
}

// Function for applying reverse alphabet encryption, suitable for both encryption and decryption
func reverseAlphabet(text string) string {

	result := ""

	// For each value in the text
	for _, r := range text {
		// If it is a lowercase alphabetic value
		if r >= 'a' && r <= 'z' {
			// The position is mirrored as a lowercase
			position := r - 'a'
			r = 'z' - position
		// If it's a uppercase alphabet value
		} else if r >= 'A' && r <= 'Z' {
			// it does the same but as a uppercase
			position := r - 'A'
			r = 'Z' - position
		}

		// Digits are left unchanged here, as they do not match the letter conditions.
		// Their values are concatenated into the result string.
		result = result + string(r)
	}
	// at the end, the result is sent back to the main function CypherTool
	return result
}

// Function for applying the Chinese conversion, suitable for both encryption and decryption
func EncryptChinese(text string, toEncrypt bool) string {

	result := ""

	// For each value in the text
	for _, r := range text {
		// If the selected action is encryption
		if toEncrypt == true {
			// and if the value is lowercase
			if r >= 'a' && r <= 'z' {
				// convert it to its Chinese representation without translating the meaning
				distance := r - 'a'
				chinese := '淡' + distance
				result += string(chinese)

			// or if the value is uppercase
			} else if r >= 'A' && r <= 'Z' {
				// convert it to its Chinese representation without translating the meaning
				distance := r - 'A'
				chinese := '中' + distance
				result += string(chinese)

			} else {
				// Digits are left unchanged here, as they do not match the letter conditions.
				// Their values are concatenated into the result string.
				result += string(r)
			}
		
			// If the selection was to decrypt
		} else { 

			// and if the value is lowercase
			if r >= '淡' && r <= '淡'+25 {
				distance := r - '淡'
				latin := 'a' + distance
				result += string(latin)

			// or if the value is uppercase
			} else if r >= '中' && r <= '中'+25 {
				distance := r - '中'
				latin := 'A' + distance
				result += string(latin)

			} else {
				// Digits are left unchanged here, as they do not match the letter conditions.
				// Their values are concatenated into the result string.
				result += string(r)
			}
		}
	}
	// at the end, the result is sent back to the main function CypherTool
	return result
}

// Function used for handle the inputs given from the user
func getInput() (toEncrypt bool, encoding string, text string) {

scanner := bufio.NewScanner(os.Stdin)

	// we are creating variable scanner and calling the package bufio for reading the input and then calling the NewScanner from inside bufio. 
	// os goes to operation system and can read input typed from the keyboard
	var operation string

	// The terminal will wait until the user enter a valid input (1 or 2)
	for {
		fmt.Println("Choose what you want to do:\n\n- Type 1 if you want to encrypt\n- Type 2 if you want to decrypt")
		// Then the input is read with scanner library
		scanner.Scan()
		// It gets trimmed from empty spaces
		operation = strings.TrimSpace(scanner.Text())
		// If the input was 1, toEncrypt will store the boolean True
		if operation == "1" {
			toEncrypt = true
			break
		}
		// If the input was 2, toEncrypt will store the boolean False
		if operation == "2" {
			toEncrypt = false
			break
		}
		// If the input was invalid, the for loop goes back to the beginning and it prints an error message
		fmt.Println("\nInvalid choice.\nPlease, try again:\n---")
	}
	// It prints the user's choice
	if operation == "1" {
		fmt.Println("You choosed to: Encrypt\n---")
	} else {
		fmt.Println("You choosed to: Decrypt\n---")
	}
	
	for {
		// The terminal will wait for user input on the type of the chosen encryp/decrypr 
		fmt.Println ("Choose a cypher:\n\n- Type 1 if you want to use Rot13\n- Type 2 if you want to use Reverse Alphabet\n- Type 3 if you want to use Chinese Converter")
		// Then the input is read with scanner library
		scanner.Scan()
		// It gets trimmed from empty spaces
		encoding = strings.TrimSpace(scanner.Text())
		// If the input was 1, encoding will store the Rot13 choice, if it was 2, encoding will store the Reverse Alphabet choice, if it was 3, encoding will store the Chinese Converter choice
		if encoding == "1" ||encoding == "2" || encoding == "3" {
			break
		}
		// If the input was invalid, the for loop goes back to the beginning and it prints an error message
		fmt.Println ("\nInvalid choice.\nPlease, try again:\n---")	
	}	
	// It prints the user's choice
	if encoding == "1" {
		fmt.Println("You choosed: Rot13\n---")
	} else if encoding == "2" {
		fmt.Println("You choosed: Reverse Alphabet\n---")
	} else {
		fmt.Println("You choosed: Chinese Converter\n---")
	}
	
	for {
		// The terminal will wait for user input on the message the user wants to encrypt or decrypt 
		fmt.Println ("Enter your text:")
		// Then the input is read with scanner library
		scanner.Scan()
		// The message gets trimmed from empty spaces and it gets stored in the variable text
		text = strings.TrimSpace(scanner.Text())
			if text != "" {
				break
			}
		// Otherwise, if the input is invalid, an error message is printed and the for loop goes back to the beginning
		fmt.Println ("\nText field cannot be empty.\nPlease, try again:\n---")
	} 

	// At the end the 3 variables are sent back to the main function CypherTool
	return toEncrypt, encoding, text

}