package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// User struct with properties mapped to json
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Initialize an slice of user pointers
var (
	users []*User
)

// Return the slice of user pointer
func GetUsers() []*User {
	return users
}

// Add new user to the slice
func AddUser(u User) {
	users = append(users, &u)
}

// Read the users from the json file
func ReadUsersFromJsonFile(file string) {
	// Open the file
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	//Read all the content of the file
	bytesFile, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	//Parse the json bytes into slices of users
	json.Unmarshal(bytesFile, &users)
}

// Write the user slices into the json file
func WritesUserToJsonFile(file string) {
	data, err := json.MarshalIndent(&users, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile(file, data, 0644)
}
