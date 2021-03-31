package userbiz

import (
	"context"

	"food_deli/common"

	"food_deli/module/user/usermodel"
)

type CreateUserStore interface {
	Create(ctx context.Context, user *usermodel.User) *common.AppError
}

type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) Create(ctx context.Context, user *usermodel.User) *common.AppError {
	if err := biz.store.Create(ctx, user); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
