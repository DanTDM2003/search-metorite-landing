package usecase

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/DanTDM2003/search-api-docker-redis/internal/session"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users"
	userUC "github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

func (uc impleUsecase) SignIn(ctx context.Context, input session.SignInInput) (session.SignInOutput, error) {
	userService := uc.locator.GetService(serviceLocator.UserService).(users.Usecase)
	user, err := userService.GetOneUser(ctx, users.GetOneUserInput{
		Email: input.Email,
	})
	if err != nil {
		if errors.Is(err, userUC.ErrUserNotFound) {
			uc.l.Warnf(ctx, "users.usecase.SignIn.repo.GetOneUserByEmail: %v", err)
			return session.SignInOutput{}, err
		}
		uc.l.Errorf(ctx, "users.usecase.SignIn.repo.GetOneUserByEmail: %v", err)
		return session.SignInOutput{}, err
	}

	if ok := uc.passwordManager.CheckPasswordHash(input.Password, user.Password); !ok {
		uc.l.Warnf(ctx, "users.usecase.SignIn.user.ComparePassword: %v", err)
		return session.SignInOutput{}, userUC.ErrWrongPassword
	}

	refreshToken, err := uc.jwtManager.GetRefreshToken(user.ID)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			refreshToken, err = uc.jwtManager.GenerateRefreshToken(user.ID)
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.GenerateRefreshToken: %v", err)
				return session.SignInOutput{}, err
			}

			err = uc.jwtManager.StoreRefreshToken(user.ID, refreshToken)
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.StoreRefreshToken: %v", err)
				return session.SignInOutput{}, err
			}
		} else {
			uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.GetRefreshToken: %v", err)
			return session.SignInOutput{}, err
		}
	}

	accessToken, err := uc.jwtManager.GetAccessToken(user.ID)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			accessToken, err = uc.jwtManager.GenerateAccessToken(pkgJWT.Payload{
				StandardClaims: jwt.StandardClaims{
					Subject:   strconv.Itoa(int(user.ID)),
					ExpiresAt: jwt.TimeFunc().Add(15 * time.Minute).Unix(),
				},
				Role: user.Role,
			})
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.GenerateAccessToken: %v", err)
				return session.SignInOutput{}, err
			}

			err = uc.jwtManager.StoreAccessToken(user.ID, accessToken)
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.StoreAccessToken: %v", err)
				return session.SignInOutput{}, err
			}
		} else {
			uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.GetAccessToken: %v", err)
			return session.SignInOutput{}, err
		}
	}

	return session.SignInOutput{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc impleUsecase) SignUp(ctx context.Context, input session.SignUpInput) (session.SignUpOutput, error) {
	userService := uc.locator.GetService(serviceLocator.UserService).(users.Usecase)
	_, err := userService.GetOneUser(ctx, users.GetOneUserInput{
		Email: input.Email,
	})
	if err != nil {
		if errors.Is(err, userUC.ErrUserNotFound) {
			uc.l.Warnf(ctx, "users.usecase.SignUp.repo.GetOneUserByEmail: %v", err)
			return session.SignUpOutput{}, err
		}
		uc.l.Errorf(ctx, "users.usecase.SignUp.repo.GetOneUserByEmail: %v", err)
		return session.SignUpOutput{}, err
	}

	user, err := userService.CreateUser(ctx, users.CreateUserInput{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignUp.repo.CreateUser: %v", err)
		return session.SignUpOutput{}, err
	}

	refreshToken, err := uc.jwtManager.GenerateRefreshToken(user.ID)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignUp.jwtManager.GenerateRefreshToken: %v", err)
		return session.SignUpOutput{}, err
	}

	err = uc.jwtManager.StoreRefreshToken(user.ID, refreshToken)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignUp.jwtManager.StoreRefreshToken: %v", err)
		return session.SignUpOutput{}, err
	}

	accessToken, err := uc.jwtManager.GenerateAccessToken(pkgJWT.Payload{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.TimeFunc().Add(15 * time.Minute).Unix(),
		},
		Role: user.Role,
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignUp.jwtManager.GenerateAccessToken: %v", err)
		return session.SignUpOutput{}, err
	}

	err = uc.jwtManager.StoreAccessToken(user.ID, accessToken)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignUp.jwtManager.StoreAccessToken: %v", err)
		return session.SignUpOutput{}, err
	}

	return session.SignUpOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

func (uc impleUsecase) SignOut(ctx context.Context, userID uint) error {
	err := uc.jwtManager.RevokeAccessToken(userID)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignOut.jwtManager.RevokeAccessToken: %v", err)
		return err
	}

	err = uc.jwtManager.RevokeRefreshToken(userID)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignOut.jwtManager.RevokeRefreshToken: %v", err)
		return err
	}

	return nil
}

func (uc impleUsecase) Refresh(ctx context.Context, input session.RefreshInput) (session.RefreshOutput, error) {
	userID, err := uc.jwtManager.VerifyRefreshToken(input.RefreshToken)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.VerifyRefreshToken: %v", err)
		return session.RefreshOutput{}, err
	}

	userService := uc.locator.GetService("userUsecase").(users.Usecase)
	user, err := userService.GetOneUser(ctx, users.GetOneUserInput{
		ID: userID,
	})
	if err != nil {
		if errors.Is(err, userUC.ErrUserNotFound) {
			uc.l.Warnf(ctx, "users.usecase.Refresh.userUC.GetOneUser: %v", err)
			return session.RefreshOutput{}, err
		}
		uc.l.Errorf(ctx, "users.usecase.Refresh.userUC.GetOneUser: %v", err)
		return session.RefreshOutput{}, err
	}

	refreshToken, err := uc.jwtManager.GetRefreshToken(user.ID)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			refreshToken, err = uc.jwtManager.GenerateRefreshToken(user.ID)
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.GenerateRefreshToken: %v", err)
				return session.RefreshOutput{}, err
			}

			err = uc.jwtManager.StoreRefreshToken(user.ID, refreshToken)
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.StoreRefreshToken: %v", err)
				return session.RefreshOutput{}, err
			}
		} else {
			uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.GetRefreshToken: %v", err)
			return session.RefreshOutput{}, err
		}
	}

	accessToken, err := uc.jwtManager.GenerateAccessToken(pkgJWT.Payload{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.TimeFunc().Add(15 * time.Minute).Unix(),
		},
		Role: user.Role,
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.GenerateAccessToken: %v", err)
		return session.RefreshOutput{}, err
	}

	err = uc.jwtManager.StoreAccessToken(user.ID, accessToken)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.StoreAccessToken: %v", err)
		return session.RefreshOutput{}, err
	}

	return session.RefreshOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
