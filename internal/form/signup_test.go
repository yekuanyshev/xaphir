package form

import (
	"errors"
	"testing"
)

func TestValidateUsername(t *testing.T) {
	testCases := []struct {
		input string
		want  error
	}{
		{input: "", want: ErrInvalidUsername},
		{input: "x", want: ErrInvalidUsername},
		{input: "xxxxx", want: ErrInvalidUsername},
		{input: "11111", want: ErrInvalidUsername},
		{input: "xxxxxx", want: nil},
		{input: "111111", want: nil},
		{input: "username", want: nil},
		{input: "username123", want: nil},
		{input: "username_123", want: nil},
	}

	form := NewSignUp()

	for _, tc := range testCases {
		got := form.validateUsername(tc.input)
		if !errors.Is(got, tc.want) {
			t.Errorf("validateUsername(%q)=%v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		input string
		want  error
	}{
		{input: "", want: ErrInvalidPassword},
		{input: "x", want: ErrInvalidPassword},
		{input: "xxxxx", want: ErrInvalidPassword},
		{input: "11111", want: ErrInvalidPassword},
		{input: "xxxxxx", want: nil},
		{input: "111111", want: nil},
		{input: "password", want: nil},
		{input: "password123", want: nil},
		{input: "password_123", want: nil},
	}

	form := NewSignUp()

	for _, tc := range testCases {
		got := form.validatePassword(tc.input)
		if !errors.Is(got, tc.want) {
			t.Errorf("validatePassword(%q)=%v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestValidateRetypedPassword(t *testing.T) {
	testCases := []struct {
		password        string
		retypedPassword string
		want            error
	}{
		{password: "password", retypedPassword: "wordpass", want: ErrPasswordMatching},
		{password: "password", retypedPassword: "password", want: nil},
	}

	form := NewSignUp()

	for _, tc := range testCases {
		form.password = tc.password
		got := form.validateRetypedPassword(tc.retypedPassword)
		if !errors.Is(got, tc.want) {
			t.Errorf(
				"validateRetypedPassword(%s)=%v, want %v, retyped password: %s",
				tc.retypedPassword, got, tc.want, tc.retypedPassword,
			)
		}
	}
}
