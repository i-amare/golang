package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/todo/note"
	"example.com/todo/todo"
)

type Saver interface {
	Save() error
	Display(any)
}

func main() {
	createTodo()
	fmt.Println("-----------------------------------------------")
	createNote()
}

func createTodo() {
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	fmt.Println("Saving the todo succeeded!")
}

func createNote() {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
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
