package helpers

import (
	"gin-second-fish/global"
	"gin-second-fish/models"
)

func SelectOrderByIdString(id string) (models.Order, error) {
	var order models.Order
	if err := global.Db.Where("id = ?", id).First(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}
