package main

func operationCheck(operation string) bool {
	if operation != "add" && operation != "remove" && operation != "list" && operation != "findById" {
		return false
	}
	return true
}
