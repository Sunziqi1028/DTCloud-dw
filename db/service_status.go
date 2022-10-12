/**
 * @Author: Sun
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2022/6/16 15:28
 */

package db

import (
	"gorm.io/gorm"
	"wg315/models"
)

// CreateDeviceStatus 插入设备状态信息
func CreateDeviceStatus(tableName string, db *gorm.DB, model models.Status) error {
	err := db.Table(tableName).Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

// GetDeviceStatusByID 通过设备ID 获取设备状态: 0 离线, 1 在线
func GetDeviceStatusByID(tableName string, db *gorm.DB, id string) (string, error) {
	var status string
	err := db.Table(tableName).Select("state").Where("name = ?", id).Find(&status).Error
	if err != nil {
		return "", err
	}
	return status, nil
}

// UpdateDeviceStatusByID 通过设备ID更新设备状态
func UpdateDeviceStatusByID(tabelName string, db *gorm.DB, id, status string) error {
	err := db.Table(tabelName).Where("name = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

// GetDeviceName 获取所有当前在线的采集通道名称
func GetDeviceName(tableName string, db *gorm.DB) (onlineDeviceName []string, err error) {
	err = db.Table(tableName).Select("name").Where("status = 1").Find(&onlineDeviceName).Error
	if err != nil {
		return nil, err
	}
	return onlineDeviceName, nil
}
