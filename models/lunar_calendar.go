package models

import (
	"github.com/jinzhu/gorm"
	"strings"
)

var LunarService = newLunarService()

func newLunarService() *lunarRepository {
	return &lunarRepository{}
}

type lunarRepository struct{}

func (this *lunarRepository) Create(db *gorm.DB, t *LunarCalendar) (err error) {
	err = db.Create(t).Error
	return
}

func (this *lunarRepository) Get(db *gorm.DB, wheres map[string]interface{}) *LunarCalendar {
	ret := &LunarCalendar{}
	db.Where(wheres).Order("id desc").First(ret)
	return ret
}

func (this *lunarRepository) GetRowByDay(db *gorm.DB, day string) *LunarCalendar {

	ret := &LunarCalendar{}

	// 格式化
	day = strings.Replace(day, "-", "", -1)

	wheres := make(map[string]interface{})
	wheres["day"] = day

	db.Where(wheres).Order("id desc").First(ret)

	return ret
}
