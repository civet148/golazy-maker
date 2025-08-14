package pay

import (
	"context"

	"main/internal/svc"
	"main/internal/types"
)

type WechatPayNotifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 微信支付回调通知(URI包含交易ID)
func NewWechatPayNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatPayNotifyLogic {
	return &WechatPayNotifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatPayNotifyLogic) WechatPayNotify(ctx context.Context, req *types.WechatPayNotifyReq) error {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return nil
}
