package userbiz

import "github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"

type CreateUserStore interface {
	Create(user *usermodel.User) error
}

type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) Create(user *usermodel.User) error {
	if err := biz.store.Create(user); err != nil {
		return err
	}
	return nil
}
