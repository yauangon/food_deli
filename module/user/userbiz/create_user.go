package userbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

type CreateUserStore interface {
	Create(ctx context.Context, user *usermodel.User) error
}

type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) Create(ctx context.Context, user *usermodel.User) error {
	if err := biz.store.Create(ctx, user); err != nil {
		return err
	}
	return nil
}
