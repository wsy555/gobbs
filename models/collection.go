package models

import "github.com/jinzhu/gorm"

var CollectionService = newCollectionService()

func newCollectionService() *collectionRepository {
	return &collectionRepository{}
}

type collectionRepository struct{}

func (this *collectionRepository) Create(db *gorm.DB, t *Favorite) (err error) {
	err = db.Create(t).Error
	return
}

func (this *collectionRepository) Get(db *gorm.DB, wheres map[string]interface{}) *Favorite {
	ret := &Favorite{}
	db.Where(wheres).Order("id desc").First(ret)
	return ret
}

func (this *collectionRepository) Delete(db *gorm.DB, userId int64, topicId int64) {
	db.Delete(&Favorite{}, "user_id = ? AND topic_id = ?", userId, topicId)
}

func (this *collectionRepository) Find(db *gorm.DB, wheres map[string]interface{}, offset int) []*Favorite {
	ret := []*Favorite{}
	db.Where(wheres).Offset(offset).Limit(10).Order("id desc").Find(&ret)
	return ret
}

func (this *collectionRepository) Count(db *gorm.DB, wheres map[string]interface{}) (count int) {
	count = 0
	db.Model(&Favorite{}).Where(wheres).Count(&count)
	return
}
