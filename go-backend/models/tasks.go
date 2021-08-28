package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	DueDate time.Time `json:"due_date"`
}

var tasks = []Task{
	{
		ID:      1,
		Title:   "タスクA",
		Content: "Aのタスクです",
		DueDate: time.Now(),
	},
	{
		ID:      2,
		Title:   "タスクB",
		Content: "Bのタスクです",
		DueDate: time.Now(),
	},
	{
		ID:      3,
		Title:   "タスクC",
		Content: "Cのタスクです",
		DueDate: time.Now(),
	},
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&tasks); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.String())

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}
