package models

import (
	"github.com/jinzhu/gorm"
)

var TopicTagService = newTopicTagRepository()

func newTopicTagRepository() *topicTagRepository {
	return &topicTagRepository{}
}

type topicTagRepository struct {
}

func (this *topicTagRepository) Get(db *gorm.DB, id int64) *TopicTag {
	ret := &TopicTag{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *topicTagRepository) Find(db *gorm.DB, wheres map[string]interface{}, offset int) []*TopicTag {
	ret := []*TopicTag{}
	db.Where(wheres).Offset(offset).Limit(20).Order("id desc").Find(ret)
	return ret
}

func (this *topicTagRepository) Create(db *gorm.DB, t *TopicTag) (err error) {
	err = db.Create(t).Error
	return
}

func (this *topicTagRepository) Update(db *gorm.DB, t *TopicTag) (err error) {
	err = db.Save(t).Error
	return
}

func (this *topicTagRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&TopicTag{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (this *topicTagRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&TopicTag{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (this *topicTagRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&TopicTag{}, "id = ?", id)
}
