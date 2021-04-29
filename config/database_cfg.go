/**
 *@Description
 *@ClassName database
 *@Date 2021/4/29 下午3:05
 *@Author ckhero
 */

package config

import "gorm.io/gorm/logger"

/**
数据库配置
*/
type Database struct {
	Dialect        string
	Database       string
	Username       string
	Password       string
	Host           string
	Port           int
	Charset        string
	MaxIdleConnNum int
	MaxOpenConnNum int
	LogMode        logger.LogLevel
}

func GetDatabaseCfg() map[string]Database {
	return appConfig.Database
}