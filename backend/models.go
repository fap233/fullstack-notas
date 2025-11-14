package main

import "time"

type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type Note struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	UserID string `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}	

type LoginResponse struct {
	Token string `json:"token"`
	User User `json:"user"`
}
