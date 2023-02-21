package entity

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`     /*,min:6,max=50" validate:"is-full-name*/
	Phone     string    `json:"phone" binding:"required"`    /*,min=10,max=10*/
	Password  string    `json:"password" binding:"required"` /*,min=8*/
	Addresses []Address `json:"addresses"`
	Balance   int       `json:"balance"`
}

type UserInfoRequest struct {
	ID int `json:"id"`
}

type UserLoginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserEditRequest struct {
	ID       int    `json:"id" binding:"required"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
}
