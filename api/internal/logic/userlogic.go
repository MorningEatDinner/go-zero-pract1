package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/xiaorui/go-zero-pract1/api/internal/svc"
	"github.com/xiaorui/go-zero-pract1/api/internal/types"
	"github.com/xiaorui/go-zero-pract1/rpc/userclient"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserResp, err error) {
	getUserResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.GetUserReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}
	resp = &types.UserResp{
		Id:    getUserResp.Id,
		Name:  getUserResp.Name,
		Phone: getUserResp.Phone,
	}

	return
}
