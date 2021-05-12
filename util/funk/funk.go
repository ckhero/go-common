/**
 *@Description
 *@ClassName funk
 *@Date 2021/1/26 上午10:38
 *@Author ckhero
 */

package funk

import (
	"fmt"
	"github.com/thoas/go-funk"
)

func GetUint64s(src interface{}, path string) []uint64 {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	if src == nil {
		return []uint64{}
	}
	res := funk.Get(src, path)
	if res == nil {
		return []uint64{}
	}
	return res.([]uint64)
}
