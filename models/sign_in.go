package models

import "github.com/jinzhu/gorm"

var SignInService = newSignInService()

func newSignInService() *signInRepository {
	return &signInRepository{}
}

type signInRepository struct{}

func (this *signInRepository) Create(db *gorm.DB, t *SignIn) (err error) {
	err = db.Create(t).Error
	return
}

func (this *signInRepository) Find(db *gorm.DB, wheres map[string]interface{}, offset int) []*SignIn {
	ret := []*SignIn{}
	if offset < 0 {
		offset = 0
	}
	db.Where(wheres).Offset(offset).Limit(20).Order("id desc").Find(&ret)
	return ret
}

func (this *signInRepository) Count(db *gorm.DB, wheres map[string]interface{}) (count int) {
	count = 0
	db.Model(&Message{}).Where(wheres).Count(&count)
	return
}

func (this *signInRepository) Updates(db *gorm.DB, wheres map[string]interface{}, columns map[string]interface{}) (err error) {
	err = db.Model(&Message{}).Where(wheres).Updates(columns).Error
	return
}

func (this *signInRepository) Get(db *gorm.DB, wheres map[string]interface{}) *SignIn {
	ret := &SignIn{}
	db.Where(wheres).Order("id desc").First(ret)
	return ret
}
