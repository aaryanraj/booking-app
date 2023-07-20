package main

import (
	"fmt"

	"github.com/aaryanraj/booking-app/pkg/common/constants"
	"github.com/aaryanraj/booking-app/pkg/common/helpers"
	"github.com/aaryanraj/booking-app/pkg/services"
	"github.com/aaryanraj/booking-app/pkg/structs"
)

var remainingTickets = &structs.Availability{
	Bengaluru: constants.BengaluruTickets,
	Mumbai:    constants.MumbaiTickets,
	Delhi:     constants.DelhiTickets,
	Kolkata:   constants.KolkataTickets,
}
var bookings = make(map[string][]structs.UserDataModel, 0)

func main() {
	var operation int

	//call the greetUser Func
	helpers.GreetUser()

	for {
		fmt.Println("Select the operation you want to perform:")
		fmt.Println("NOTE: Integer input expected!")
		fmt.Printf("1. Book Ticket \n2. Update Ticket \n3. Get Booking\n")
		fmt.Printf("4. List Bookings \n5. Delete Booking \n6. Check Availability\n7. Exit\n")
		fmt.Scan(&operation)

		switch operation {
		case 1:
			services.Book(remainingTickets, bookings)
		case 2:
			services.Update()
		case 3:
			services.Get(bookings)
		case 4:
			services.List(bookings)
		case 5:
			services.Delete()
		case 6:
			services.CheckAvalability(remainingTickets)
		case 7:
			return
		default:
			{
				fmt.Println("Invalid operation selected. Please try again!")
			}
		}
	}
}
