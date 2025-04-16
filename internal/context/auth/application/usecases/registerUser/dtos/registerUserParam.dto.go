package dtos

type RegisterUserParam struct {
	ID       int
	Name     string
	Password string
	Email    string
	Provider string
	Img      string
}
