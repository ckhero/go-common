/**
 *@Description
 *@ClassName funck_test
 *@Date 2021/1/26 上午10:55
 *@Author ckhero
 */

package funk

import (
	"fmt"
	"github.com/thoas/go-funk"
	"testing"
)

type Foo struct {
	ID        int
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int
}

func  TestGetUint64s(t *testing.T) {
	//bar := &Bar{
	//	Name: "Test",
	//}

	foo1 := &Foo{
		ID:        1,
		FirstName: "Dark",
		LastName:  "Vador",
		Age:       20,
		//Bar:       bar,
	}

	foo2 := &Foo{
		ID:        1,
		FirstName: "Dark",
		LastName:  "Vador",
		Age:       30,
	} // foo2.Bar is nil

	fmt.Println(funk.Get([]*Foo{foo1, foo2}, "LastName")) // []string{"Test"}
	//funk.Get(foo2, "Bar.Name") /
}
