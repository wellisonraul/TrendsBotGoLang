package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	s := NewServer()
	defer s.Close()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	s.Start(port)
}
