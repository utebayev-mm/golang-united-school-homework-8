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
		return fmt.Errorf("operation flag not specified")
	}
	if args["fileName"] == "" {
		return fmt.Errorf("fileName flag not specified")
	}
	fmt.Println(args)
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
