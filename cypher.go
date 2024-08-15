package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// rot13 encryption/decryption function
// encryption and decription are reversible functions
func rot13(s string) string {
	// This function loops through every character of the string s and checks if the given charachter is between 'a' and 'z' or 'A' and 'Z'
	// If the character is between 'a' and 'z' or 'A' and 'Z', the function executes.

	result := make([]string, 0, len(s))
	for _, chr := range s {
		if 'a' <= chr && chr <= 'z' {
			chr = ((chr - 'a' + 13) % 26) + 'a'
			// Converts chr from ASCII value to a number between 0-25, corresponding to its position in the alphabet. Then adds 13, shifting the letter 13 positions forward.
			// %26 wraps around the alphabet, if the shift goes past 'z'. And finally adds 'a', to convert the number back into a ASCII character.
		}
		if 'A' <= chr && chr <= 'Z' {
			chr = ((chr - 'A' + 13) % 26) + 'A'
		}
		result = append(result, string(chr))
		// Saves chr in the result variable.
	}
	output := strings.Join(result, "")
	return output
}

func reverse(plaintext string) string {

	var ciphertext string
	// String is array of runes. Encrypting each element of array, like in task "Reverse Alphabet"
	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			// Apply the cipher to upcase letter
			ciphertext += string('Z' - (char - 'A'))

		} else if char >= 'a' && char <= 'z' {
			// Apply the cipher to lowcase letter
			ciphertext += string('z' - (char - 'a'))

		} else {
			// If not letter
			ciphertext += string(char)
		}
	}

	return ciphertext
}

// The function shows that the string consists only of letters
func isOnlyLetters(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}

// Get key from terminal
func getKey() string {
	reader := bufio.NewReader(os.Stdin)
	// Entering the key
	var key string
	for {
		fmt.Printf("\nEnter the key: ")
		key, _ = reader.ReadString('\n')
		key = strings.TrimSpace(key)
		if isOnlyLetters(key) {
			break
		} else {
			fmt.Println("The key must consist of Latin letters only. Try again.")
		}
	}
	key = strings.ToUpper(key)
	return key
}

// Function to encrypt a string using the Vigenere cipher
func encrypt_vigenere(plaintext, key string) string {

	var ciphertext string
	keyIndex := 0

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			// applying the cipher only to letters
			shift := key[keyIndex] - 'A'
			encryptedChar := (char-'A'+rune(shift))%26 + 'A'
			ciphertext += string(encryptedChar)

			// Shifting the key index
			keyIndex = (keyIndex + 1) % len(key)
		} else if char >= 'a' && char <= 'z' {
			// Applying the cipher to lowercase letters
			shift := key[keyIndex] - 'A'
			encryptedChar := (char-'a'+rune(shift))%26 + 'a'
			ciphertext += string(encryptedChar)

			// Shifting the key index
			keyIndex = (keyIndex + 1) % len(key)
		} else {
			// If the symbol is not a letter, just add it
			ciphertext += string(char)
		}
	}

	return ciphertext
}

// Function to decrypt a string encrypted using the Vigenere cipher
func decrypt_vigenere(ciphertext, key string) string {
	var plaintext string

	keyIndex := 0

	for _, char := range ciphertext {
		if char >= 'A' && char <= 'Z' {
			// applying the cipher only to letters
			shift := key[keyIndex] - 'A'
			decryptedChar := (char-'A'-rune(shift)+26)%26 + 'A'
			plaintext += string(decryptedChar)
			// Shifting the key index
			keyIndex = (keyIndex + 1) % len(key)
		} else if char >= 'a' && char <= 'z' {
			// Applying the cipher to lowercase letters
			shift := key[keyIndex] - 'A'
			decryptedChar := (char-'a'-rune(shift)+26)%26 + 'a'
			plaintext += string(decryptedChar)
			// Shifting the key index
			keyIndex = (keyIndex + 1) % len(key)
		} else {
			// If the symbol is not a letter, just add it
			plaintext += string(char)
		}
	}
	return plaintext
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {
	return rot13(s)
}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {
	return rot13(s)
}

// Encrypt the message with reverse
func encrypt_reverse(s string) string {
	return reverse(s)
}

// Decrypt the message with reverse
func decrypt_reverse(s string) string {
	return reverse(s)
}

// Get the input data required for the operation
func getInput() (toEncrypt bool, encoding string, message string) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Cypher Tool!")

	for {
		//Get direction
		fmt.Println("")
		fmt.Println("Select an operation (1/2):")
		fmt.Println("1. Encrypt.")
		fmt.Println("2. Decrypt.")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}
		input = strings.TrimSpace(input)

		if input == "1" {
			toEncrypt = true
			break
		} else if input == "2" {
			toEncrypt = false
			break
		} else {
			fmt.Println("Only numbers 1 and 2 are accepted.")
		}
	}

	for {
		//Get type of cypher
		fmt.Println("")
		fmt.Println("Select cypher (1/3):")
		fmt.Println("1. ROT13.")
		fmt.Println("2. Reverse.")
		fmt.Println("3. Vigenère.")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}
		input = strings.TrimSpace(input)

		if input == "1" {
			encoding = "ROT13"
			break
		} else if input == "2" {
			encoding = "Reverse"
			break
		} else if input == "3" {
			encoding = "Vigenere"
			break
		} else {
			fmt.Println("Only numbers 1, 2 and 3 are accepted.")
		}
	}
	//Get string
	for {
		fmt.Println("Enter the message:")
		message, _ = reader.ReadString('\n')
		message = strings.TrimSpace(message)
		if message == "" {
			fmt.Println("Please add your message.")
			fmt.Println("")
		} else {
			break
		}
	}

	return toEncrypt, encoding, message

}

func main() {
	toEncrypt, encoding, message := getInput()
	var text string
	if toEncrypt == true {

		switch encoding {
		case "ROT13":
			text = encrypt_rot13(message)
		case "Reverse":
			text = encrypt_reverse(message)
		case "Vigenere":
			text = encrypt_vigenere(message, getKey())
		}
		fmt.Printf("\nEncrypted message using %v: \n%v \n", encoding, text)
	} else {
		switch encoding {
		case "ROT13":
			text = decrypt_rot13(message)
		case "Reverse":
			text = decrypt_reverse(message)
		case "Vigenere":
			text = decrypt_vigenere(message, getKey())
		}
		fmt.Printf("\nDecrypted message using %v: \n%v \n", encoding, text)
	}
}
