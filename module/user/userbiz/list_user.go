package userbiz

import (
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

type ListUserStore interface {
	ListUsers(paging *common.Paging) ([]usermodel.User, error)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListUsers(paging *common.Paging) ([]usermodel.User, error) {
	data, err := biz.store.ListUsers(paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
