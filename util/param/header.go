/**
 *@Description
 *@ClassName header
 *@Date 2020/12/28 下午10:13
 *@Author ckhero
 */

package param

import (
	headerConst "github.com/ckhero/go-common/constant/header"
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) string {

	token := c.GetHeader(headerConst.HeaderToken)
	if len(token) == 0 {
		token = c.Query(headerConst.HeaderToken)
	}
	return token
}

