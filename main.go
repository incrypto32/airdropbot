package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Llongfile)

	// Load ENV FIle
	if err := godotenv.Load(); err != nil {

		log.Panic("Failed loading env variables", err)
	}

	fmt.Println("Hello World")
	fmt.Println(os.Getenv("BOT_TOKEN"))
}
