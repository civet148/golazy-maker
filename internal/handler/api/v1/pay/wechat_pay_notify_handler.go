package pay

import (
	"github.com/gin-gonic/gin"
	"main/internal/types"
	"net/http"

	"github.com/civet148/log"
	"main/internal/logic/api/v1/pay"
	"main/internal/svc"
)

// @Summary 微信支付回调通知(URI包含交易ID)
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param WechatPayNotify body types.WechatPayNotifyReq true "params description"
// @Success 200 {object} nil
// @Router /api/v1/pay/wechat/notify/:tid [get]
func WechatPayNotifyHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.WechatPayNotifyReq
		if err := c.ShouldBindUri(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := pay.NewWechatPayNotifyLogic(c, svcCtx)

		err := l.WechatPayNotify(c, &req)
		if err != nil {
			log.Errorf("call WechatPayNotify failed, err: %v", err.Error())
			return
		}

	}
}
