package structs

//validateUserResp is returned by validateUser func
type ValidateUserResp struct {
	IsValidName        bool
	IsValidEmail       bool
	IsValidUserTickets bool
	IsTicketOutOfBound bool
	IsValidCity        bool
	AreDetailsValid    bool
}

//userData holds the details about the users
type UserDataModel struct {
	FirstName   string
	LastName    string
	Email       string
	UserTickets uint
	City        string
	BookingID   string
}

type Availability struct {
	Bengaluru uint
	Mumbai    uint
	Delhi     uint
	Kolkata   uint
}
