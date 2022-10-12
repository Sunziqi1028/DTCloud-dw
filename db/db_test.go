/**
 * @Author: Sun
 * @Description:
 * @File:  db_test
 * @Version: 1.0.0
 * @Date: 2022/6/15 20:05
 */

package db

import (
	"fmt"
	"testing"
	"wg315/config"
	"wg315/global"
)

type WaterTest struct {
	Name     string `gorm:"name"`
	Datatype string `gorm:"datatype"`
	Value    string `gorm:"value"`
	Error    string `gorm:"error"`
	ErrMsg   string `gorm:"err_msg"`
}

func TestNewPostgresql(t *testing.T) {
	global.Config = config.NewConfig()
	db, err := NewPostgresql()
	if err != nil {
		t.Error(err)
	}
	Water := WaterTest{
		Datatype: "",
		Value:    "0.009",
		Error:    "",
		ErrMsg:   "",
	}
	err = db.Table("iot_water").Create(&Water).Error
	if err != nil {
		fmt.Println("insert table err:", err)
	}
}
