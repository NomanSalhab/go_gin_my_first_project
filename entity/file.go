package entity

type File struct {
	// gorm.Model // GORM model that contains the ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Filename string `json:"file_name" gorm:"not null"`   // Filename of the file. Cannot be null.
	UUID     string `json:"uuid" gorm:"unique;not null"` // UUID of the file. Must be unique and cannot be null.
}
