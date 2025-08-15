package v1

import (
	"github.com/gin-gonic/gin"
	"main/internal/types"
	"net/http"

	"github.com/civet148/log"
	"main/internal/logic/api/v1"
	"main/internal/svc"
)

// @Summary 用户注册
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param UserSignUp body types.UserSignUpReq true "params description"
// @Success 200 {object} types.UserSignUpRsp
// @Router /api/v1/sign_up [post]
func UserSignUpHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.UserSignUpReq
		if err := c.ShouldBind(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := v1.NewUserSignUpLogic(c, svcCtx)

		resp, err := l.UserSignUp(c, &req)
		if err != nil {
			log.Errorf("call UserSignUp failed, err: %v", err.Error())
		}
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
