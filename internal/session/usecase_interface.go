package session

import "context"

type Usecase interface {
	SignIn(ctx context.Context, input SignInInput) (SignInOutput, error)
	SignUp(ctx context.Context, input SignUpInput) (SignUpOutput, error)
	SignOut(ctx context.Context, userID uint) error
	Refresh(ctx context.Context, input RefreshInput) (RefreshOutput, error)
}
