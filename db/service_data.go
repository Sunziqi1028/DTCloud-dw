/**
 * @Author: Sun
 * @Description:
 * @File:  service_data
 * @Version: 1.0.0
 * @Date: 2022/6/16 16:01
 */

package db

import (
	"gorm.io/gorm"
	"wg315/models"
)

func CreateData(tableName string, db *gorm.DB, model models.Data) error {
	err := db.Table(tableName).Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

//func GetData(tableName string, db *gorm.DB, name string) (*models.IdData, error) {
//	var model models.IdData
//	err := db.Table(tableName).Where("name = ?", name).First(&model).Error
//	if err != nil {
//		return nil, err
//	}
//	return &model, nil
//}

//func UpdateData(tabelName string, db *gorm.DB, name, status string) error {
//	err := db.Table(tabelName).Where("name = ?", name).Update("status", status).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
