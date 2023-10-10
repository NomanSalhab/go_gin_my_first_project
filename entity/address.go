package entity

type Address struct {
	ID        int     `json:"id"`
	UserId    int     `json:"user_id"`   /* binding:"min=2,max=200" validate:"is-cool" */ /* xml:"title" form:"title" validate:"email" binding:"required"*/
	Name      string  `json:"name"`      /* binding:"max=200" */
	Latitude  float32 `json:"latitude"`  /*  binding:"required,url" */
	Longitude float32 `json:"longitude"` /*  binding:"required,url" */
}

type UserAddressesRequest struct {
	UserId int `json:"user_id"`
}

type UserDeleteAddressRequest struct {
	UserId int `json:"user_id"`
	ID     int `json:"id"`
}

type OrderAddress struct {
	ID        int     `json:"id" binding:"required"`
	UserId    int     `json:"user_id"`   /* binding:"min=2,max=200" validate:"is-cool" */ /* xml:"title" form:"title" validate:"email" binding:"required"*/
	Name      string  `json:"name"`      /* binding:"max=200" */
	Latitude  float32 `json:"latitude"`  /*  binding:"required,url" */
	Longitude float32 `json:"longitude"` /*  binding:"required,url" */
}

type AddAddressRequest struct {
	UserId    int     `json:"user_id"`
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
