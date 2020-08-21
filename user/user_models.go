package user

//User :
type User struct {
	FirstName string
	LastName  string
	Address   Address
	UserID    string
}

//Address :
type Address struct {
	Home   AddressDetails
	Office AddressDetails
	Email  string
	Phone  string
}

//AddressDetails :
type AddressDetails struct {
	Address  string
	Location Location
	Type     string
}

//Location :
type Location struct {
	Lattitude float64
	Longitude float64
}
