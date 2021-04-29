/**
 *@Description
 *@ClassName map
 *@Date 2020/12/29 上午12:28
 *@Author ckhero
 */

package _map

import "strings"

func SeparateMapKeyByBlank(data map[string]struct{}) string {
	res := ""
	for k, _ := range data {
		res += k + " "
	}
	return strings.Trim(res, " ")
}
