/**
 * @Author: Sun
 * @Description:
 * @File:  device_device
 * @Version: 1.0.0
 * @Date: 2022/6/15 17:22
 */

package controller

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
	"wg315/db"
	"wg315/global"
	"wg315/http"
	"wg315/models"
	"wg315/utils"
)

type WG024YX struct {
}

func (w WG024YX) HandData(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	j, err := simplejson.NewJson(data)
	if err != nil {
		log.Println("NewJson error:%s", err.Error())
		return
	}
	typeData, _ := j.Get("type").String()
	time, _ := j.Get("time").String()                                        // 获取时间字符串
	year, month, week, day, hour, minute, second := utils.HandTime2Int(time) // 将时间类型分解单个的字段

	switch typeData {
	case global.BEAT:
		var beatData models.DeviceStatus
		json.Unmarshal(data, &beatData)
		beatData.Year = year
		beatData.Month = month
		beatData.Week = week
		beatData.Day = day
		beatData.Hour = hour
		beatData.Minute = minute
		beatData.Second = second
		code := http.PostData(data)
		go func() {
			for _, device := range beatData.Status {
				status, err := db.GetDeviceStatusByID("iot_device", global.DB, device.Name)
				if err != nil {
					fmt.Println("获取设备状态失败, err:", err)
					continue
				}
				if len(status) == 0 {
					err = db.CreateDeviceStatus("iot_device", global.DB, device)
					if err != nil {
						fmt.Println("插入设备状态信息失败， err:", err)
						continue
					}
				}
				if len(status) != 0 && status != device.State {
					err = db.UpdateDeviceStatusByID("iot_device", global.DB, device.Name, device.State)
					if err != nil {
						fmt.Println("更新设备状态信息失败， err:", err)
					}
				}
			}
		}()
		fmt.Println(code)
		c.JSON(200, "successfully~")
	case global.REAlDATA:
		code := http.PostData(data)
		fmt.Println(code)
		if err != nil {
			fmt.Println("get online device name err:", err)
		}

		go func() {
			onlineDeviceNames, _ := db.GetDeviceName("iot_device", global.DB) // 拿到数据库中当前在线的通道名称
			for _, onlineDeviceName := range onlineDeviceNames {
				realData, _ := j.Get("data").Get(onlineDeviceName).MarshalJSON()
				var modelRealData []models.Data
				json.Unmarshal(realData, &modelRealData)
				for _, v := range modelRealData { // 通过遍历获取当前在线通道下的所有设备（之前是把通道写死的，下面注释的代码）
					v.Year = year
					v.Month = month
					v.Week = week
					v.Day = day
					v.Hour = hour
					v.Minute = minute
					v.Second = second
					if strings.Contains(onlineDeviceName, "water") {
						db.CreateData("iot_water", global.DB, v)
					} else {
						db.CreateData("iot_electric", global.DB, v)
					}

				}
			}
			//electricBytes, _ := j.Get("data").Get("electric_D1").MarshalJSON() // 获取电表信息
			//if len(electricBytes) > 0 {
			//	var electric []models.Data
			//	json.Unmarshal(electricBytes, &electric)
			//	fmt.Println(electric, "line:72")
			//	for _, v := range electric {
			//		v.Year = year
			//		v.Month = month
			//		v.Week = week
			//		v.Day = day
			//		v.Hour = hour
			//		v.Minute = minute
			//		v.Second = second
			//		db.CreateData("iot_electric", global.DB, v)
			//	}
			//}
			//
			//waterBytes, _ := j.Get("data").Get("water_D1").MarshalJSON() // 获取水表信息
			//if len(waterBytes) > 0 {
			//	var water []models.Data
			//	json.Unmarshal(waterBytes, &water)
			//	fmt.Println(water, "line:72")
			//	for _, v := range water {
			//		v.Year = year
			//		v.Month = month
			//		v.Week = week
			//		v.Day = day
			//		v.Hour = hour
			//		v.Minute = minute
			//		v.Second = second
			//		db.CreateData("iot_water", global.DB, v)
			//	}
			//}
		}()

		c.JSONP(200, "send data successfully~")
	}
}
