package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Arguments map[string]string

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Age   uint   `json:"age"`
}

var users []User

func Perform(args Arguments, writer io.Writer) error {
	if args["operation"] == "" {
		return fmt.Errorf("-operation flag has to be specified")
	}
	if args["fileName"] == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}

	operation := args["operation"]
	if !operationCheck(operation) {
		return fmt.Errorf("Operation " + operation + " not allowed!")
	}
	filename := args["fileName"]

	switch operation {
	case "add":
		if args["item"] == "" {
			return fmt.Errorf("-item flag has to be specified")
		}
		item := args["item"]

		users, err := addNewItem(item, filename, writer)
		if err != nil {
			log.Fatal(err)
		}
		return writeToFile(filename, users)

	case "list":
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(content, &users)
		if err != nil {
			log.Fatal(err)
		}
		usersToPrint, err := json.Marshal(users)
		if err != nil {
			log.Fatal(err)
		}
		writer.Write(usersToPrint)

	case "remove":
		if args["id"] == "" {
			return fmt.Errorf("-id flag has to be specified")
		}
		id := args["id"]
		users, err := removeUser(id, filename, writer)
		if err != nil {
			log.Fatal(err)
		}
		return writeToFile(filename, users)

	case "findById":
		if args["id"] == "" {
			return fmt.Errorf("-id flag has to be specified")
		}
		id := args["id"]
		if err := findById(id, filename, writer); err != nil {
			return fmt.Errorf("failed to find by ID %s: %w", filename, err)
		}
	}

	return nil
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {
	idFlag := flag.String("id", "", "")
	itemFlag := flag.String("item", "", "")
	operationFlag := flag.String("operation", "", "")
	filenameFlag := flag.String("fileName", "", "")
	flag.Parse()

	return Arguments{
		"id":        *idFlag,
		"item":      *itemFlag,
		"operation": *operationFlag,
		"fileName":  *filenameFlag,
	}
}
