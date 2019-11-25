package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
)

var TopicService = newTopicService()

type topicRepository struct{}

func newTopicService() *topicRepository {
	return &topicRepository{}
}

func (this *topicRepository) Get(db *gorm.DB, id int64) *Topic {
	ret := &Topic{}
	db.First(ret, "id = ?", id)
	return ret
}

func (this *topicRepository) Find(db *gorm.DB, wheres map[string]interface{}, offset int) []*Topic {
	ret := []*Topic{}
	if offset < 0 {
		offset = 0
	}
	db.Where(wheres).Offset(offset).Limit(20).Order("id desc").Find(&ret)
	return ret
}

func (this *topicRepository) FindIn(db *gorm.DB, wheres []int64, offset int, limit int) []*Topic {
	ret := []*Topic{}
	db.Where(wheres).Offset(offset).Limit(limit).Order("id desc").Find(&ret)
	return ret
}

func (this *topicRepository) Create(db *gorm.DB, t *Topic) (err error) {
	err = db.Create(t).Error
	return
}

func (this *topicRepository) Update(db *gorm.DB, t *Topic) (err error) {
	err = db.Save(t).Error
	return
}

func (this *topicRepository) Updates(db *gorm.DB, wheres map[string]interface{}, columns map[string]interface{}) (err error) {
	err = db.Model(&Topic{}).Where(wheres).Updates(columns).Error
	return
}

func (this *topicRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&Topic{}, "id = ?", id)
}

func (this *topicRepository) Count(db *gorm.DB, wheres map[string]interface{}) (count int) {
	count = 0
	db.Model(&Topic{}).Where(wheres).Count(&count)
	return
}

func (this *topicRepository) AddTopic(db *gorm.DB, params map[string]string) (*Topic, error) {
	title := strings.TrimSpace(params["title"])
	content := strings.TrimSpace(params["content"])
	categoryId := strings.TrimSpace(params["category_id"])
	experience := strings.TrimSpace(params["experience"])
	userId := strings.TrimSpace(params["user_id"])

	if len(title) == 0 {
		return nil, errors.New("标题不能为空")
	}

	// 验证密码
	if len(content) == 0 {
		return nil, errors.New("内容不能为空")
	}

	categoryIdInt64, _ := strconv.ParseInt(categoryId, 10, 64)
	userId64, _ := strconv.ParseInt(userId, 10, 64)
	experience64, _ := strconv.ParseInt(experience, 10, 64)
	// 分享需要审核 1正常 2审核 3删除
	var topicStatus int
	if categoryIdInt64 == int64(1) {
		topicStatus = 1
	} else if categoryIdInt64 == int64(2) {
		topicStatus = 2
	}

	topic := &Topic{
		Title:           title,
		Content:         content,
		CategoryId:      categoryIdInt64,
		Experience:      experience64,
		UserId:          userId64,
		Status:          topicStatus,
		LastCommentTime: time.Now().UTC(),
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}

	if err := this.Create(db, topic); err != nil {
		return nil, err
	}
	return topic, nil
}

func (this *topicRepository) CountAdd(db *gorm.DB, id int64, name string) {
	db.Model(&Topic{}).Where("id = ?", id).UpdateColumn(name, gorm.Expr(name+"+ ?", 1))
	return
}

//本周热议
func (this *topicRepository) GetTopicHot(db *gorm.DB) []*Topic {
	ret := []*Topic{}
	db.Where("status = ?", 1).Offset(0).Limit(10).Order("comment_count desc").Find(&ret)
	return ret
}
