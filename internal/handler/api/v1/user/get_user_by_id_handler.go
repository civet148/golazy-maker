package user

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"main/internal/logic/api/v1/user"
	"main/internal/svc"
	"main/internal/types"
	"net/http"
)

// @Summary 根据ID查询用户
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param GetUserById body types.GetUserByIdReq true "params description"
// @Success 200 {object} types.GetUserByIdRsp
// @Router /api/v1/user/:id [get]
func GetUserByIdHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req types.GetUserByIdReq
		if err := c.ShouldBindUri(&req); err != nil {
			if err != nil {
				log.Errorf("call ShouldBind/ShouldBindUri failed, err: %v", err.Error())
			}
			c.JSON(http.StatusOK, svc.JsonResponse(nil, err))
			return
		}
		log.Infof("request: %+v", req)

		l := user.NewGetUserByIdLogic(c, svcCtx)

		resp, err := l.GetUserById(c, &req)
		if err != nil {
			log.Errorf("call GetUserById failed, err: %v", err.Error())
		}
		c.JSON(http.StatusOK, svc.JsonResponse(resp, err))

	}
}
