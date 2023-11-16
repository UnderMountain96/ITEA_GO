package main

import (
	"fmt"
	"os"
)

const (
	CommandList   = "list"
	CommandCreate = "create"
	CommandUpdate = "update"
	CommandDelete = "delete"
	CommandExit   = "exit"
	CommandHelp   = "help"
)

var users = map[int]string{}

func main() {
	printGoLogo()
	fmt.Println("Welcome to UDBS (User DataBase Shell)\n")

	fmt.Printf("Entry command '%s' to get a list commands\n", CommandHelp)
	fmt.Printf("Entry command '%s' or press Ctrl+C to exit\n\n", CommandExit)

	readCommand()
}

func readCommand() {
	for {
		fmt.Print("Entry command: ")
		var command string
		fmt.Scan(&command)
		fmt.Println()

		switch command {
		case CommandList:
			listUsers()
		case CommandCreate:
			createUser()
		case CommandUpdate:
			updateUser()
		case CommandDelete:
			deleteUser()
		case CommandExit:
			exit()
		case CommandHelp:
			help()
		default:
			fmt.Printf("Unknown command '%s', entry command '%s' to get a list commands\n", command, CommandHelp)
		}

		fmt.Println()
	}
}

func listUsers() {
	if len(users) == 0 {
		fmt.Println("List users is empty")
		return
	}

	fmt.Println("List users:\n")
	for id, name := range users {
		fmt.Printf("ID: %d,\tNAME: %s\n", id, name)
	}
}

func createUser() {
	fmt.Print("Entry ID to create user: ")
	var id int
	fmt.Scan(&id)

	_, exists := users[id]

	if exists {
		fmt.Printf("User by ID: %d - exist\n", id)
		return
	}

	fmt.Print("Entry user name: ")
	var name string
	fmt.Scan(&name)

	users[id] = name

	fmt.Printf("User by ID: %d NAME: %s - created\n", id, name)
}

func updateUser() {
	fmt.Print("Entry ID to update user: ")
	var id int
	fmt.Scan(&id)

	name, exists := users[id]

	if !exists {
		fmt.Printf("User by ID: %d - not found\n", id)
		return
	}

	fmt.Print("Entry new name: ")
	var newName string
	fmt.Scan(&newName)

	users[id] = newName

	fmt.Printf("User by ID: %d NAME: %s - set new name '%s'\n", id, name, newName)
}

func deleteUser() {
	fmt.Print("Entry ID to delete user: ")
	var id int
	fmt.Scan(&id)

	name, exists := users[id]

	if !exists {
		fmt.Printf("User by ID: %d - not found\n", id)
		return
	}

	delete(users, id)

	fmt.Printf("User by ID: %d, NAME: %s - deleted\n", id, name)
}

func help() {
	fmt.Println("List of available commands:\n")
	fmt.Println(CommandList, "\t- list users")
	fmt.Println(CommandCreate, "\t- create user")
	fmt.Println(CommandUpdate, "\t- update user")
	fmt.Println(CommandDelete, "\t- delete user")
	fmt.Println(CommandExit, "\t- exit the program ")
	fmt.Println(CommandHelp, "\t- list commands")
}

func exit() {
	os.Exit(1)
}

func printGoLogo() {
	fmt.Println(`-------------------------------------------------------------------------------------
                                 ++++++++++++++                 ++++++++++++         
                              ++++++++++++++++++++          ++++++++++++++++++++     
                            ++++++++++++++++++++++++     ++++++++++++++++++++++++    
                          +++++++++++++++++++++++++++   +++++++++++++++++++++++++++  
      +++++++++++++++++  +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ 
      ++++++++++++++++  +++++++++++         +++++++  ++++++++++++         ++++++++++ 
+++++++++++++++++++++  ++++++++++                   +++++++++++            ++++++++++
++++++++++++++++++++  ++++++++++     +++++++++++++++++++++++++              +++++++++
          ++++++++++  ++++++++++     ++++++++++++++++++++++++              ++++++++++
          ++++++++++  ++++++++++    +++++++++++++++++++++++++              ++++++++++
                      ++++++++++    +++++++++++++++++++++++++             ++++++++++ 
                       +++++++++++        +++++++++++++++++++++         +++++++++++  
                       ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  
                        +++++++++++++++++++++++++++  ++++++++++++++++++++++++++++    
                         ++++++++++++++++++++++++     ++++++++++++++++++++++++++     
                           ++++++++++++++++++++         +++++++++++++++++++++        
                             ++++++++++++++                +++++++++++++++           
-------------------------------------------------------------------------------------`)
}
