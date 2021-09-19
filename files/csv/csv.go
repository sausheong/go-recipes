package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type User struct {
	Id        int
	firstName string
	lastName  string
	email     string
}

func main() {
	writeOne()
}

func writeOne() {
	file, err := os.Create("new_users2.csv")
	if err != nil {
		log.Println("Cannot create CSV file:", err)
	}
	defer file.Close()

	data := [][]string{
		{"id", "first_name", "last_name", "email"},
		{"1", "Sau Sheong", "Chang", "someemail@random.com"},
		{"2", "John", "Doe", "john@email.com"},
	}

	writer := csv.NewWriter(file)
	for _, row := range data {
		err = writer.Write(row)
		if err != nil {
			log.Println("Cannot write to CSV file:", err)
		}
		writer.Flush()
	}

}

func write() {
	file, err := os.Create("new_users.csv")
	if err != nil {
		log.Println("Cannot create CSV file:", err)
	}
	defer file.Close()

	data := [][]string{
		{"id", "first_name", "last_name", "email"},
		{"1", "Sau Sheong", "Chang", "someemail@random.com"},
		{"2", "John", "Doe", "john@email.com"},
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	if err != nil {
		log.Println("Cannot write to CSV file:", err)
	}

}

func read() {
	file, err := os.Open("users2.csv")
	if err != nil {
		log.Println("Cannot open CSV file:", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.Comment = '#'
	rows, err := reader.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}

	var users []User
	for _, row := range rows {
		id, _ := strconv.ParseInt(row[0], 0, 0)
		user := User{Id: int(id),
			firstName: row[1],
			lastName:  row[2],
			email:     row[3],
		}
		users = append(users, user)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
