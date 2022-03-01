package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/lpalma92/learning-go/models"
)

func main() {

	// Define a flag to allow user add parameters on execution
	new := flag.Bool("new", false, "Add new user to the json file")
	// Define a flag to show the list of users
	list := flag.Bool("list", false, "List all user loades from the json file")
	// Define a flag to set the users json file
	file := flag.String("file", "users.json", "Json file where the users will be readed or writed")

	//Parse all define flags to pointers
	flag.Parse()

	//Validate if file exist if not created
	//TODO

	// Call the funcion in the models package to read users from a json file and load in to memory
	models.ReadUsersFromJsonFile(*file)

	// if the new flag is true add new user to the json
	if *new {
		// Ask to the user to add their information to store on the json file
		AddNewUser(*file)
	}

	// If the list flag is true show the  user
	if *list {
		ListUsers()
	}
}

func ListUsers() {
	// Call method to get the user loaded from memorie
	user := models.GetUsers()
	for u := range user {
		fmt.Printf("Name: '%v', Email: '%v' \n", user[u].Name, user[u].Email)
	}
}

func AddNewUser(file string) {
	// Ask to the user to add their information to store on the json file
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter your email: ")
	email, _ := reader.ReadString('\n')

	//Create a user and trim the user information from caracters added from reader
	user := models.User{}
	user.Name = strings.TrimSuffix(name, "\r\n")
	user.Email = strings.TrimSuffix(email, "\r\n")

	//Add the user to the slice
	models.AddUser(user)

	//Call the method to store the user to the file
	models.WritesUserToJsonFile(file)

	fmt.Println("User added to json file")
}
