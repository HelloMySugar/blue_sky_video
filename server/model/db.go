package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/HelloMySugar/service-video/global/settings"
)

var DBObj *gorm.DB

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id,string,omitempty"`
	CreatedAt time.Time  `json:"createAt"`
	UpdatedAt time.Time  `json:"updateAt"`
	DeletedAt *time.Time `sql:"index" json:"deleteAt"`
}

func Setup() {
	var err error
	DBObj, err = gorm.Open(settings.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		settings.DatabaseSetting.User,
		settings.DatabaseSetting.Password,
		settings.DatabaseSetting.Host,
		settings.DatabaseSetting.Name))

	if err != nil {
		log.Panicf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return settings.DatabaseSetting.TablePrefix + defaultTableName
	}

	DBObj.LogMode(true)
	DBObj.SingularTable(true)
	DBObj.DB().SetMaxIdleConns(10)
	DBObj.DB().SetMaxOpenConns(100)
	/*
		DBObj.DropTable(
			&Category{},
		)
	*/
	DBObj.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(
		&Category{},
		&Video{},
	)
}
