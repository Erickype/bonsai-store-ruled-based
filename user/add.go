package user

import (
	"context"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

// CreateUser creates a user based on firebase token to control
// admin acces due to firebase cloud functions free tier restrictions
//
//encore:api auth method=POST path=/user
func CreateUser(context context.Context) error {
	uuid, exists := auth.UserID()
	if !exists {
		return &errs.Error{
			Code:    errs.NotFound,
			Message: "User not found",
		}
	}
	query := `insert into "user" (uuid, isAdmin)
			  values($1, false)`
	_, err := sqldb.Exec(context, query, uuid)
	if err != nil {
		return err
	}
	return nil
}
