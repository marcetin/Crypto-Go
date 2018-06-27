package main

import (
	"fmt"

	"github.com/mitsukomegumi/FakeCrypto/src/database"
)

func main() {
	db, err := database.ReadDatabase("127.0.0.1")

	if err != nil {
		panic(err)
	}

	acc, err := database.FindAccount(db, "test")

	if err != nil {
		panic(err)
	}

	fmt.Println(acc)
}
