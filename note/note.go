package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title     string    `json:"title`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Println("Your ntoe titled %v has following content : \n\n %v")
}

func (note Note) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if content == "" {
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		Text:      content,
	}, nil
}
