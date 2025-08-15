package user

import (
	"github.com/gin-gonic/gin"
	"main/internal/types"
	"net/http"

	"github.com/civet148/log"
	"main/internal/logic/api/v1/user"
	"main/internal/svc"
)

// @Summary 添加用户
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param AddUser body types.AddUserReq true "params description"
// @Success 200 {object} types.AddUserRsp
// @Router /api/v1/user/add [put]
func AddUserHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.AddUserReq
		if err := c.ShouldBind(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := user.NewAddUserLogic(c, svcCtx)

		resp, err := l.AddUser(c, &req)
		if err != nil {
			log.Errorf("call AddUser failed, err: %v", err.Error())
		}
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
