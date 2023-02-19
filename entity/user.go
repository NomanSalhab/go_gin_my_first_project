package entity

type User struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Addresses []Address `json:"addresses"`
	Balance   float32   `json:"balance"`
}
