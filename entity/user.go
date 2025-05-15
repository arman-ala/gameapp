package entity

type User struct {
	ID          uint
	PhoneNumber string
	Name        string
	Password    string // Password field always keep the hased version of the real password
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
}
