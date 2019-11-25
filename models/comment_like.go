package models

import "github.com/jinzhu/gorm"

var CommentLikeService = newCommentLikeService()

func newCommentLikeService() *commentLikeRepository {
	return &commentLikeRepository{}
}

type commentLikeRepository struct{}

func (*commentLikeRepository) GetLikes(db *gorm.DB, userId int64, CommentIds []int64) []*CommentLike {
	ret := []*CommentLike{}
	db.Where("user_id = ? AND comment_id in (?)", userId, CommentIds).Offset(0).Limit(20).Find(&ret)
	return ret
}

func (this *commentLikeRepository) Create(db *gorm.DB, t *CommentLike) (err error) {
	err = db.Create(t).Error
	return
}

func (this *commentLikeRepository) Delete(db *gorm.DB, commentId int64, userId int64) {
	db.Delete(&CommentLike{}, "comment_id = ? and user_id = ?", commentId, userId)
}
