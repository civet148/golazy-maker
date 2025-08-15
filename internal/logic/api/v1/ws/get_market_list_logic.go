package ws

import (
	"context"

	"main/internal/svc"
)

type GetMarketListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 市场行情（websocket方式）
func NewGetMarketListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMarketListLogic {
	return &GetMarketListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMarketListLogic) GetMarketList(ctx context.Context) error {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return nil
}
