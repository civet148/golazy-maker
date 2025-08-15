package v1

import (
	"github.com/gin-gonic/gin"
	"main/internal/types"
	"net/http"

	"github.com/civet148/log"
	"main/internal/logic/api/v1"
	"main/internal/svc"
)

// @Summary 用户登录
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param UserSignIn body types.UserSignInReq true "params description"
// @Success 200 {object} types.UserSignInRsp
// @Router /api/v1/sign_in [post]
func UserSignInHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.UserSignInReq
		if err := c.ShouldBind(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := v1.NewUserSignInLogic(c, svcCtx)

		resp, err := l.UserSignIn(c, &req)
		if err != nil {
			log.Errorf("call UserSignIn failed, err: %v", err.Error())
		}
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
