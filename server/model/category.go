package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Category struct {
	Model

	Name string `gorm:"type:varchar(32);index;unique;not null" json:"name"`
}

type CategoryList []Category

func (m *Category) FetchAllList(db *gorm.DB) (categoryList CategoryList, total int, err error) {
	categoryList = make(CategoryList, 0)
	if err := db.Model(m).Count(&total).Error; err != nil {
		return nil, 0, errors.WithMessage(err, "get count")
	}
	if err := db.Where(m).Find(&categoryList).Error; err != nil {
		return nil, 0, errors.WithMessage(err, "get db")
	}
	return categoryList, total, nil
}
