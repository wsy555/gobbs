package models

import "github.com/jinzhu/gorm"

var MessageService = newMessageService()

func newMessageService() *messageRepository {
	return &messageRepository{}
}

type messageRepository struct{}

func (this *messageRepository) Create(db *gorm.DB, t *Message) (err error) {
	err = db.Create(t).Error
	return
}

func (this *messageRepository) Find(db *gorm.DB, wheres map[string]interface{}, offset int) []*Message {
	ret := []*Message{}
	if offset < 0 {
		offset = 0
	}
	db.Where(wheres).Offset(offset).Limit(20).Order("id desc").Find(&ret)
	return ret
}

func (this *messageRepository) Count(db *gorm.DB, wheres map[string]interface{}) (count int) {
	count = 0
	db.Model(&Message{}).Where(wheres).Count(&count)
	return
}

func (this *messageRepository) Updates(db *gorm.DB, wheres map[string]interface{}, columns map[string]interface{}) (err error) {
	err = db.Model(&Message{}).Where(wheres).Updates(columns).Error
	return
}
