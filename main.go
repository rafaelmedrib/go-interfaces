package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()
	err = note.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return
	}

	fmt.Println("saving the note succeeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
