package configs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {

	exp, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Dir(exp)

	fmt.Println(path)
	log.Print(path)
	fmt.Println(filepath.Join(path))

	if err := godotenv.Load(filepath.Join(path, ".env")); err != nil {
		log.Fatal(path)
	}
	return os.Getenv("MONGO_URI")
}
