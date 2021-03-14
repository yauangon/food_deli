package userbiz

import (
	"context"
	"errors"

	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

// DeleteUserStore delete action biz
type DeleteUserStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*usermodel.User, error)
	Delete(ctx context.Context, id int) error
}

type deleteUserBiz struct {
	store DeleteUserStore
}

// NewDeleteUserBiz biz main action
func NewDeleteUserBiz(store DeleteUserStore) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUser(ctx context.Context, userID int) error {
	user, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": userID})
	if err != nil {
		return err
	}
	if user.Status == 0 {
		return errors.New("User has been deleted before")
	}
	if err := biz.store.Delete(ctx, userID); err != nil {
		return err
	}
	return nil
}
