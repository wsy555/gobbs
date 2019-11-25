package models

import (
	"github.com/jinzhu/gorm"
)

var TopicCommentService = newTopicCommentService()

type topicCommentRepository struct{}

func newTopicCommentService() *topicCommentRepository {
	return &topicCommentRepository{}
}

func (this *topicCommentRepository) Get(db *gorm.DB, id int64) *TopicComment {
	ret := &TopicComment{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *topicCommentRepository) Find(db *gorm.DB, wheres map[string]interface{}, offset int) []*TopicComment {
	ret := []*TopicComment{}
	if offset < 0 {
		offset = 0
	}
	db.Where(wheres).Offset(offset).Limit(20).Order("id desc").Find(&ret)
	return ret
}

func (this *topicCommentRepository) Create(db *gorm.DB, t *TopicComment) (err error) {
	err = db.Create(t).Error
	return
}

func (this *topicCommentRepository) Update(db *gorm.DB, t *TopicComment) (err error) {
	err = db.Save(t).Error
	return
}

func (this *topicCommentRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&TopicComment{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (this *topicCommentRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&TopicComment{}, "id = ?", id)
}

func (this *topicCommentRepository) LikeAdd(db *gorm.DB, id int64, num int) {
	db.Model(&TopicComment{}).Where("id = ?", id).UpdateColumn("like_count", gorm.Expr("like_count + ?", num))
	return
}
