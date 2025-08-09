package v1

import (
	"context"
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/logic/api/v1"
	"test/internal/svc"
	"test/internal/types"
)

// @Summary 用户登录
// @Description 用户登录接口
// @Tags
// @Accept json
// @Produce json
// @Param login body types.UserSignInReq true "登录信息"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} svc.Response
// @Failure 500 {object} svc.Response
// @Router /v1/login [post]
func UserSignInHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.UserSignInReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Debugf("request [%+v]", req)
		l := v1.NewUserSignInLogic(context.Background(), svcCtx)
		resp, err := l.UserSignIn(&req)
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
