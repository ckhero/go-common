/**
 *@Description
 *@ClassName slice
 *@Date 2020/12/12 上午11:24
 *@Author ckhero
 */

package slice

import "strconv"

func ContainsUint64(target []uint64, search uint64) bool {
	for i := 0; i < len(target); i++ {
		if target[i] == search {
			return true
		}
	}
	return false
}

func ConvertStrToUint64(src []string) []uint64 {
	res := make([]uint64, 0)
	for _, item := range src {
		tmp, _ := strconv.ParseUint(item, 10, 64)
		res = append(res, tmp)
	}
	return res
}
