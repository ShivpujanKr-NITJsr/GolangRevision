package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/shiv/podcast/auth"
	"github.com/shiv/podcast/user"
)

// need to make fisrt module initialize for creating package

func main() {
	auth.LoginWithCredentials("shiv", "password123")

	session := auth.GetSession()

	fmt.Println("Current session status:", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}

	// fmt.Println("User details:", user)

	color.Red(user.Email)
}

// go mod tidy remove all packages if not used or change to indirect /direct as used
