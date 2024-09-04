package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"talentpitch/src/shared/dependencies"
)

func main() {
	godotenv.Load()
	err := dependencies.BuildMainDependencies().Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
	}
}
