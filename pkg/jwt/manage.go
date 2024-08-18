package jwt

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	pkgRedis "github.com/DanTDM2003/search-api-docker-redis/pkg/redis"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

type Manager interface {
	GenerateAccessToken(payload Payload) (string, error)
	VerifyAccessToken(token string) (Payload, error)
	StoreAccessToken(userID uint, token string) error
	GetAccessToken(userID uint) (string, error)
	RevokeAccessToken(userID uint) error

	VerifyRefreshToken(token string) (uint, error)
	GenerateRefreshToken(userID uint) (string, error)
	StoreRefreshToken(userID uint, token string) error
	GetRefreshToken(userID uint) (string, error)
	RevokeRefreshToken(userID uint) error
}

type Payload struct {
	jwt.StandardClaims
	Role    string `json:"role"`
	Refresh bool   `json:"refresh"`
}

type impleManager struct {
	secretKey string
	redis     *pkgRedis.RedisClient
}

func New(
	secretKey string,
	redis *pkgRedis.RedisClient,
) Manager {
	return &impleManager{
		secretKey: secretKey,
		redis:     redis,
	}
}

func (manager impleManager) GenerateAccessToken(payload Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager impleManager) VerifyAccessToken(token string) (Payload, error) {
	ctx := context.Background()

	payload := Payload{}
	jwtToken, err := jwt.ParseWithClaims(token, &payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secretKey), nil
	})
	if err != nil {
		log.Printf("failed to parse access token: %v", err)
		return Payload{}, err
	}

	if !jwtToken.Valid {
		log.Printf("invalid access token: %v", ErrInvalidToken)
		return Payload{}, ErrInvalidToken
	}

	userID, err := strconv.Atoi(payload.Subject)
	if err != nil {
		log.Printf("failed to convert subject to uint: %v", err)
		return Payload{}, ErrInvalidToken
	}

	storedToken, err := manager.redis.Get(ctx, fmt.Sprintf("access_token:%d", uint(userID)))
	if err == redis.Nil {
		log.Printf("access token not found in redis: %v", ErrInvalidToken)
		return Payload{}, ErrInvalidToken
	} else if err != nil {
		log.Printf("failed to check access token in redis: %v", err)
		return Payload{}, err
	}

	if token != storedToken {
		log.Printf("access token not found in redis: %v", ErrInvalidToken)
		return Payload{}, ErrInvalidToken
	}

	return payload, nil
}

func (manager impleManager) StoreAccessToken(userID uint, token string) error {
	ctx := context.Background()

	err := manager.redis.Set(ctx, fmt.Sprintf("access_token:%d", userID), token)
	if err != nil {
		log.Printf("failed to set access token: %v", err)
		return err
	}

	err = manager.redis.Expire(ctx, fmt.Sprintf("access_token:%d", userID), time.Second*time.Duration(pkgRedis.ONE_MINUTES*15))
	if err != nil {
		log.Printf("failed to set expiration time for access token: %v", err)
		return err
	}

	return nil
}

func (manager impleManager) GetAccessToken(userID uint) (string, error) {
	ctx := context.Background()

	token, err := manager.redis.Get(ctx, fmt.Sprintf("access_token:%d", userID))
	if err != nil {
		log.Printf("failed to get access token: %v", err)
		return "", err
	}

	return token, nil
}

func (manager impleManager) RevokeAccessToken(userID uint) error {
	ctx := context.Background()

	err := manager.redis.Del(ctx, fmt.Sprintf("access_token:%d", userID))
	if err != nil {
		log.Printf("failed to delete access token: %v", err)
		return err
	}

	return nil
}

func (manager impleManager) GenerateRefreshToken(userID uint) (string, error) {
	payload := Payload{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(int(userID)),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
		Refresh: true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager impleManager) VerifyRefreshToken(token string) (uint, error) {
	// Step 1: Check if the token exists in Redis
	ctx := context.Background()

	// Step 2: Verify the token
	payload := Payload{}
	jwtToken, err := jwt.ParseWithClaims(token, &payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secretKey), nil
	})
	if err != nil {
		log.Printf("failed to parse refresh token: %v", err)
		return 0, err
	}

	if !jwtToken.Valid || !payload.Refresh {
		log.Printf("invalid refresh token: %v", ErrInvalidToken)
		return 0, ErrInvalidToken
	}

	userID, err := strconv.Atoi(payload.Subject)
	if err != nil {
		log.Printf("failed to convert subject to int: %v", err)
		return 0, err
	}

	storedToken, err := manager.redis.Get(ctx, fmt.Sprintf("refresh_token:%d", uint(userID)))
	if err == redis.Nil {
		log.Printf("refresh token not found in redis: %v", ErrInvalidToken)
		return 0, ErrInvalidToken
	} else if err != nil {
		log.Printf("failed to check refresh token in redis: %v", err)
		return 0, err
	}

	if token != storedToken {
		log.Printf("refresh token not found in redis: %v", ErrInvalidToken)
		return 0, ErrInvalidToken
	}

	return uint(userID), nil
}

func (manager impleManager) StoreRefreshToken(userID uint, token string) error {
	ctx := context.Background()

	err := manager.redis.Set(ctx, fmt.Sprintf("refresh_token:%d", userID), token)
	if err != nil {
		log.Printf("failed to set refresh token: %v", err)
		return err
	}

	err = manager.redis.Expire(ctx, fmt.Sprintf("refresh_token:%d", userID), time.Second*time.Duration(pkgRedis.ONE_DAY*30))
	if err != nil {
		log.Printf("failed to set expiration time for refresh token: %v", err)
		return err
	}

	return nil
}

func (manager impleManager) GetRefreshToken(userID uint) (string, error) {
	ctx := context.Background()

	token, err := manager.redis.Get(ctx, fmt.Sprintf("refresh_token:%d", userID))
	if err != nil {
		log.Printf("failed to get refresh token: %v", err)
		return "", err
	}

	return token, nil
}

func (manager impleManager) RevokeRefreshToken(userID uint) error {
	ctx := context.Background()

	err := manager.redis.Del(ctx, fmt.Sprintf("refresh_token:%d", userID))
	if err != nil {
		log.Printf("failed to delete refresh token: %v", err)
		return err
	}

	return nil
}
