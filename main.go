package main

import (
	"fmt"
	"strings"
)

// package level variables
var conferenceName = "Go conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	// fmt.Printf("conferencTickets is %T, remainingTickets is %T and conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	for {
		firstName, lastName, email, userTickets:= getUserInput()
		isValidName, isValidEmail, isValidTicketNmber:= validateUserInput(firstName, lastName, email, userTickets)
		
		if isValidName && isValidEmail && isValidTicketNmber {
			bookTickets( userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are all the bookings: %v\n", firstNames)	

			if remainingTickets == 0 {
				fmt.Printf("Our conference is fully booked. Come back next year.\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("The irst name or last name you entered is too short!\n")
			}
			if !isValidEmail{
				fmt.Printf("The email address you provided is invalid!\n")
			}
			if !isValidTicketNmber {
				fmt.Printf("The number of tickets you enterd can not be proces!\n")
			}
		}
	}
}


func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames()[]string {
	firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint)(bool, bool, bool) {
	isValidName :=len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNmber := userTickets > 0 && userTickets <= remainingTickets
		return isValidName, isValidEmail, isValidTicketNmber
}

func getUserInput()(string, string, string, uint) {
	var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets
}

func bookTickets( userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tikets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}