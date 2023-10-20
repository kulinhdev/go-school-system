package helpers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var jwtKey = []byte("LVMz9Ze4vG8hC5naqihFLgdAqSBaCHiaM3jVpdLJvY2xWc1PkB50kAmCgAHcLewiR6cBSdC6wpLYx9bSj2Dnz3j5UwP1eb9ESD79kzwawxeTNKj78rgbnbPA")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %w", err)
	}
	return string(bytes), nil
}

func CheckPassword(providedPassword string, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func CreateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func ValidateToken(signedToken string, publicKey string) (sub interface{}, err error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode public key: %w", err)
	}

	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	})

	if err != nil {
		return nil, fmt.Errorf("validate token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	sub = claims["sub"]

	return sub, nil
}
