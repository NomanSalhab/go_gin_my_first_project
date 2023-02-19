package entity

type Address struct {
	ID        int8    `json:"id"`
	UserId    int8    `json:"user_id"`   /* binding:"min=2,max=200" validate:"is-cool" */ /* xml:"title" form:"title" validate:"email" binding:"required"*/
	Name      string  `json:"name"`      /* binding:"max=200" */
	Latitude  float32 `json:"latitude"`  /*  binding:"required,url" */
	Longitude float32 `json:"longitude"` /*  binding:"required,url" */
}
