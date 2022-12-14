package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// package level variables
var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// empty list of maps
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// to prevent main thread from exiting program before sendingTicket" is done executing
var wg = sync.WaitGroup{}

func main() {
	greetUsers()
	// fmt.Printf("conferencTickets is %T, remainingTickets is %T and conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNmber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNmber {
			bookTickets(userTickets, firstName, lastName, email)
			// "go" creates a "Goroutine" and makes code concurrent by having "sendTicket" execute in separate thread
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

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
			if !isValidEmail {
				fmt.Printf("The email address you provided is invalid!\n")
			}
			if !isValidTicketNmber {
				fmt.Printf("The number of tickets you enterd can not be proces!\n")
			}
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tikets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// simulating delay in execution
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending ticket(s):\n %v to email address %v\n", ticket, email)
	fmt.Printf("############\n")
	// removes thread from waiting list
	wg.Done()
}
