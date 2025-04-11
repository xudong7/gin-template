package helpers

import (
	"gin-second-fish/global"
	"gin-second-fish/models"
)

func SelectCategoryByIdString(id string) (models.Category, error) {
	var category models.Category
	if err := global.Db.Where("id = ?", id).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil

}
