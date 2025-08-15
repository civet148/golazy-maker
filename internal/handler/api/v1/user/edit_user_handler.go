package user

import (
	"github.com/gin-gonic/gin"
	"main/internal/types"
	"net/http"

	"github.com/civet148/log"
	"main/internal/logic/api/v1/user"
	"main/internal/svc"
)

// @Summary 修改用户
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param EditUser body types.EditUserReq true "params description"
// @Success 200 {object} types.EditUserRsp
// @Router /api/v1/user/edit [post]
func EditUserHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.EditUserReq
		if err := c.ShouldBind(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := user.NewEditUserLogic(c, svcCtx)

		resp, err := l.EditUser(c, &req)
		if err != nil {
			log.Errorf("call EditUser failed, err: %v", err.Error())
		}
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
