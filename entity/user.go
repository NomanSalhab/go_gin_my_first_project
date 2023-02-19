package entity

type User struct {
	ID        int8      `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Addresses []Address `json:"addresses"`
	Balance   float32   `json:"balance"`
}

type UserInfoRequest struct {
	ID int8 `json:"id"`
}
