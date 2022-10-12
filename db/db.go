/**
 * @Author: Sun
 * @Description:
 * @File:  db
 * @Version: 1.0.0
 * @Date: 2022/6/15 19:47
 */

package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"wg315/global"
)

func NewPostgresql() (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		global.Config.PGConfig.Host,
		global.Config.PGConfig.UserName,
		global.Config.PGConfig.Password,
		global.Config.PGConfig.DBName,
		global.Config.PGConfig.DBPort)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("db connectde failed, err:", err)
		return nil, err
	}
	return db, nil
}
