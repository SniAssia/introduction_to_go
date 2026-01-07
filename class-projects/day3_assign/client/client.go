package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type student struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address string   `json:"address"`
	Courses []string `json:"courses"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/students")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	message, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	users := []student{}
	err = json.Unmarshal(message, &users)
	if  err != nil {
		log.Fatal(err)
	}

	fmt.Println("got students:", len(users))
	for _, v := range users {
		printStudent(v)
	}

	newStudent := student{
		Name:"assia",
		Age:10,
		Address:"assouar mek",
		Courses:[]string{"ai", "ml"},
	}

	body, err := json.Marshal(newStudent)
	if err != nil {
		log.Fatal(err)
	}

	resp1, err := http.Post(
		"http://localhost:8080/students",
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp1.Body.Close()

	fmt.Println("POST status:", resp1.Status)

	resp, err = http.Get("http://localhost:8080/students")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	message, _ = io.ReadAll(resp.Body)
	json.Unmarshal(message, &users)

	fmt.Println("After insert:")
	for _, v := range users {
		printStudent(v)
	}
}

func printStudent(v student) {
	fmt.Println("name:", v.Name)
	fmt.Println("age:", v.Age)
	fmt.Println("address:", v.Address)
	fmt.Println("courses:", v.Courses)
	fmt.Println("=========== ok ============")
}
