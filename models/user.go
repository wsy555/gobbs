package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
	"yehelaoren/common"
)

var UserService = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct {
}

type Result struct {
	Id         int
	Nickname   string
	Avatar     string
	TopicNum   int
	CommentNum int
	Experience int
}

func (this *userRepository) Get(db *gorm.DB, id int64) *User {
	ret := &User{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *userRepository) Take(db *gorm.DB, where ...interface{}) *User {
	ret := &User{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *userRepository) Find(db *gorm.DB, wheres map[string]interface{}, offset int) []*User {
	ret := []*User{}
	db.Where(wheres).Offset(offset).Limit(20).Order("id desc").Find(&ret)
	return ret
}

func (this *userRepository) FindIn(db *gorm.DB, wheres []int64, offset int) []*User {
	ret := []*User{}
	db.Where(wheres).Offset(offset).Limit(20).Order("id desc").Find(&ret)
	return ret
}

func (this *userRepository) Create(db *gorm.DB, t *User) (err error) {
	err = db.Create(t).Error
	return
}

func (this *userRepository) Update(db *gorm.DB, t *User) (err error) {
	err = db.Save(t).Error
	return
}

func (this *userRepository) Updates(db *gorm.DB, wheres map[string]interface{}, columns map[string]interface{}) (err error) {
	err = db.Model(&User{}).Where(wheres).Updates(columns).Error
	return
}

func (this *userRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&User{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (this *userRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&User{}, "id = ?", id)
}

func (this *userRepository) GetByKey(db *gorm.DB, key string, value string) *User {
	return this.Take(db, key+" = ?", value)
}

// 注册
func (this *userRepository) DoRegister(db *gorm.DB, params map[string]string) (*User, error) {
	email := strings.TrimSpace(params["email"])
	mobile := strings.TrimSpace(params["mobile"])
	nickname := strings.TrimSpace(params["nickname"])
	password := strings.TrimSpace(params["password"])
	rePassword := strings.TrimSpace(params["rePassword"])
	avatar := strings.TrimSpace(params["avatar"])

	if len(nickname) == 0 {
		return nil, errors.New("昵称不能为空")
	}

	// 验证密码
	if password != rePassword {
		return nil, errors.New("2次密码不一致")
	}

	// 如果设置了邮箱，那么需要验证邮箱
	if len(email) > 0 {
		if common.IsValidateEmail(email) != nil {
			return nil, errors.New("邮箱格式错误")
		}

		if this.GetByKey(db, "email", email) != nil {
			return nil, errors.New("邮箱：" + email + " 已被占用")
		}
	}

	if common.IsValidateMobile(mobile) != nil {
		return nil, errors.New("手机号格式错误")
	}

	// 验证用户名是否存在
	oldUser := this.GetByKey(db, "nickname", nickname)
	if oldUser != nil {
		return nil, errors.New("昵称已经被占用")
	}

	user := &User{
		Email:      email,
		Nickname:   nickname,
		Password:   common.Md5Encode(password),
		Avatar:     avatar,
		Status:     1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Experience: 20,
	}

	if err := this.Create(db, user); err != nil {
		return nil, err
	}
	return user, nil
}

// 登录
func (this *userRepository) DoLogin(db *gorm.DB, email, password string) (*User, error) {
	if len(password) == 0 {
		return nil, errors.New("密码不能为空")
	}

	var user *User
	if err := common.IsValidateEmail(email); err == nil { // 如果用户输入的是邮箱
		user = this.GetByKey(db, "email", email)
	} else {
		err := common.IsValidateMobile(email)
		if err == nil { //输入是手机号
			user = this.GetByKey(db, "mobile", email)
		}
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	if user.Password != common.Md5Encode(password) {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (this *userRepository) CountAdd(db *gorm.DB, id int64, name string, num int) {
	db.Model(&User{}).Where("id = ?", id).UpdateColumn(name, gorm.Expr(name+"+ ?", num))
	return
}

func (this *userRepository) UploadAvatar(db *gorm.DB, id int64, name string, num int) {

	return
}

// 最新回帖榜
func (this *userRepository) GetUserReplyHot(db *gorm.DB) []*Result {
	var ret []*Result
	// 接口过滤明感字段
	db.Model(&User{}).Select("id, nickname, avatar, comment_num").Offset(0).Limit(12).Order("comment_num desc").Scan(&ret)
	return ret
}
