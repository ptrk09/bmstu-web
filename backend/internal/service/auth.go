package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
	"web/internal/model"
	"web/internal/repository"

	"github.com/dgrijalva/jwt-go/v4"
)

const (
	salt       = "flexomomomomomomomomo"
	tokenTTL   = 12 * time.Hour
	signingKey = "qwqjljljqojlqkjw"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"userId"`
	Role   string `json:"role"`
}

type UserContextData struct {
	UserId int    `json:"userId"`
	Role   string `json:"role"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

// AddUser godoc
// @Summary Adds user to database
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Failure 400 {object} model.User
// @Failure 500 {object} model.User
// @Param request body model.User true "User's email and password and name"
// @Router /users/ [post]
func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.At(time.Now()),
		},
		user.ID,
		user.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (UserContextData, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return UserContextData{0, ""}, err
	}

	claim, ok := token.Claims.(*tokenClaims)
	if !ok {
		return UserContextData{0, ""}, errors.New("token claims are not of type *tokenClaims")
	}

	return UserContextData{claim.UserId, claim.Role}, nil
}

// func generatePasswordHash(password string) string {
// 	var passwordBytes = []byte(password)
// 	passwordBytes = append(passwordBytes, []byte(salt)...)
// 	hash := sha256.Sum256(passwordBytes)
// 	return hex.EncodeToString(hash)
// 	// hash := sha1.New()
// 	// hash.Write([]byte(password))
// 	// return base64.URLEncoding.EncodeToString(hash.Sum([]byte(salt)))
// }

func generatePasswordHash(password string) string {
	passwordBytes := []byte(password)
	sha256Hasher := sha256.New()
	sha256Hasher.Write(passwordBytes)

	hashedPasswordBytes := sha256Hasher.Sum(nil)
	hashedPasswordHex := hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}
