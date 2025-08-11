package ws

import (
	"context"
	"github.com/gin-gonic/gin"
	"test/internal/logic/api/v1/ws"
	"test/internal/svc"
)

// @Summary 获取市场列表(WebSocket)
// @Description WebSocket协议接口，可能返回文本或二进制数据
// @Tags WebSocket市场数据
// @Accept json
// @Produce plain
// @Security ApiKeyAuth
// @Param authorization header string true "Bearer token"
// @Param market query string false "市场名称过滤" Enums(BTC,ETH,USDT)
// @Success 200 {string} string "成功返回文本或二进制数据"
// @Failure 400 {string} string "错误信息"
// @Router /api/v1/ws/market [get]
func WsMarketListHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		l := ws.NewWsMarketListLogic(context.Background(), svcCtx)
		_ = l.WsMarketList(c)

	}
}
