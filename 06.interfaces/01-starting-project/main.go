package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/todo/note"
	"example.com/todo/todo"
)

type File interface {
	Save() error
	Display()
}

func SaveFile(file File) error {
	file.Save()

	err := file.Save()

	if err != nil {
		fmt.Println("Saving the file failed.")
		return err
	}

	fmt.Println("Saving the file succeeded!")
	return nil
}

func DisplayFile(file File) {
	fmt.Println("-----------------------------------------------")
	file.Display()
	fmt.Println("-----------------------------------------------")
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

	DisplayFile(todo)
	SaveFile(todo)
}

func createNote() {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	DisplayFile(userNote)
	SaveFile(userNote)
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
