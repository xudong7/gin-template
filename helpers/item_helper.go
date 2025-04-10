package helpers

import (
	"gin-second-fish/global"
	"gin-second-fish/models"
)

func SelectItemByIdString(id string) (models.Item, error) {
	var item models.Item

	if err := global.Db.AutoMigrate(&models.Item{}); err != nil {
		return item, err
	}

	if err := global.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return item, err
	}

	return item, nil
}
