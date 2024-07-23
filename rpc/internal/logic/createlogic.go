package logic

import (
	"context"
	"github.com/xiaorui/go-zero-pract1/models"
	"github.com/xiaorui/go-zero-pract1/rpc/internal/svc"
	"github.com/xiaorui/go-zero-pract1/rpc/user"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateReq) (*user.CreateResp, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &models.Users{
		Id:    in.Id,
		Name:  in.Name,
		Phone: in.Phone,
	})
	if err != nil {
		log.Println(" l.svcCtx.UserModel.Insert err", err.Error())
	}
	log.Println("res:", res)
	return &user.CreateResp{}, nil
}
