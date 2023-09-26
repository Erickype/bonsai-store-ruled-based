package user

import (
	"context"

	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

// MakeAdmin asigns the admin role to a user
//
//encore:api auth method=POST path=/user/admin/:email
func MakeAdmin(context context.Context, email string) error {
	user, err := GetUserByEmail(context, email)
	if err != nil {
		return &errs.Error{
			Code:    errs.NotFound,
			Message: "User not found",
		}
	}
	query := `update "user"
			  set isAdmin = true
			  where uuid = $1`
	_, err = sqldb.Exec(context, query, user.UID)
	if err != nil {
		return err
	}
	return nil
}

// CheckAdminResponse is the response of CheckAdmin endpoint
type CheckAdminResponse struct {
	IsAdmin bool `json:"isAdmin"`
}

// CheckAdmin checks if the user is an admin
//
//encore:api auth method=GET path=/user/admin/:email
func CheckAdmin(context context.Context, email string) (*CheckAdminResponse, error) {
	user, err := GetUserByEmail(context, email)
	if err != nil {
		return &CheckAdminResponse{IsAdmin: false}, err
	}
	query := `select admin from "user"
			  where uuid = $1`
	sqldb.Exec(context, query, user.UID)
	return &CheckAdminResponse{IsAdmin: true}, nil
}
