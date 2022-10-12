/**
 * @Author: Sun
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/6/14 14:53
 */

package main

import (
	"fmt"
	"wg315/config"
	"wg315/db"
	"wg315/global"
	"wg315/router"
)

func init() {
	var err error
	global.Config = config.NewConfig() // 解析全局配置

	global.DB, err = db.NewPostgresql() // 全局DB
	if err != nil {
		fmt.Println("初始化DB err:", err)
	}
}

func main() {
	r := router.NewRouter()
	fmt.Println(global.Config)
	r.Run(fmt.Sprintf("%s:%s", global.Config.IPAddr.Addr, global.Config.IPAddr.Port))
}
