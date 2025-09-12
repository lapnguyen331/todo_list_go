package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json: "status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func main() {
	now := time.Now().UTC()
	item := TodoItem{1,"hgege","đây là test","Doing",&now, nil,}

	jsonData, err :=json.Marshal(item)

	if err != nil {
		fmt.Print(err)
		return
	}else{
		fmt.Println(string(jsonData))
	}

	var items2 TodoItem

	 json.Unmarshal(jsonData,&items2)
	fmt.Println(items2)
}