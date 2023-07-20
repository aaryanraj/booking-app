package helpers

import (
	"fmt"
	"strings"

	"github.com/aaryanraj/booking-app/pkg/common/constants"
	"github.com/aaryanraj/booking-app/pkg/structs"

	"github.com/google/uuid"
)

//validateUserInput func validates the data entered by the user
func ValidateUserInput(uData structs.UserDataModel,
	remTickets *structs.Availability) structs.ValidateUserResp {
	//validating user input
	var isTicketOutOfBound bool
	isValidName := len(uData.FirstName) >= 2 && len(uData.LastName) >= 2
	isValidEmail := strings.Contains(uData.Email, "@")
	isValidUserTickets := uData.UserTickets > 0
	var isInValidCity bool
	switch uData.City {
	case "Mumbai", "mumbai":
		{
			isTicketOutOfBound = uData.UserTickets > remTickets.Mumbai
		}
	case "Delhi", "delhi":
		{
			isTicketOutOfBound = uData.UserTickets > remTickets.Delhi
		}
	case "Kolkata", "kolkata":
		{
			isTicketOutOfBound = uData.UserTickets > remTickets.Kolkata
		}
	case "Bengaluru", "bengaluru":
		{
			isTicketOutOfBound = uData.UserTickets > remTickets.Bengaluru
		}
	default:
		isInValidCity = true
	}

	areDetailsValid := isValidName && isValidEmail && isValidUserTickets && !isInValidCity && !isTicketOutOfBound
	return structs.ValidateUserResp{
		IsValidName:        isValidName,
		IsValidEmail:       isValidEmail,
		IsValidUserTickets: isValidUserTickets,
		IsTicketOutOfBound: isTicketOutOfBound,
		IsValidCity:        !isInValidCity,
		AreDetailsValid:    areDetailsValid,
	}
}

// collectUserData function Collects the user's data
func CollectUserData() structs.UserDataModel {
	var uData structs.UserDataModel

	fmt.Println("Enter your first name:")
	fmt.Scan(&uData.FirstName)

	fmt.Println("Enter the last name:")
	fmt.Scan(&uData.LastName)

	fmt.Println("Enter your email id:")
	fmt.Scan(&uData.Email)

	fmt.Println("Select the city from the lsit, where you want to attend:")
	fmt.Printf("Mumbai \nBengaluru \nDelhi \nKolkata\n")
	fmt.Scan(&uData.City)

	fmt.Println("Enter the number of tickets required:")
	fmt.Scan(&uData.UserTickets)

	return uData
}

// greetUser function greets the user
func GreetUser() {
	fmt.Printf("Welcome to %v booking application\n", constants.ConferenceName)
	fmt.Println("Get your tickets here to attend")
}

// getFirstNames func gets only the first names of the attendees
func GetFirstNames(bookings map[string][]structs.UserDataModel, city string) []string {
	//fmt.Println(bookings)
	firstNames := []string{}
	pList, _ := bookings[city]

	for _, participant := range pList {
		name := participant.FirstName
		firstNames = append(firstNames, name)
	}
	return firstNames
}

func BookTockets(userData structs.UserDataModel,
	remTickets *structs.Availability, bookings map[string][]structs.UserDataModel) {
	newUUID := uuid.New()
	userData.BookingID = newUUID.String()

	city := userData.City

	switch city {
	case "Mumbai", "mumbai":
		{
			remTickets.Mumbai = remTickets.Mumbai - userData.UserTickets
			fmt.Println("The Conference will be held at 'Hotel Taj' on 25th of this month")
			fmt.Println("The Conference starts at 12 noon, Entry closes at 11:55 AM")
		}
	case "Delhi", "delhi":
		{
			remTickets.Delhi = remTickets.Delhi - userData.UserTickets
			fmt.Println("The Conference will be held at 'Hotel JW Marriot' on 27th of this month")
			fmt.Println("The Conference starts at 3 PM, Entry closes at 02:55 PM")
		}
	case "Kolkata", "kolkata":
		{
			remTickets.Kolkata = remTickets.Kolkata - userData.UserTickets
			fmt.Println("The Conference will be held at 'Radisson Blu' on 29th of this month")
			fmt.Println("The Conference starts at 12 noon, Entry closes at 11:55 AM")
		}
	case "Bengaluru", "bengaluru":
		{
			remTickets.Bengaluru = remTickets.Bengaluru - userData.UserTickets
			fmt.Println("The Conference will be held at 'Hotel JW Marriot' on 30th of this month")
			fmt.Println("The Conference starts at 3 PM, Entry closes at 02:55 PM")
		}
	}

	bookings[city] = append(bookings[city], userData)

	fmt.Printf("Thankyou %v %v for booking %v tickets, a confirmation email will be sent to %v\n",
		userData.FirstName, userData.LastName, userData.UserTickets, userData.Email)
	fmt.Printf("Your booking ID: %v\n", newUUID)
}

func CheckAvalability(remTickets *structs.Availability) bool {
	if remTickets.Bengaluru > 0 || remTickets.Delhi > 0 ||
		remTickets.Kolkata > 0 || remTickets.Mumbai > 0 {
		return true
	}
	return false
}
