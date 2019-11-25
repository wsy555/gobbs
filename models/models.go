package models

import "time"

var Models = []interface{}{
	&User{}, &Category{}, &TopicComment{}, &Favorite{},
	&Topic{}, &TopicTag{}, &Message{}, &SysConfig{},
	&CommentLike{},
}

type Model struct {
	Id int64 `gorm:"PRIMARY_KEY;unsigned;AUTO_INCREMENT"`
}

//用户表
type User struct {
	Model
	Email       string `gorm:"size:128;unique;not null;index:idx_email;"`
	Mobile      string `gorm:"size:11;unique;index:idx_mobile;"`
	Nickname    string `gorm:"size:16;not null"`
	Avatar      string `gorm:"size:200"`
	Password    string `gorm:"size:200;not null"`
	Status      int    `gorm:"type:tinyint;size:1;unsigned;not null"` // 1正常 2屏蔽
	Roles       string `gorm:"size:200"`
	Description string `gorm:"size:512"`
	TopicNum    int    `gorm:"int;unsigned;not null;default 0;"`
	CommentNum  int    `gorm:"int;unsigned;not null;default 0;"`
	Experience  int    `gorm:"int;unsigned;not null;default 0;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// 分类
type Category struct {
	Model
	Name        string `gorm:"size:32;unique;not null"`
	Description string `gorm:"size:512"`
	Status      int    `gorm:"type:tinyint;size:1;unsigned;not null"`
	CreatedAt   time.Time
}

// 评论
type TopicComment struct {
	Model
	UserId    int64  `gorm:"unsigned;index:idx_user_id;not null"`  // 用户编号
	TopicId   int64  `gorm:"unsigned;index:idx_topic_id;not null"` // 被评论实体编号
	Content   string `gorm:"type:text;not null"`                   // 内容
	ParentId  int64  `gorm:"int;not null;unsigned;default:0;" `    // 引用的评论编号
	LikeCount int64  `gorm:"int;not null;unsigned" `               // 赞数
	IsAccept  int    `gorm:"tinyint;size:1;unsigned"`              // 是否采纳 1是 0否
	Status    int    `gorm:"tinyint;size:1;unsigned;"`             // 状态：0：待审核、1：审核通过、2：审核失败
	CreatedAt time.Time
}

// 话题点赞
type CommentLike struct {
	Model
	UserId    int64     `gorm:"not null;unsigned;index:idx_user_id;"`    // 用户
	CommentId int64     `gorm:"not null;unsigned;index:idx_comment_id;"` // 主题编号
	CreatedAt time.Time // 创建时间
}

// 收藏
type Favorite struct {
	Model
	UserId    int64  `gorm:"unsigned;index:idx_user_id;not null" `  // 用户编号
	TopicId   int64  `gorm:"unsigned;index:idx_topic_id;not null" ` // 编号
	Title     string `gorm:"size:20;"`
	CreatedAt time.Time
}

// 主题
type Topic struct {
	Model
	UserId          int64     `gorm:"not null;unsigned;index:idx_user_id;" `                                  // 用户
	CategoryId      int64     `gorm:"type:tinyint;size:1;unsigned;not null;default:0;index:idx_category_id;"` // 用户
	IsTop           int64     `gorm:"type:tinyint;size:1;unsigned;not null;default:0;"`                       // 用户
	IsBest          int64     `gorm:"type:tinyint;size:1;unsigned;not null;default:0;"`                       // 用户
	Title           string    `gorm:"size:128;" json:"title" `                                                // 标题
	Content         string    `gorm:"type:text" json:"content" `                                              // 内容
	ViewCount       int64     `gorm:"int;unsigned;default:0;" `                                               // 查看数量
	CommentCount    int64     `gorm:"int;unsigned;default:0;" `                                               // 跟帖数量
	Experience      int64     `gorm:"int;unsigned;default:0;" `                                               // 经验数量
	Status          int       `gorm:"tinyint;size:1;index:idx_status;default:0;"`                             // 状态：1：正常、2：待审核 3删除
	LastCommentTime time.Time // 最后回复时间
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// 主题标签
type TopicTag struct {
	Model
	TopicId   int64     `gorm:"not null;unsigned;index:idx_topic_id;"`             // 主题编号
	TagId     int64     `gorm:"smallint;unsigned;not null;index:idx_tag_id;"`      // 标签编号
	Status    int64     `gorm:"tinyint;size:1;unsigned;not null;index:idx_status"` // 状态：正常、删除
	CreatedAt time.Time // 创建时间
}

// 消息
type Message struct {
	Model
	FromId    int64     `gorm:"not null;unsigned;" json:"fromId"`      // 消息发送人
	UserId    int64     `gorm:"not null;unsigned;index:idx_user_id;" ` // 用户编号(消息接收人)
	Content   string    `gorm:"type:text;not null"`                    // 消息内容
	ObjType   int       `gorm:"tinyint;not null;unsigned" `            // 消息类型1:系统 2:帖子回复
	ObjId     int64     `gorm:"tinyint;not null;unsigned" `            // 消息关联ID
	Status    int       `gorm:"tinyint;not null;unsigned" `            // 状态：0：未读、1：已读
	CreatedAt time.Time // 创建时间
}

// 系统配置
type SysConfig struct {
	Model
	Key         string    `gorm:"not null;size:128;unique" ` // 配置key
	Value       string    `gorm:"type:text"`                 // 配置值
	Name        string    `gorm:"not null;size:32"`          // 配置名称
	Description string    `gorm:"size:128"`                  // 配置描述
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间
}

type SignIn struct {
	Model
	UserId     int64  `gorm:"unsigned;not null;index:idx_user_id;"` //用户
	Days       int    `gorm:"unsigned;not null;default:0;"`         // 次数
	Experience int    `gorm:"unsigned;not null;default:0;"`         //经验
	Day        string `gorm:"size:10;"`                             //日期
	CreatedAt  time.Time
}

//日历库
type LunarCalendar struct {
	Model
	Day        string `gorm:"size:10;unique" `
	LunarYear  string `gorm:"size:5;"`
	LunarMonth string `gorm:"size:5;"`
	LunarDay   string `gorm:"size:5;"`
	Suit       string `gorm:"size:200;"`
	Taboo      string `gorm:"size:50;"`
	Jieqi      string `gorm:"size:50;"`
	CreatedAt  time.Time
}
