package domain

type User struct {
	ID          int    `json:"id"`
	PhoneNumber int    `json:"phone_number"`
	Password    string `json:"password"`
}

type SignInInput struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
