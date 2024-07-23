package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/xiaorui/go-zero-pract1/rpc/internal/svc"
	"github.com/xiaorui/go-zero-pract1/rpc/user"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义方法
func (l *GetUserLogic) GetUser(req *user.GetUserReq) (*user.GetUserResp, error) {
	if u, ok := users[req.Id]; ok {
		return &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}, nil
	}

	return &user.GetUserResp{}, nil
}
