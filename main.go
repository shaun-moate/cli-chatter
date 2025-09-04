package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Initializing IRC CLI Chat...")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Printf("You said: %s\n", input)
		fmt.Print("> ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
