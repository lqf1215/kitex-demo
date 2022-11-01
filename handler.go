package main

import (
	"context"
	"errors"
	"github.com/lqf1215/kitex-demo/global"
	"github.com/lqf1215/kitex-demo/kitex_gen/service"
	"github.com/lqf1215/kitex-demo/modal"
)

type Sex int8

const (
	SexUnKnown Sex = iota
	SexMan
	SexGirl
)

var SexMap = map[Sex]string{
	SexUnKnown: "未知",
	SexMan:     "男",
	SexGirl:    "女",
}

func SexVal(sex Sex) string {
	v, ok := SexMap[sex]
	if !ok {
		v = SexMap[SexUnKnown]
	}
	return v
}

var _ service.UserService = &UserServiceImpl{}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUserList implements the UserImpl interface.
func (s *UserServiceImpl) GetUserList(ctx context.Context, req *service.GetUserListReq) (resp *service.GetUserListResp, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var list []*service.UserData
	var users []modal.User
	db := global.DB.Model(&modal.User{})
	if req.Username != "" {
		db.Where("username LIKE ?", "%"+req.Username+"%")
	}
	err = db.Limit(int(limit)).Offset(int(offset)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		list = append(list, &service.UserData{
			Id:       int64(user.ID),
			Username: user.Username,
			Sex:      SexVal(Sex(user.Sex)),
		})
	}

	resp = &service.GetUserListResp{Data: list}
	return
}

// GetUserById implements the UserImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, req *service.GetUserByIdReq) (resp *service.UserData, err error) {

	if req.Id == 0 {
		return nil, errors.New("请输入ID")
	}

	var user modal.User
	err = global.DB.Model(&modal.User{}).Where("id=?", req.Id).Limit(1).Find(&user).Error

	if err != nil {
		return nil, err
	}
	resp = &service.UserData{
		Id:       int64(user.ID),
		Username: user.Username,
		Sex:      SexVal(Sex(user.Sex)),
	}
	return
}
