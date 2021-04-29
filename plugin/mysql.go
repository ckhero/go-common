/**
 *@Description
 *@ClassName mysql
 *@Date 2020/11/25 9:25 上午
 *@Author ckhero
 */

package plugin

import (
	"fmt"
	"github.com/ckhero/go-common/config"
	"github.com/ckhero/go-common/db/mysql"
)

type Mysql struct {
}

func NewPluginMysql() *Mysql {
	return &Mysql{
	}
}

func (r *Mysql) InitPlugin() error {
	mysql.ConnectDB(config.GetDatabaseCfg())
	return nil
}

func (r *Mysql) Release() {
	mysql.CloseDB()
	fmt.Println("redis release")

}