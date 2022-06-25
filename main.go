package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Arguments map[string]string

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Age   uint   `json:"age"`
}

func Perform(args Arguments, writer io.Writer) error {
	if args["operation"] == "" {
		return fmt.Errorf("-operation flag has to be specified")
	}
	if args["fileName"] == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}

	operation := args["operation"]
	filename := args["filename"]
	fmt.Println(filename)
	switch operation {
	case "add":
		if args["item"] == "" {
			return fmt.Errorf("-item flag has to be specified")
		}
		item := args["item"]
		fmt.Println(item)
	case "list":
		fmt.Println("list")
	case "remove":
		if args["id"] == "" {
			return fmt.Errorf("-id flag has to be specified")
		}
		fmt.Println("remove")
	case "findById":
		if args["id"] == "" {
			return fmt.Errorf("-id flag has to be specified")
		}
		fmt.Println("find")
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
	filenameFlag := flag.String("filename", "", "")
	flag.Parse()

	return Arguments{
		"id":        *idFlag,
		"item":      *itemFlag,
		"operation": *operationFlag,
		"fileName":  *filenameFlag,
	}
}
