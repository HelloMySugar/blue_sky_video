package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/HelloMySugar/service-video/types"
)

type Video struct {
	Model

	Name        string `gorm:"type:varchar(64);index;unique;not null"`
	Visited     uint32
	Keyword     string
	Category    *Category `gorm:"association_autoupdate:false"`
	CategoryID  uint
	Url         string `gorm:"type:varchar(128)"`
	CoverUrl    string `gorm:"type:varchar(128)"`
	ReleaseTime time.Time
}

type VideoList []Video

func (m *Video) Create(db *gorm.DB) error {
	err := db.Create(m).Error
	if err != nil {
		return errors.WithMessage(err, "write db")
	}
	return nil
}

func (m *Video) Update(db *gorm.DB) error {
	err := db.Model(m).Updates(*m).Error
	if err != nil {
		return errors.WithMessage(err, "write db")
	}
	return nil
}

func (m *Video) Delete(db *gorm.DB) error {
	err := db.Delete(m).Error
	if err != nil {
		return errors.WithMessage(err, "write db")
	}
	return nil
}

func (m *Video) Fetch(db *gorm.DB) error {
	if err := db.Where(m).First(m).Error; err != nil {
		return errors.WithMessage(err, "get db")
	}
	return nil
}

func (m *Video) FetchList(db *gorm.DB, page types.Pager) (videoList VideoList, total int, err error) {
	scope := db.NewScope(m)
	fieldReleaseTime, _ := scope.FieldByName("ReleaseTime")
	videoList = make(VideoList, 0)
	if err := db.Model(m).Where(m).Count(&total).Error; err != nil {
		return nil, 0, errors.WithMessage(err, "get count")
	}
	if err := db.Where(m).Order(fmt.Sprintf("%s desc", fieldReleaseTime.DBName)).Limit(page.Limit).Offset(page.Offset).Find(&videoList).Error; err != nil {
		return nil, 0, errors.WithMessage(err, "get db")
	}
	return videoList, total, nil
}

func (m *Video) FetchListNameLike(db *gorm.DB, page types.Pager, name string) (videoList VideoList, total int, err error) {
	scope := db.NewScope(m)
	fieldName, _ := scope.FieldByName("Name")
	videoList = make(VideoList, 0)
	if err := db.Model(m).Where(fmt.Sprintf("%s LIKE '%%%s%%'", fieldName.DBName, name)).Count(&total).Error; err != nil {
		return nil, 0, errors.WithMessage(err, "get count")
	}
	if err := db.Where(fmt.Sprintf("%s LIKE '%%%s%%'", fieldName.DBName, name)).Limit(page.Limit).Offset(page.Offset).Find(&videoList).Error; err != nil {
		return nil, 0, errors.WithMessage(err, "get db")
	}
	return videoList, total, nil
}
