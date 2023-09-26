package user

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

// GetUserByEmail is the firebase function call to fbAuth.GetUserByEmail
func GetUserByEmail(context context.Context, email string) (*auth.UserRecord, error) {
	user, err := fbAuth.GetUserByEmail(context, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
