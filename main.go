package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

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

		isValidName, isValidEmail, isValidTicketNumber :=
			validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
			fmt.Printf("You will receive a confirmation email at %v\n", email)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			fmt.Printf("The whole slice: %v\n", bookings)
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

func greetUsers(confName string, confTickets int, remainTickets uint) {
	fmt.Printf("Welcome to %v\n", confName)
	fmt.Printf("Total tickets: %v, Remaining: %v\n", confTickets, remainTickets)
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

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}