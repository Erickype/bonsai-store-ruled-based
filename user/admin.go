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
	user, err := fbAuth.GetUserByEmail(context, email)
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
