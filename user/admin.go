package user

import (
	"context"

	"encore.dev/storage/sqldb"
	"encore.dev/types/uuid"
)

//encore:api auth method=POST path=/user/admin/:uuid
func MakeAdmin(context context.Context, uuid uuid.UUID) error {
	err := CheckUser(context, uuid)
	if err != nil {
		return err
	}
	query := `update user
			  set isAdmin = true
			  where uuid = $1`
	_, err = sqldb.Exec(context, query, uuid)
	if err != nil {
		return err
	}
	return nil
}