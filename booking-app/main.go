package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" // Syntatic Sugar
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings = []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		// Validate firstName and lastName validation:
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// check userTickets input less than or equal to remainingTickets
		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}           // slices of first name
			for _, booking := range bookings { // underline to ignore a variable u don't want to use
				var names = strings.Fields(booking) // Splits the string with white space as separator.
				var firstName = names[0]            // Get the first name of the whole name
				firstNames = append(firstNames, firstName)
			}
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
