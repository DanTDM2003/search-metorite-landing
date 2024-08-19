package hasher

type Password interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hashedPassword string) bool
}

type passwordManager struct{}

func New() Password {
	return &passwordManager{}
}
