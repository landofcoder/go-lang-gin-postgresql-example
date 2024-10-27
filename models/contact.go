package models

type Contact struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email" gorm:"unique"`
	JobTitle    string `json:"job_title"`
	Company     string `json:"company"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	Country     string `json:"country"`
	Tags        string `json:"tags"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateContactInput struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email" binding:"required"`
	JobTitle    string `json:"job_title"`
	Company     string `json:"company"`
	Address     string `json:"address"`
	City        string `json:"City"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	Country     string `json:"country"`
	Tags        string `json:"tags"`
}

type UpdateContactInput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	JobTitle    string `json:"job_title"`
	Company     string `json:"company"`
	City        string `json:"City"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	Country     string `json:"country"`
	Tags        string `json:"tags"`
}
