package main

import (
	"fmt"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("Saving the todo failed")
		return
	}
	fmt.Println("Saving the todo succeeded!")

	userNote, err := note.New(title, content)

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succeeded!")
}

func getNoteData() {

}
