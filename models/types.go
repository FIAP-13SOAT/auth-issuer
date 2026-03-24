package models

type Input struct {
	Document string `json:"document"`
}

type AdminInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Output struct {
	Token string `json:"token"`
}
