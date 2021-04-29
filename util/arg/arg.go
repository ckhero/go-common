/**
 *@Description
 *@ClassName arg
 *@Date 2020/11/23 3:05 下午
 *@Author ckhero
 */

package arg

import (
	"os"
	"regexp"
	"strings"
)

func GetArg(name, defaultValue string) string {
	args := os.Args
	pattern := name + "="
	for _, arg := range args {
		if match, _ := regexp.MatchString(pattern, arg); match {
			value := strings.Split(arg, "=")
			return strings.TrimSpace(value[1])
		}
	}
	return strings.TrimSpace(defaultValue)
}
