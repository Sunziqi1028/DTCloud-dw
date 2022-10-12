/**
 * @Author: Sun
 * @Description:
 * @File:  model
 * @Version: 1.0.0
 * @Date: 2022/6/14 14:59
 */

package models

// DeviceStatus 设备心跳状态
type DeviceStatus struct {
	Type    string   `json:"type,omitempty"`
	Gateway string   `json:"gateway"`
	Time    string   `json:"time"`
	Status  []Status `json:"status"`
	Year    int      `json:"year,omitempty"`
	Month   int      `json:"month,omitempty"`
	Week    int      `json:"week,omitempty"`
	Day     int      `json:"day,omitempty"`
	Hour    int      `json:"hour,omitempty"`
	Minute  int      `json:"minute,omitempty"`
	Second  int      `json:"second,omitempty"`
}

type Status struct {
	Name  string `json:"id,omitempty"`
	State string `json:"status,omitempty"` // 0：离线 1：在线
}

// DeviceData 上传的设备数据
type DeviceData struct {
	Type    string              `json:"type"`
	Gateway string              `json:"gateway"`
	Time    string              `json:"time"`
	Data    map[string]struct{} `json:"data"`
}

// Data 用于存表的数据
type Data struct {
	Name     string `json:"id" ,gorm:"name"`
	Datatype string `json:"datatype" ,gorm:"datatype"`
	Value    string `json:"value"`
	Error    string `json:"err"`
	ErrMsg   string `json:"err_msg"`
	Year     int    `json:"year,omitempty"`
	Month    int    `json:"month,omitempty"`
	Week     int    `json:"week,omitempty"`
	Day      int    `json:"day,omitempty"`
	Hour     int    `json:"hour,omitempty"`
	Minute   int    `json:"minute,omitempty"`
	Second   int    `json:"second,omitempty"`
}
