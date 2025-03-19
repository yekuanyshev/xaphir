package form

import (
	"errors"
	"regexp"

	"github.com/charmbracelet/huh"
)

const (
	usernameRegex = `^[A-Za-z0-9_]{6,32}$` //nolint:gosec
	passwordRegex = `^[A-Za-z0-9_]{6,32}$` //nolint:gosec
)

var (
	ErrInvalidUsername  = errors.New("invalid username")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrPasswordMatching = errors.New("passwords doesn't match")
)

type SignUp struct {
	username                string
	password                string
	usernameValidationRegex *regexp.Regexp
	passwordValidationRegex *regexp.Regexp
}

func NewSignUp() *SignUp {
	return &SignUp{
		usernameValidationRegex: regexp.MustCompile(usernameRegex),
		passwordValidationRegex: regexp.MustCompile(passwordRegex),
	}
}

func (su *SignUp) Run() error {
	inputs := []*huh.Input{
		su.usernameInput(),
		su.passwordInput(),
		su.retypedPasswordInput(),
	}

	for _, input := range inputs {
		group := huh.NewGroup(input)
		form := huh.NewForm(group)
		err := form.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func (su *SignUp) Username() string {
	return su.username
}

func (su *SignUp) Password() string {
	return su.password
}

func (su *SignUp) usernameInput() *huh.Input {
	return huh.NewInput().
		Title("Username").
		Validate(su.validateUsername).
		Value(&su.username)
}

func (su *SignUp) passwordInput() *huh.Input {
	return huh.NewInput().
		Title("Password").
		EchoMode(huh.EchoModePassword).
		Validate(su.validatePassword).
		Value(&su.password)
}

func (su *SignUp) retypedPasswordInput() *huh.Input {
	return huh.NewInput().
		Title("Retype password").
		EchoMode(huh.EchoModePassword).
		Validate(su.validateRetypedPassword)
}

func (su *SignUp) validateUsername(username string) error {
	match := su.usernameValidationRegex.MatchString(username)
	if !match {
		return ErrInvalidUsername
	}
	return nil
}

func (su *SignUp) validatePassword(password string) error {
	match := su.passwordValidationRegex.MatchString(password)
	if !match {
		return ErrInvalidPassword
	}

	return nil
}

func (su *SignUp) validateRetypedPassword(retypedPassword string) error {
	if su.password != retypedPassword {
		return ErrPasswordMatching
	}

	return nil
}
