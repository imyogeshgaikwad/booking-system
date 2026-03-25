package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go conference"

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber :=
			validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets, bookings = bookTicket(
				userTickets,
				firstName,
				lastName,
				email,
			)

			fmt.Printf("Slice length: %v\n", len(bookings))

			firstNames := getFirstNames(bookings)
			fmt.Printf("First names: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email must contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid number of tickets")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("Total tickets: %v, Remaining: %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}



func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint,
	firstName string, lastName string, email string) (uint, []string) {

	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation email at %v\n", email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("The whole slice: %v\n", bookings)

	return remainingTickets, bookings
}