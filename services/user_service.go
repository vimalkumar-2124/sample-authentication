package services

import (
	"errors"
	"strings"
	"time"

	"github.com/vimalkumar-2124/sample-authentication/models"
	"github.com/vimalkumar-2124/sample-authentication/repositories"
	"github.com/vimalkumar-2124/sample-authentication/tokens"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepo
}

type UserServiceRoutes interface {
	SignUp(body models.SignUpBody) (string, error)
	SignIn(body models.SignInBody) (string, error)
	GetEncryptedPassword(password string) string    // Password encrypt
	CompareHashPassword(password, hash string) bool // Compare encrypt password while signin
	LogOut(authToken string) error
	ChangePassword(body models.ChangeUserPassword) error
}

func NewInstanceOfUserService(userRepo repositories.UserRepo) UserService {
	return UserService{userRepo: userRepo}
}

func (u *UserService) GetEncryptedPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes)
}

func (u *UserService) CompareHashPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func (u *UserService) signIn(email string, password string) (string, error) {

	// Grab user
	found, user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if !found {
		return "", errors.New("Unauthorized")
	}
	match := u.CompareHashPassword(user.Password, password)
	if !match {
		return "", errors.New("error : Password not matched")
	}

	token, err := tokens.GenerateJWT(email)
	if err != nil {
		return "", err
	}

	// Create Session
	now := time.Now()
	expiryDate := now.AddDate(0, 0, 1)
	newSession := models.Session{
		Email:       email,
		Created:     now,
		Expiry:      expiryDate,
		TokenString: token,
	}

	err = u.userRepo.SaveSession(newSession)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserService) SignIn(body models.SignInBody) (string, error) {
	emailLowerCase := strings.ToLower(body.Email)
	emailTrim := strings.Trim(emailLowerCase, " ")
	return u.signIn(emailTrim, body.Password)
}

func (u *UserService) SignUp(body models.SignUpBody) (string, error) {
	emailLowerCase := strings.ToLower(body.Email)
	emailTrim := strings.Trim(emailLowerCase, " ")
	password := u.GetEncryptedPassword(body.Password)

	// Sign Up user
	newUser := models.Users{
		Email:    emailTrim,
		Password: password,
		Name:     body.Name,
		Role:     body.Role,
		Mobile:   body.Mobile,
		Created:  time.Now(),
	}
	err := u.userRepo.SaveUser(newUser)
	if err != nil {
		return "", err
	}

	// Sign in user
	// return u.signIn(emailTrim, body.Password)
	return "User signed up successfully!!", nil
}

func (u *UserService) LogOut(authToken string) error {
	// Grab session
	found, _, err := u.userRepo.GetSessinById(authToken)
	if err != nil {
		return err
	}
	if !found {
		return errors.New("error : Unauthorized")
	}

	// Mark as expired

	err = u.userRepo.MarkSessionAsExpired(authToken)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ChangePassword(body models.ChangeUserPassword) error {
	found, user, err := u.userRepo.GetUserByEmail(body.Email)
	if err != nil {
		return err
	}
	if found {
		match := u.CompareHashPassword(user.Password, body.Old_Password)
		if match {
			new_pass := u.GetEncryptedPassword(body.New_Password)
			updatedUserPassword := models.SignInBody{
				Email:    user.Email,
				Password: new_pass,
			}
			err = u.userRepo.UpdateUser(updatedUserPassword)
			if err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("Password is not matched")
		}

	} else {
		return errors.New("User not found!!!")
	}
	// return nil
}
