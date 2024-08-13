package jwt

import (
	"log"
	"strconv"
	"time"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Manager interface {
	GenerateAccessToken(payload Payload) (string, error)
	GenerateRefreshToken(userID int) (string, error)
	VerifyAccessToken(token string) (Payload, error)
	VerifyRefreshToken(token string) (int, error)
	StoreRefreshToken(userID uint, token string, expiresAt time.Time) error
}

type Payload struct {
	jwt.StandardClaims
	Role    string `json:"role"`
	Refresh bool   `json:"refresh"`
}

type impleManager struct {
	secretKey string
	db        *gorm.DB
}

func New(
	secretKey string,
	db *gorm.DB,
) Manager {
	return &impleManager{
		secretKey: secretKey,
		db:        db,
	}
}

func (manager impleManager) GenerateAccessToken(payload Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager impleManager) GenerateRefreshToken(userID int) (string, error) {
	payload := Payload{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(userID),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
		Refresh: true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager impleManager) VerifyAccessToken(token string) (Payload, error) {
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

	return payload, nil
}

func (manager impleManager) VerifyRefreshToken(token string) (int, error) {
	var claims jwt.StandardClaims
	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secretKey), nil
	})
	if err != nil {
		log.Printf("failed to parse refresh token: %v", err)
		return 0, err
	}

	if !jwtToken.Valid {
		log.Printf("invalid refresh token: %v", ErrInvalidToken)
		return 0, ErrInvalidToken
	}

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		log.Printf("failed to convert subject to int: %v", err)
		return 0, err
	}

	return userID, nil
}

func (manager impleManager) StoreRefreshToken(userID uint, token string, expiresAt time.Time) error {
	return manager.db.Create(&models.RefreshToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	}).Error
}
