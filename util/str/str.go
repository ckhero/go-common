/**
 *@Description
 *@ClassName str
 *@Date 2021/2/7 下午4:52
 *@Author ckhero
 */

package str

func NumToStr(i int) string {
	arr := []string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O",
		"P","Q","R","S","T","U","V","W","X","Y","Z"}
	if i < len(arr) {
		return arr[i]
	}
	return ""
}
