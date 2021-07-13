package main

import (
	"log"
	"os"

	"github.com/google/uuid"
)

func createIDS() uuid.UUID {
	id := uuid.New()

	return id
}

func writeToFile(id uuid.UUID) {
	file, err := os.Create("uuids.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for i := 0; i < 1_000_000; i++ {
		_, err := file.WriteString(createIDS().String() + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	writeToFile(createIDS())
}
