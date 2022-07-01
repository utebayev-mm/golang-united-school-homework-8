package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func operationCheck(operation string) bool {
	if operation != "add" && operation != "remove" && operation != "list" && operation != "findById" {
		return false
	}
	return true
}

func addNewItem(item, filename string, writer io.Writer) ([]User, error) {
	var newUser User
	if err := json.Unmarshal([]byte(item), &newUser); err != nil {
		return nil, fmt.Errorf("JSON failed to unmarshal: %w", err)
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var users []User
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", filename, err)
	}
	if err := json.Unmarshal(bytes, &users); err != nil {
		users = make([]User, 0, 1)
	} else {
		for i := 0; i < len(users); i++ {
			if users[i].ID == newUser.ID {
				writer.Write([]byte("Item with id " + users[i].ID + " already exists"))
				return users, nil
			}
		}
	}
	users = append(users, newUser)
	return users, nil
}

func removeUser(id, filename string, writer io.Writer) ([]User, error) {
	var users []User
	bytes, _ := os.ReadFile(filename)
	if err := json.Unmarshal(bytes, &users); err != nil {
		return nil, fmt.Errorf("JSON failed to unmarshal: %w", err)
	}

	newUsers := make([]User, 0, len(users)-1)
	for _, user := range users {
		if user.ID != id {
			newUsers = append(newUsers, user)
		}
	}

	if len(newUsers) == len(users) {
		writer.Write([]byte("Item with id " + id + " not found"))
		return users, nil
	}
	return newUsers, nil
}

func findById(id, filename string, writer io.Writer) error {
	var users []User
	bytes, _ := os.ReadFile(filename)
	if err := json.Unmarshal(bytes, &users); err == nil {
		for i := 0; i < len(users); i++ {
			if users[i].ID == id {
				bytes, err := json.Marshal(users[i])
				if err != nil {
					return fmt.Errorf("JSON failed to marshal: %w", err)
				}
				writer.Write(bytes)
			}
		}
	}
	return nil
}

func writeToFile(filename string, users []User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("JSON failed to marshal: %w", err)
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", filename, err)
	}

	if _, err = file.Write(bytes); err != nil {
		return fmt.Errorf("failed to write data to %s: %w", filename, err)
	}

	if err = file.Close(); err != nil {
		return fmt.Errorf("failed to close %s: %w", filename, err)
	}
	return nil
}
