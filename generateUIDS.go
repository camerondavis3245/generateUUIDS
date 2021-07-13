package main

import (
	"log"
	"os"

	"github.com/google/uuid"
)

const SIZE = 1_000_000

//Generates uuids
func createIDS() uuid.UUID {
	id := uuid.New()

	return id
}

// Creates file and writes to file
func writeToFile(id uuid.UUID) {
	file, err := os.Create("uuids.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for i := 0; i < SIZE; i++ {
		_, err := file.WriteString(createIDS().String() + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	writeToFile(createIDS())
}
