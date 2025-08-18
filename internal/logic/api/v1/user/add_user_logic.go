package user

import (
	"context"

	"github.com/gin-gonic/gin"

	"main/internal/svc"
	"main/internal/types"
)

type AddUserLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 添加用户
func NewAddUserLogic(c *gin.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(ctx context.Context, req *types.AddUserReq) (resp *types.AddUserRsp, err error) {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return &types.AddUserRsp{}, nil
}
