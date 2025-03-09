package form

import "github.com/charmbracelet/huh"

type SignIn struct {
	username string
	password string
}

func NewSignIn() *SignIn {
	return &SignIn{}
}

func (si *SignIn) Run() error {
	inputs := []*huh.Input{
		si.usernameInput(),
		si.passwordInput(),
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

func (si *SignIn) Username() string {
	return si.username
}

func (si *SignIn) Password() string {
	return si.password
}

func (si *SignIn) usernameInput() *huh.Input {
	return huh.NewInput().
		Title("Username").
		Value(&si.username)
}

func (si *SignIn) passwordInput() *huh.Input {
	return huh.NewInput().
		Title("Password").
		EchoMode(huh.EchoModePassword).
		Value(&si.password)
}
