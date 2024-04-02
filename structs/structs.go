package main

import (
	"fmt"

	"example.com/note/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthday := getUserData("Please enter your birthday (MM/DD/YY): ")

	var appUser *user.User

	// appUser = &user.User{}

	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthday,
	// 	createdAt: time.Now(),
	// }

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthday)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("tony@example.com", "test123")
	admin.User.OutputUserDetails()
	admin.User.ClearUserName()
	admin.User.OutputUserDetails()

	// outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(infoText string) string {
	var userInput string
	fmt.Print(infoText)
	fmt.Scanln(&userInput)

	return userInput
}
