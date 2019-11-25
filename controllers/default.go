package controllers

import (
	"fmt"
	"strconv"
	"time"
	"yehelaoren/models"
)

type MainController struct {
	BaseController
}

func (c *MainController) Home() {
	c.TplName = "index.html"
	//置顶帖子3条
	topTopicWheres := make(map[string]interface{})
	topTopicWheres["is_top"] = 1
	topTopicData := models.TopicService.Find(db, topTopicWheres, 0)
	var l int
	if len(topTopicData) > 3 {
		l = 3
	} else {
		l = len(topTopicData)
	}
	topTopicList := topTopicData[0:l]
	//帖子列表
	topicWheres := make(map[string]interface{})
	rtList := models.TopicService.Find(db, topicWheres, 0)

	//用户信息
	userIdArr := make([]int64, 0)
	for _, v := range rtList {
		userIdArr = append(userIdArr, v.UserId)
	}
	for _, v := range topTopicList {
		userIdArr = append(userIdArr, v.UserId)
	}

	userList := models.UserService.FindIn(db, userIdArr, 0)
	userListMap := make(map[int64]string)
	userAvatarMap := make(map[int64]string)
	for _, v := range userList {
		userListMap[v.Id] = v.Nickname
		userAvatarMap[v.Id] = v.Avatar
	}

	// 分类
	cateGoryWhere := make(map[string]interface{})
	categoryList := models.CategoryService.Find(db, cateGoryWhere, 0)

	categoryListMap := make(map[int64]string)

	for _, v := range categoryList {
		categoryListMap[v.Id] = v.Name
	}

	//本周热议
	topicHots := models.TopicService.GetTopicHot(db)
	c.Data["topicHots"] = topicHots

	c.Data["topicList"] = rtList
	c.Data["topTopicList"] = topTopicList
	c.Data["categoryListMap"] = categoryListMap
	c.Data["userListMap"] = userListMap
	c.Data["userAvatarMap"] = userAvatarMap
}

// 回帖榜

func (c *MainController) TopReply() {
	//回帖周榜
	userReplyHots := models.UserService.GetUserReplyHot(db)
	rtData := make(map[string]interface{})
	rtData["data"] = userReplyHots
	rtData["status"] = 0
	c.Data["json"] = rtData
	c.ServeJSON()
	c.StopRun()
}

// 帖子列表
func (c *MainController) Column() {
	c.TplName = "jie/index.html"

	categoryId, _ := c.GetInt(":categoryId")
	page, _ := c.GetInt("page", 1)

	if page < 0 {
		page = 1
	}

	offset := (page - 1) * 20

	topicWheres := make(map[string]interface{})
	if categoryId > 0 {
		topicWheres["category_id"] = categoryId
	}
	topicWheres["status"] = 1

	rtList := models.TopicService.Find(db, topicWheres, offset)

	//用户信息
	userIdArr := make([]int64, 0)
	for _, v := range rtList {
		userIdArr = append(userIdArr, v.UserId)
	}
	userList := models.UserService.FindIn(db, userIdArr, 0)
	userListMap := make(map[int64]string)
	userAvatarMap := make(map[int64]string)
	for _, v := range userList {
		userListMap[v.Id] = v.Nickname
		userAvatarMap[v.Id] = v.Avatar

	}

	// 分类
	cateGoryWhere := make(map[string]interface{})
	categoryList := models.CategoryService.Find(db, cateGoryWhere, 0)

	categoryListMap := make(map[int64]string)

	for _, v := range categoryList {
		categoryListMap[v.Id] = v.Name
	}

	//本周热议
	topicHots := models.TopicService.GetTopicHot(db)
	c.Data["topicHots"] = topicHots

	c.Data["pagePrevious"] = page - 1
	c.Data["pageNext"] = page + 1
	c.Data["categoryId"] = categoryId
	c.Data["topicList"] = rtList
	c.Data["categoryListMap"] = categoryListMap
	c.Data["userListMap"] = userListMap
	c.Data["userAvatarMap"] = userAvatarMap
}

// 详情
func (c *MainController) Detail() {
	id, _ := c.GetInt64(":id")
	topic := models.TopicService.Get(db, id)

	if topic == nil {
		c.redirect(c.URLFor("MainController.Home"))
	}
	c.Data["topic"] = topic
	//回复列表
	wheres := make(map[string]interface{})
	wheres["topic_id"] = id
	replyList := models.TopicCommentService.Find(db, wheres, 0)
	c.Data["replyList"] = replyList

	// 用户信息
	user := models.UserService.Get(db, topic.UserId)
	c.Data["user"] = user

	// 回复列表用户信息
	userIdArr := make([]int64, 0)
	commentIdArr := make([]int64, 0)
	for _, v := range replyList {
		userIdArr = append(userIdArr, v.UserId)
		commentIdArr = append(commentIdArr, v.Id)
	}
	userList := models.UserService.FindIn(db, userIdArr, 0)
	userListMap := make(map[int64]string)
	userAvatarMap := make(map[int64]string)
	for _, v := range userList {
		userListMap[v.Id] = v.Nickname
		userAvatarMap[v.Id] = v.Avatar
	}
	//我是否点赞
	var commentLikes []*models.CommentLike
	if c.userId > 0 {
		commentLikes = models.CommentLikeService.GetLikes(db, c.userId, commentIdArr)
	}
	likeListMap := make(map[int64]bool)
	for _, v := range commentLikes {
		likeListMap[v.CommentId] = true
	}

	//本周热议
	topicHots := models.TopicService.GetTopicHot(db)
	c.Data["topicHots"] = topicHots

	c.Data["likeListMap"] = likeListMap
	c.Data["userListMap"] = userListMap
	c.Data["userAvatarMap"] = userAvatarMap

	c.TplName = "jie/detail.html"
}

func (c *MainController) Add() {
	if c.userId < 0 {
		c.redirect(c.URLFor("UserController.Login"))
	}
	topicId, _ := c.GetInt64("topic_id")
	topic := models.TopicService.Get(db, topicId)

	cateGoryWhere := make(map[string]interface{})
	categoryList := models.CategoryService.Find(db, cateGoryWhere, 0)
	fmt.Println("topicID:", topic.Id)
	c.Data["topic"] = topic
	c.Data["categoryList"] = categoryList
	c.TplName = "jie/add.html"
}

func (c *MainController) AddDo() {
	if c.userId > 0 {
		var params = make(map[string]string, 0)
		params["title"] = c.GetString("title")
		params["category_id"] = c.GetString("category_id")
		params["content"] = c.GetString("content")
		params["experience"] = c.GetString("experience")
		params["user_id"] = strconv.FormatInt(c.userId, 10)
		topicId, _ := c.GetInt64("topic_id")
		// 如果是编辑直接更新
		if topicId > 0 {
			wheres := make(map[string]interface{})
			updates := make(map[string]interface{})
			wheres["user_id"] = c.userId
			wheres["id"] = topicId
			updates["title"] = params["title"]
			updates["content"] = params["content"]
			updates["category_id"], _ = c.GetInt64("category_id")
			models.TopicService.Updates(db, wheres, updates)
			c.Data["json"] = map[string]interface{}{"status": 0, "msg": "编辑成功！", "action": "/"}
			c.ServeJSON()
			c.StopRun()
		}

		// 飞吻不够不能发帖
		experience, _ := c.GetInt("experience")
		user := models.UserService.Get(db, c.userId)
		if user == nil || user.Experience < experience {
			c.Data["json"] = map[string]interface{}{"status": 1, "msg": "积分不够~查看积分说明"}
			c.ServeJSON()
			c.StopRun()
		}

		_, err := models.TopicService.AddTopic(db, params)
		if err != nil {
			c.Data["json"] = map[string]interface{}{"status": 1, "msg": err.Error()}
		} else {
			// 用户帖子数+1
			models.UserService.CountAdd(db, c.userId, "topic_num", 1)
			// 发帖减少积分
			models.UserService.CountAdd(db, c.userId, "experience", -experience)

			c.Data["json"] = map[string]interface{}{"status": 0, "msg": "发表成功", "action": "/"}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"status": 1, "msg": "请先登录", "action": "/user/login.html"}
	}
	c.ServeJSON()
	c.StopRun()
}

// 回复
func (c *MainController) ReplyDo() {
	if c.userId > 0 {
		var t = &models.TopicComment{}

		t.TopicId, _ = c.GetInt64("topic_id")
		t.Content = c.GetString("content")
		t.UserId = c.userId
		t.CreatedAt = time.Now()

		err := models.TopicCommentService.Create(db, t)
		if err != nil {
			c.Data["json"] = map[string]interface{}{"status": 1, "msg": err.Error()}
		} else {
			// 话题评论数+1
			models.TopicService.CountAdd(db, t.TopicId, "comment_count")
			topic := models.TopicService.Get(db, t.TopicId)
			// 用户回帖数+1
			models.UserService.CountAdd(db, c.userId, "comment_num", 1)

			//发送站内信
			m := models.Message{}
			m.FromId = c.userId
			m.UserId = topic.UserId
			m.Content = fmt.Sprintf("<a href=\"/user/home/%d.html\">神秘人</a> 回答了您的求解《<a href=\"/column/detail/%d.html\">%s</a>》",
				c.userId, topic.Id, topic.Title)
			m.ObjType = 2
			m.ObjId = topic.Id
			m.CreatedAt = time.Now()

			models.MessageService.Create(db, &m)

			c.Data["json"] = map[string]interface{}{"status": 0, "msg": "回复成功！"}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"status": 1, "msg": "请先登录", "action": "/user/login.html"}
	}
	c.ServeJSON()
	c.StopRun()
}

func (c *MainController) ZanDo() {
	//收藏判断顺便更新帖子阅读量
	if c.userId < 0 {
		c.Data["json"] = map[string]interface{}{"status": 1, "msg": "请登录~"}
		c.ServeJSON()
		c.StopRun()
	}
	ok, _ := c.GetBool("ok")
	commentId, _ := c.GetInt64("id")
	num := -1
	if ok == false {
		num = 1
	}
	models.TopicCommentService.LikeAdd(db, commentId, num)

	//点赞记录 其他逻辑
	if ok == false {
		var t = models.CommentLike{CommentId: commentId, UserId: c.userId, CreatedAt: time.Now()}
		_ = models.CommentLikeService.Create(db, &t)
	} else {
		//取消赞记录
		models.CommentLikeService.Delete(db, commentId, c.userId)
	}

	c.Data["json"] = map[string]interface{}{"status": 0, "data": ""}
	c.ServeJSON()
	c.StopRun()
}
