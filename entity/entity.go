package entity

type Contact struct {
	FirstName   string `json:"firstName" db:"first_name"`
	LastName    string `json:"lastName" db:"last_name"`
	Address     string `json:"address" db:"address"`
	PhoneNumber int    `json:"phoneNumber" db:"phone_number"`
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}
