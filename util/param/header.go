/**
 *@Description
 *@ClassName header
 *@Date 2020/12/28 下午10:13
 *@Author ckhero
 */

package param

import (
	"github.com/gin-gonic/gin"
	"youmi-micro-cluster/src/common/constant"
)

func GetToken(c *gin.Context) string {

	token := c.GetHeader(constant.HeadAuthorization)
	if len(token) == 0 {
		token = c.Query(constant.HeadAuthorization)
	}
	return token
}

