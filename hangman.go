package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
	"time"
)

func addSpace(word string) string {
    var result []rune

    for _, char := range word {
        // Append the current character and a space to the result slice
        result = append(result, char, ' ')
    }

    // Convert the slice of runes back to a string
    return string(result)
}

func check(word string, guesses string) bool {
	guesses = strings.ReplaceAll(guesses, " ", "") //Fix guesses by removing spaces

	if word == guesses {
		return true
	}
	return false //Else
}

func clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported operating system")
	}
}

func processChar() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ")
		panic(err)
	}
	//Weird input fix
	runes := []rune(input)
	input = string(runes[0])

	//Make sure strings are lowercase for consistency
	input = strings.ToLower(input)
	return input
}

func processString() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ")
		panic(err)
	}
	//Fix input problem removing input spaces
	input = strings.TrimSpace(input)
	//Make sure strings are lowercase for consistency
	input = strings.ToLower(input)
	return input
}



func main() {
	tries := 6 //Head, arm, arm, body, leg, leg
	var guessed string
	var letters []string
	fmt.Println("Welcome to hangman!")
	fmt.Println("Enter the word you would like to be guessed...")
	word := processString()
	println("Clearing screen...")
	time.Sleep(1 * time.Second) //Wait 1 second then clear
	clear()

	//Create empty "_" string for each letter in the word
	for range word {
		guessed += "_ "
	}
	guessed = strings.TrimSpace(guessed) //Remove ending space

	fmt.Println("You will have", tries, "guesses")
	fmt.Println("Let's begin!")

	//Main loop
	//for tries != 6 { //Change to 0 when done
	for tries != 0 {
		fmt.Println("")
		fmt.Println(guessed)
		fmt.Println("Guess a letter...")
		guess := processChar() //Returns string

		//Check if letter has already been guessed, if so continue to next iteration
		if slices.Contains(letters, guess) {
			fmt.Println("You already guessed that letter", guess, "try again.")
			continue
		}

		//Check if letter (guess) is in the word (word)
		if strings.Contains(word, guess) {
			//The guessed letter exists somewhere in the word
			//Replace instances of "_" in guessed where the letter should be then continue loop

			//Get indexes in word where the letter is and then adjust (guessed accordingly)
			var indexes []int
			for index, letter := range word { //Check each letter
				if string(letter) == guess {
					//Add index number to indexes slice where the letter exists
					indexes = slices.Insert(indexes, len(indexes), index)
				}
			}

			newSlice := strings.Split(guessed, " ") //This is now a slice of strings "_"

			//Replace every instance of guess
			for i := 0; i < len(indexes); i++  {
				newSlice[indexes[i]] = guess
			}
			//Change slice of strings to just one string (guessed)
			guessed = ""
			for index := range newSlice {
				guessed += newSlice[index]
			}

			if check(word, guessed) { //Check if the word has been guessed
				fmt.Println("You solved it!")
				fmt.Println("Thanks for playing!")
				os.Exit(0)
			}			
			guessed = addSpace(guessed) //Revert guesses back to initial setup
			continue
		}

		//If you end up here the letter wasn't in the word
		fmt.Println("That letter isnt in that word!")
		//Add to letters that have been guessed and decrease tries
		letters = slices.Insert(letters, len(letters), guess)
		slices.Sort(letters)
		tries -= 1
		fmt.Println(tries, "attempts remaining.")
	}
	//Your attempts ran out
	laughing := "\U0001F602"
	fmt.Println("")
	fmt.Println("Looks like you couldn't hang", laughing, laughing, laughing)
	fmt.Println("The word was", word)
	fmt.Println("Thanks for playing!")
}