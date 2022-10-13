package main

import (
	"fmt"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := vaildateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTickets {

			bookTicket(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("These are all the first names of bookings: %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our %v tickets are sold out. Please come back next year \n", conferenceName)
				break
			}
		} else {

			if !isValidName {
				fmt.Println("The firstname or lastname you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email you entered doesnot have @ sign")
			}
			if !isValidTickets {
				fmt.Println("The no of tickets you entered is invalid")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
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

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Println(bookings)

	fmt.Printf("Thank you  %v %v for booking %v tickets. You will receive a confirmation mail at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v \n", remainingTickets, conferenceName)
}
