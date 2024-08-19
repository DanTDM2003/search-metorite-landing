package user

type GetOneUserInput struct {
	ID    uint
	Email string
}

type CreateUserInput struct {
	Username string
	Email    string
	Password string
}
