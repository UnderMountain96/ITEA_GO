package main

import (
	"fmt"
	"os"
)

const (
	cmdHelp   = "help"
	cmdList   = "list"
	cmdCreate = "create"
	cmdUpdate = "update"
	cmdDelete = "delete"
	cmdExit   = "exit"
)

var cmdDescriptions = map[string]string{
	cmdHelp:   "this help",
	cmdList:   "show all users in database",
	cmdCreate: "create new user and add to database",
	cmdUpdate: "update user data in database",
	cmdDelete: "delete user from database",
	cmdExit:   "exit the program",
}

var users = map[int]string{}

func main() {
	printWelcome()

	for {
		cmd, err := getCommand()
		if err != nil {
			fmt.Printf("Command input error: %s", err)
		}
		handleCommand(cmd)
	}
}

func printWelcome() {
	const msg = "Welcome to UDBS (User DataBase Shell) v2"
	borderLen := len(msg) + 2

	printBorder(borderLen)
	fmt.Printf(" %s \n", msg)
	printBorder(borderLen)

	fmt.Printf("Entry command %q to get list of available commands.\n", cmdHelp)
}

func printBorder(len int) {
	for i := 0; i < len; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}

func getCommand() (string, error) {
	fmt.Printf("\nEnter command > ")

	var cmd string
	_, err := fmt.Scan(&cmd)
	if err != nil {
		return "", fmt.Errorf("getCommand: invalid command value entered: %w", err)
	}

	return cmd, nil
}

func handleCommand(cmd string) {
	switch cmd {
	case cmdHelp:
		handleHelpCmd()
	case cmdList:
		handleListCmd()
	case cmdCreate:
		if err := handleCreateCmd(); err != nil {
			fmt.Printf("Failed cteate user: %s", err)
			return
		}
	case cmdUpdate:
		if err := handleUpdateCmd(); err != nil {
			fmt.Printf("Failed update user: %s", err)
			return
		}
	case cmdDelete:
		if err := handleDeleteCmd(); err != nil {
			fmt.Printf("Failed delete user: %s", err)
			return
		}
	case cmdExit:
		handleExitCmd()
	default:
		fmt.Printf("Unknown command: %s\n\n", cmd)
		handleHelpCmd()
	}
}

func handleListCmd() {
	if len(users) == 0 {
		fmt.Println("Database is empty")
		return
	}

	fmt.Println("Users in database:\n")
	for id, name := range users {
		fmt.Printf("ID: %d,\tNAME: %s\n", id, name)
	}
}

func handleCreateCmd() error {
	fmt.Println("Please provide info about new user:")

	fmt.Print("ID: ")
	var id int
	_, err := fmt.Scan(&id)
	if err != nil {
		return fmt.Errorf("handleCreateCmd: invalid ID value entered: %w", err)
	}

	_, exists := users[id]

	if exists {
		fmt.Printf("User with ID: %d - is already exist\n", id)
		return nil
	}

	fmt.Print("Name: ")
	var name string
	_, err = fmt.Scan(&name)
	if err != nil {
		return fmt.Errorf("handleCreateCmd: invalid name value entered: %w", err)
	}

	users[id] = name

	fmt.Printf("User with ID: %d NAME: %s - successfully created!\n", id, name)

	return nil
}

func handleUpdateCmd() error {
	fmt.Print("Entry user ID to update: ")
	var id int
	_, err := fmt.Scan(&id)
	if err != nil {
		return fmt.Errorf("handleUpdateCmd: invalid ID value entered: %w", err)
	}

	name, exists := users[id]

	if !exists {
		fmt.Printf("User with ID: %d - not found\n", id)
		return nil
	}

	fmt.Print("Entry new name: ")
	var newName string
	_, err = fmt.Scan(&newName)
	if err != nil {
		return fmt.Errorf("handleUpdateCmd: invalid name value entered: %w", err)
	}

	users[id] = newName

	fmt.Printf("User with ID: %d NAME: %s - successfully set new name %q!\n", id, name, newName)

	return nil
}

func handleDeleteCmd() error {
	fmt.Print("Entry ID to delete user: ")
	var id int
	_, err := fmt.Scan(&id)
	if err != nil {
		return fmt.Errorf("handleDeleteCmd: invalid ID value entered: %w", err)
	}

	name, exists := users[id]

	if !exists {
		fmt.Printf("User with ID: %d - not found\n", id)
		return nil
	}

	delete(users, id)

	fmt.Printf("User with ID: %d, NAME: %s - successfully deleted!\n", id, name)

	return nil
}

func handleHelpCmd() {
	fmt.Println("Available commands:")

	for cmd, description := range cmdDescriptions {
		fmt.Printf("%s \t- %s.\n", cmd, description)
	}
}

func handleExitCmd() {
	os.Exit(1)
}
