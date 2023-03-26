package main

import (
	"booking-app/helper"
	"fmt"
)

// Package Level Variables
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// initilize slice with map type
var bookings = make([]UserData, 0)

// Struct data type
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	// Call func greetUsers with param confName
	greetUsers()

	for {
		// get user input
		firstName, lastName, email, userTickets := getUserInput()

		// return validate user input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// check userTickets input less than or equal to remainingTickets
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			// call func firstNames and return values of list of firstName
			firstNames := getFirstNames()

			fmt.Printf("The first name of our bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Comeback next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short:")
			}
			if !isValidEmail {
				fmt.Println("Email you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	// execute code for booking New York conference tickets
	// case "London", "Berlin":
	// 	// execute code for London & Berlin conference tickets
	// case "Hong Kong", "Singapore":
	// 	// execute code for Hong Kong & Singapore conference tickets
	// case "Mexico City":
	// 	// execute code for Mexico City conference tickets
	// default:
	// 	fmt.Println("No valid city selected")
	// }

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}           // slices of first name
	for _, booking := range bookings { // underline to ignore a variable u don't want to use
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// as user for their info
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("Enter your number tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
