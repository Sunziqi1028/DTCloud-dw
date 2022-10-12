/**
 * @Author: Sun
 * @Description:
 * @File:  global
 * @Version: 1.0.0
 * @Date: 2022/6/15 17:12
 */

package global

import (
	"gorm.io/gorm"
	"wg315/config"
)

const (
	BEAT     = "beat"
	REAlDATA = "realdata"
)

var Config *config.Config

var DB *gorm.DB
