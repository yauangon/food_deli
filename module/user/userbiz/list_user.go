package userbiz

import (
	"context"

	"food_deli/common"

	"food_deli/module/user/usermodel"
)

type ListUserStore interface {
	ListUsers(ctx context.Context, paging *common.Paging) ([]usermodel.User, *common.AppError)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListUsers(ctx context.Context, paging *common.Paging) ([]usermodel.User, *common.AppError) {
	data, err := biz.store.ListUsers(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}
	return data, nil
}
