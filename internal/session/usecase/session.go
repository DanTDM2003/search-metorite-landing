package usecase

import (
	"context"
	"errors"
	"strconv"
	"time"

	userSrv "github.com/DanTDM2003/search-api-docker-redis/internal/application/user"
	userUC "github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

func (uc impleUsecase) SignIn(ctx context.Context, input SignInInput) (SignInOutput, error) {
	userService := uc.locator.GetService(serviceLocator.UserService).(userSrv.UserUsecase)
	user, err := userService.GetOneUser(ctx, userSrv.GetOneUserInput{
		Email: input.Email,
	})
	if err != nil {
		if errors.Is(err, userUC.ErrUserNotFound) {
			uc.l.Warnf(ctx, "users.usecase.SignIn.repo.GetOneUserByEmail: %v", err)
			return SignInOutput{}, err
		}
		uc.l.Errorf(ctx, "users.usecase.SignIn.repo.GetOneUserByEmail: %v", err)
		return SignInOutput{}, err
	}

	if ok := uc.passwordManager.CheckPasswordHash(input.Password, user.Password); !ok {
		uc.l.Warnf(ctx, "users.usecase.SignIn.user.ComparePassword: %v", err)
		return SignInOutput{}, userUC.ErrWrongPassword
	}

	refreshToken, err := uc.jwtManager.GetRefreshToken(user.ID)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			refreshToken, err = uc.jwtManager.GenerateRefreshToken(user.ID)
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.GenerateRefreshToken: %v", err)
				return SignInOutput{}, err
			}

			err = uc.jwtManager.StoreRefreshToken(user.ID, refreshToken)
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.StoreRefreshToken: %v", err)
				return SignInOutput{}, err
			}
		} else {
			uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.GetRefreshToken: %v", err)
			return SignInOutput{}, err
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
				return SignInOutput{}, err
			}

			err = uc.jwtManager.StoreAccessToken(user.ID, accessToken)
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.StoreAccessToken: %v", err)
				return SignInOutput{}, err
			}
		} else {
			uc.l.Errorf(ctx, "users.usecase.SignIn.jwtManager.GetAccessToken: %v", err)
			return SignInOutput{}, err
		}
	}

	return SignInOutput{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc impleUsecase) SignUp(ctx context.Context, input SignUpInput) (SignUpOutput, error) {
	userService := uc.locator.GetService(serviceLocator.UserService).(userSrv.UserUsecase)
	_, err := userService.GetOneUser(ctx, userSrv.GetOneUserInput{
		Email: input.Email,
	})
	if err != nil {
		if errors.Is(err, userUC.ErrUserNotFound) {
			uc.l.Warnf(ctx, "users.usecase.SignUp.repo.GetOneUserByEmail: %v", err)
			return SignUpOutput{}, err
		}
		uc.l.Errorf(ctx, "users.usecase.SignUp.repo.GetOneUserByEmail: %v", err)
		return SignUpOutput{}, err
	}

	user, err := userService.CreateUser(ctx, userSrv.CreateUserInput{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignUp.repo.CreateUser: %v", err)
		return SignUpOutput{}, err
	}

	refreshToken, err := uc.jwtManager.GenerateRefreshToken(user.ID)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignUp.jwtManager.GenerateRefreshToken: %v", err)
		return SignUpOutput{}, err
	}

	err = uc.jwtManager.StoreRefreshToken(user.ID, refreshToken)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignUp.jwtManager.StoreRefreshToken: %v", err)
		return SignUpOutput{}, err
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
		return SignUpOutput{}, err
	}

	err = uc.jwtManager.StoreAccessToken(user.ID, accessToken)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignUp.jwtManager.StoreAccessToken: %v", err)
		return SignUpOutput{}, err
	}

	return SignUpOutput{
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

func (uc impleUsecase) Refresh(ctx context.Context, input RefreshInput) (RefreshOutput, error) {
	userID, err := uc.jwtManager.VerifyRefreshToken(input.RefreshToken)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.VerifyRefreshToken: %v", err)
		return RefreshOutput{}, err
	}

	userService := uc.locator.GetService("userUsecase").(userSrv.UserUsecase)
	user, err := userService.GetOneUser(ctx, userSrv.GetOneUserInput{
		ID: userID,
	})
	if err != nil {
		if errors.Is(err, userUC.ErrUserNotFound) {
			uc.l.Warnf(ctx, "users.usecase.Refresh.userUC.GetOneUser: %v", err)
			return RefreshOutput{}, err
		}
		uc.l.Errorf(ctx, "users.usecase.Refresh.userUC.GetOneUser: %v", err)
		return RefreshOutput{}, err
	}

	refreshToken, err := uc.jwtManager.GetRefreshToken(user.ID)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			refreshToken, err = uc.jwtManager.GenerateRefreshToken(user.ID)
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.GenerateRefreshToken: %v", err)
				return RefreshOutput{}, err
			}

			err = uc.jwtManager.StoreRefreshToken(user.ID, refreshToken)
			if err != nil {
				uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.StoreRefreshToken: %v", err)
				return RefreshOutput{}, err
			}
		} else {
			uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.GetRefreshToken: %v", err)
			return RefreshOutput{}, err
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
		return RefreshOutput{}, err
	}

	err = uc.jwtManager.StoreAccessToken(user.ID, accessToken)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Refresh.jwtManager.StoreAccessToken: %v", err)
		return RefreshOutput{}, err
	}

	return RefreshOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
