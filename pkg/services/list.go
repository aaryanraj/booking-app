package services

import (
	"fmt"

	"github.com/aaryanraj/booking-app/pkg/structs"
)

func List(bookings map[string][]structs.UserDataModel) {
	var city string
	var bookingList []structs.UserDataModel
	fmt.Println("Enter the city for whic you want to list bookings:")
	fmt.Scan(&city)

	if value, found := bookings[city]; found {
		bookingList = value

		fmt.Printf("\nThe Bookings list for %v\n", city)
		for _, value := range bookingList {
			fmt.Println(value)
		}
	} else {
		fmt.Println("\nCity Not Found")
		fmt.Println("Details did not match. please try again with correct details!")
	}

	fmt.Println("Enter any key to return to main menu.")
	var input string
	fmt.Scan(&input)
}
