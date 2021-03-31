package userbiz

import (
	"context"

	"food_deli/common"

	"food_deli/module/user/usermodel"
)

type GetUserStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*usermodel.User, *common.AppError)
}

type getUserBiz struct {
	store GetUserStore
}

func NewGetUserBiz(store GetUserStore) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz getUserBiz) GetUserByID(ctx context.Context, id int) (*usermodel.User, *common.AppError) {
	user, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	return user, nil
}
