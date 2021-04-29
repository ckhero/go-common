/**
 *@Description
 *@ClassName json
 *@Date 2020/11/2 6:07 下午
 *@Author ckhero
 */

package json

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func init()  {
	extra.RegisterFuzzyDecoders()
}
/**
 * json序列化复制
 */
func DeepCopy(src, dst interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func JSONToMap(src string, dst interface{}) error {

	return json.Unmarshal([]byte(src), dst)
}

func DeepCopyPHP(src, dst interface{}) error {

	b, err := jsoniter.Marshal(src)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(b, dst)
}

func DeepCopyAndRtn(src, dst interface{}) (interface{}, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, dst)
	return dst, err
}