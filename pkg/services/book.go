package services

import (
	"fmt"

	"github.com/aaryanraj/booking-app/pkg/common/helpers"
	"github.com/aaryanraj/booking-app/pkg/structs"
)

func Book(remTickets *structs.Availability, bookings map[string][]structs.UserDataModel) {

	if !helpers.CheckAvalability(remTickets) {
		fmt.Println("Our conference is booked out at all locationa. See you next time")
	} else {
		//call the function to get user data
		userData := helpers.CollectUserData()

		//call the function to validate user
		isValidUserInput := helpers.ValidateUserInput(userData, remTickets)

		if isValidUserInput.AreDetailsValid {
			//call bookTickets func
			helpers.BookTockets(userData, remTickets, bookings)

			//Call getFirstNames func
			firstNames := helpers.GetFirstNames(bookings, userData.City)
			fmt.Printf("Here is the participant's list of %v : %v\n", userData.City, firstNames)

		} else {
			if !isValidUserInput.IsValidName {
				fmt.Println("You have entered a invalid user name. First & last name should be atleast 2 char long.")
			}
			if !isValidUserInput.IsValidEmail {
				fmt.Println("You have entered a invalid email. Email should have @ present")
			}
			if !isValidUserInput.IsValidCity {
				fmt.Println("Invalid city selectd. Please make sure valid city is selected from the list.")
			}
			if !isValidUserInput.IsValidUserTickets {
				fmt.Println("Tickets cant be zero or negative value")
			}
			if isValidUserInput.IsTicketOutOfBound {
				fmt.Printf("You can't book %v tickets in %v. Check ticket Avalibility and try again!!!\n", userData.UserTickets, userData.City)
			}
			fmt.Println("Please re-enter the details")
		}
	}
}
