package authorization

import (
	"github.com/gin-gonic/gin"
)

type HeaderParams struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
}

// CheckTokenAuth 检查token完整性、有效性中间件
func CheckTokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerParams := HeaderParams{}
		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			//response.TokenErrorParam(context, consts.JwtTokenMustValid+err.Error())
			return
		}
		// 判断token有效性
		//token := strings.Split(headerParams.Authorization, " ")

	}
}
