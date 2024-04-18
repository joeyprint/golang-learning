package user

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

const (
	SUPER_ADMIN string = "SUPER_ADMIN"
	ADMIN       string = "ADMIN"
	CUSTOMER    string = "CUSTOMER"
)

var mockUsers = []User{
	{Id: "1", FirstName: "Wasin", LastName: "Wachirapusitanun", Role: SUPER_ADMIN},
	{Id: "2", FirstName: "Marina", LastName: "Hickman", Role: ADMIN},
	{Id: "3", FirstName: "Marissa", LastName: "Lee", Role: CUSTOMER},
	{Id: "4", FirstName: "Tomasz", LastName: "Sanford", Role: CUSTOMER},
}
