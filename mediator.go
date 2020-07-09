package main

import (
	"fmt"
	"time"
)

var chatRoom = &ChatRoom{}

type ChatRoom struct{}

func (*ChatRoom) ShowMessage(user *User, message string) {
	fmt.Println(time.Now(), "[", user, ": ", message, "]")
}

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(str string) {
	u.Name = str
}

func (u *User) SendMessage(message string) {
	chatRoom.ShowMessage(u, message)
}

func main() {
	robert := NewUser("Robert")
	john := NewUser("John")

	robert.SendMessage("hi John")
	john.SendMessage("hi Robert!")
}
