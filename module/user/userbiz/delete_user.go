package userbiz

import (
	"errors"

	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

// DeleteUserStore delete action biz
type DeleteUserStore interface {
	FindDataWithCondition(condition map[string]interface{}) (*usermodel.User, error)
	Delete(id int) error
}

type deleteUserBiz struct {
	store DeleteUserStore
}

// NewDeleteUserBiz biz main action
func NewDeleteUserBiz(store DeleteUserStore) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUser(userID int) error {
	user, err := biz.store.FindDataWithCondition(map[string]interface{}{"id": userID})
	if err != nil {
		return err
	}
	if user.Status == 0 {
		return errors.New("User has been deleted before")
	}
	if err := biz.store.Delete(userID); err != nil {
		return err
	}
	return nil
}
