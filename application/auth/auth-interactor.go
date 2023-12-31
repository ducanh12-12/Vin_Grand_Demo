package auth

import (
	"base-go/application/requires"
	"base-go/domain/model"
	"context"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthInteractor struct {
	userService requires.UsersService
}

func NewAuthInteractor(userService requires.UsersService) AuthInteractor {
	return AuthInteractor{userService}
}
func Create(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() //Token hết hạn sau 12 giờ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func (interactor *AuthInteractor) Login(ctx context.Context, username string, phone_number string, password string) (string, error) {
	user := &model.User{}
	var err error
	if username != "" {
		user, err = interactor.userService.GetUserByUserName(ctx, username)
		if err != nil {
			return "", fmt.Errorf("Incorrect username or password")
		}
	}
	if phone_number != "" {
		user, err = interactor.userService.GetUserByPhoneNumber(ctx, phone_number)
		if err != nil {
			return "", fmt.Errorf("Incorrect phone number or password")
		}
	}
	hashedPassword := user.Password
	err = CheckPasswordHash(*hashedPassword, password)

	if err != nil {
		return "", fmt.Errorf("Incorrect password")
	}
	token, errCreate := Create(username)
	if errCreate != nil {
		return "", err
	}
	return token, nil
}

// func (interactor *UsersInteractor) GetUser(ctx context.Context, id int) (*GetUserResp, error) {
// 	return interactor.userService.Retrieve(ctx, id)
// }
// func (interactor *UsersInteractor) GetUserByUserName(ctx context.Context, username string) (*GetUserResp, error) {
// 	return interactor.userService.GetUserByUserName(ctx, username)
// }
// func (interactor *UsersInteractor) GetUsers(ctx context.Context) (*GetUserResp, error) {
// 	return interactor.userService.GetUsers(ctx)
// }
