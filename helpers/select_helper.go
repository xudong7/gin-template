package helpers

import (
	"gin-second-fish/global"
	"gin-second-fish/models"
)

// select category by id string
func SelectCategoryByIdString(id string) (models.Category, error) {
	var category models.Category
	if err := global.Db.Where("id = ?", id).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil

}

// select item by id string
func SelectItemByIdString(id string) (models.Item, error) {
	var item models.Item
	if err := global.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

// select order by id string
func SelectOrderByIdString(id string) (models.Order, error) {
	var order models.Order
	if err := global.Db.Where("id = ?", id).First(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

// select user by id string
func SelectUserByIdString(id string) (models.User, error) {
	var user models.User

	if err := global.Db.AutoMigrate(&models.User{}); err != nil {
		return user, err
	}

	if err := global.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
