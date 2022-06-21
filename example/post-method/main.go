package main

import (
	"encoding/json"
	"log"

	"github.com/muhfaris/request"
)

type User struct {
	Name string
	Job  string
}

func main() {
	user := User{
		Name: "faris",
		Job:  "software developer",
	}

	raw, err := json.Marshal(user)
	if err != nil {
		log.Printf("error marshal the user data, %v", err)
		return
	}

	config := &request.Config{
		URL:         "https://reqres.in/api/users",
		Body:        raw,
		ContentType: "application/json",
	}

	resp := config.Post()
	if resp.Error != nil {
		log.Printf("error create new post user, %v", resp.Error)
		return
	}

	log.Println("Resp")
	log.Println(string(resp.Body))
}
