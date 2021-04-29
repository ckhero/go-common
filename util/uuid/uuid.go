/**
 *@Description
 *@ClassName uuid
 *@Date 2020/12/2 12:57 下午
 *@Author ckhero
 */

package uuid

import (
	"github.com/ckhero/go-common/util/idcreator/snowflake"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
)

func GenUUID() uint64 {
	// 后面再修改
	return snowflake.NextID()
}

func Gen8UUID() string {
	chars := []string{
		"a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5",
		"6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
		"W", "X", "Y", "Z",
	}
	uuidStr := uuid.NewV4().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	resArr := make([]string, 8)
	for i := 0; i < 8; i++ {
		tmp := uuidStr[4*i : 4*(i+1)]
		tmpHex, _ := strconv.ParseInt(tmp, 16, 32)
		resArr[i] = chars[tmpHex%62]

	}
	return strings.Join(resArr, "")
}