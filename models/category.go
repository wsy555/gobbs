package models

import (
	"github.com/jinzhu/gorm"
)

var CategoryService = newCategoryRepository()

func newCategoryRepository() *categoryRepository {
	return &categoryRepository{}
}

type categoryRepository struct {
}

func (this *categoryRepository) Get(db *gorm.DB, id int64) *Category {
	ret := &Category{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *categoryRepository) Find(db *gorm.DB, wheres map[string]interface{}, offset int) []*Category {
	ret := []*Category{}
	db.Where(wheres).Offset(offset).Limit(20).Order("id asc").Find(&ret)
	return ret
}

func (this *categoryRepository) Create(db *gorm.DB, t *Category) (err error) {
	err = db.Create(t).Error
	return
}

func (this *categoryRepository) Update(db *gorm.DB, t *Category) (err error) {
	err = db.Save(t).Error
	return
}

func (this *categoryRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&Category{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (this *categoryRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&Category{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (this *categoryRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&Category{}, "id = ?", id)
}
