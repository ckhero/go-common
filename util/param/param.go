/**
 *@Description
 *@ClassName param
 *@Date 2021/1/7 下午3:29
 *@Author ckhero
 */

package param

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"youmi-micro-cluster/src/common/errors"
	"youmi-micro-cluster/src/common/util"
)

func GetParamUint64(c *gin.Context, key string) (uint64, error) {
	var res uint64
	resStr := c.Param(key)

	res, err := strconv.ParseUint(resStr, 10, 64)
	if err != nil {
		return 0, errors.ErrorParseParam
	}
	return res, nil
}

func GetHeaderParamUint64(c *gin.Context, key string) (uint64, error) {
	var res uint64
	resStr := c.GetHeader(key)

	res, err := strconv.ParseUint(resStr, 10, 64)
	if err != nil {
		return 0, errors.ErrorParseParam
	}
	return res, nil
}

func GetRpcReqFromGin(c *gin.Context, ginReqPointer, rpcReqPointer interface{}) error {
	err := c.ShouldBindJSON(ginReqPointer)
	if err != nil {
		return err
	}
	err = util.DeepCopyPHP(ginReqPointer, rpcReqPointer)
	return err
}

func GetQueryUint64(c *gin.Context, key string) uint64 {
	resStr := c.Query(key)
	var res uint64
	if len(resStr) > 0 {
		res, _ = strconv.ParseUint(resStr, 10, 64)
	}
	return res
}
