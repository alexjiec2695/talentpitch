package main

import (
	"github.com/joho/godotenv"
	"talentpitch/src/shared/dependencies"
)

func main() {
	godotenv.Load()
	dependencies.BuildMainDependencies()
}
