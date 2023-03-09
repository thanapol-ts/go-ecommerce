package intializer

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("ENV")
	fmt.Println("env == ", env)
	if "" == env {
		env = "dev"
	}

	err := godotenv.Load("./env/" + env + ".env")
	if err != nil {
		log.Fatal("ERROR loading .env file")
	}
}
