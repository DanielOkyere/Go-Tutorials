package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content`
	CreatedAt time.Time `json:"created_at`
}

func (n Note) Display() {
	fmt.Printf("Your note titled:%v has the following content:\n\n%v\n\n",n.title,n.content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.title, " ", "_")
	fileName = strings.ToLower(fileName)
	json, err := json.Marshal(n)
	if err != nil {
		return err
	}
	
	return os.WriteFile(fileName, json, 0644)
	
}

func New(title, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("Invalid Input")
	}
	return Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}