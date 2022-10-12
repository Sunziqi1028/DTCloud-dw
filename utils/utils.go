/**
 * @Author: Sun
 * @Description:
 * @File:  utils
 * @Version: 1.0.0
 * @Date: 2022/6/18 16:48
 */

package utils

import (
	"time"
)

func HandTime2Int(t string) (year, month, day, week, hour, minute, second int) {
	tmpTime, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	//t.Format("2006-01-02 15:04:05")
	year, week = tmpTime.ISOWeek()

	return year, int(tmpTime.Month()), week, tmpTime.Day(), tmpTime.Hour(), tmpTime.Minute(), tmpTime.Second()

}
