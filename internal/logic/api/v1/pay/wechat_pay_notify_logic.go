package pay

import (
	"context"

	"github.com/gin-gonic/gin"

	"main/internal/svc"
	"main/internal/types"
)

type WechatPayNotifyLogic struct {
	ginCtx *gin.Context
	svcCtx *svc.ServiceContext
}

// 微信支付回调通知(URI包含交易ID)
func NewWechatPayNotifyLogic(c *gin.Context, svcCtx *svc.ServiceContext) *WechatPayNotifyLogic {
	return &WechatPayNotifyLogic{
		ginCtx: c,
		svcCtx: svcCtx,
	}
}

func (l *WechatPayNotifyLogic) WechatPayNotify(ctx context.Context, req *types.WechatPayNotifyReq) error {
	// you can call ctx.(*gin.Context) convert to gin context
	// todo: add your logic here and delete this line

	return nil
}
