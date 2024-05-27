package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type User struct {
	Uuid         uuid.UUID
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	Role         string `json:"role"`
	Token        string
	RefreshToken string
	User_Id      string
	Created_at   time.Time
	Updated_at   time.Time
}
type SignedTokens struct {
	User_ID   string
	FirstName string
	Email     string
	Password  string
	jwt.StandardClaims
}
