package main

import "time"

type Todo struct {
	ID   int
	Task string
}

type UpdateTodoInput struct {
	Task *string `json:"task"`
	Done *bool   `json:"done"`
}

type TodoResp struct {
	ID        int       `json:"ID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Task      string    `json:"task"`
	Done      bool      `json:"done"`
}
