package controllers

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"math/rand"
	"os"
	"strconv"
	"time"
	"yehelaoren/common"
	"yehelaoren/models"
)

type UserController struct {
	BaseController
}

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	// 设置验证码长度
	cpt.ChallengeNums = 4
	// 设置验证码模板高度
	cpt.StdHeight = 50
	// 设置验证码模板宽度
	cpt.StdWidth = 120
}

func (c *UserController) Index() {
	user := models.UserService.Get(db, c.userId)
	if user == nil {
		c.redirect(c.URLFor("UserController.Login"))
	}
	c.Data["user"] = user

	// 我的帖子
	topicWheres := make(map[string]interface{})
	topicWheres["user_id"] = c.userId
	rtList := models.TopicService.Find(db, topicWheres, 0)

	// 我的收藏
	favWheres := make(map[string]interface{})
	favWheres["user_id"] = c.userId
	favList := models.CollectionService.Find(db, favWheres, 0)

	// 帖子总数
	wheres := make(map[string]interface{})
	wheres["user_id"] = c.userId
	totalTopic := models.TopicService.Count(db, wheres)

	// 我的收藏总数
	wheresFav := make(map[string]interface{})
	wheresFav["user_id"] = c.userId
	totalFavorite := models.CollectionService.Count(db, wheresFav)

	c.Data["totalTopic"] = totalTopic
	c.Data["totalFavorite"] = totalFavorite
	c.Data["topicList"] = rtList
	c.Data["favoriteList"] = favList

	c.TplName = "user/index.html"
}

func (c *UserController) Home() {
	userId, _ := c.GetInt64(":id")

	user := models.UserService.Get(db, userId)
	if user == nil {
		c.redirect(c.URLFor("MainController.Home"))
	}
	c.Data["user"] = user

	//他的帖子
	topicWheres := make(map[string]interface{})
	topicWheres["user_id"] = userId
	rtList := models.TopicService.Find(db, topicWheres, 0)
	c.Data["topicList"] = rtList
	//他的回贴
	topicCommentWheres := make(map[string]interface{})
	topicCommentWheres["user_id"] = userId
	commentList := models.TopicCommentService.Find(db, topicCommentWheres, 0)
	c.Data["commentList"] = commentList

	//回复关联帖子名称
	topicIdArr := make([]int64, 0)
	for _, v := range commentList {
		topicIdArr = append(topicIdArr, v.TopicId)
	}
	topicList := models.TopicService.FindIn(db, topicIdArr, 0, 10)
	topicListMap := make(map[int64]string)
	for _, v := range topicList {
		topicListMap[v.Id] = v.Title
	}
	c.Data["topicListMap"] = topicListMap

	c.TplName = "user/home.html"
}

func (c *UserController) Reg() {
	c.TplName = "user/reg.html"
}

// 注册do
func (c *UserController) RegDo() {

	if !cpt.VerifyReq(c.Ctx.Request) {
		c.Data["json"] = map[string]interface{}{"status": 1, "msg": "验证码错误,刷新重新输入！"}
		c.ServeJSON()
		c.StopRun()
	}

	var params = make(map[string]string, 0)

	params["email"] = c.GetString("email")
	params["mobile"] = c.GetString("mobile")
	params["nickname"] = c.GetString("username")
	params["password"] = c.GetString("pass")
	params["rePassword"] = c.GetString("repass")
	// 随机数
	rand.Seed(time.Now().UnixNano())
	avatarNum := rand.Intn(20)
	params["avatar"] = "/static/images/avatar/" + strconv.Itoa(avatarNum) + ".jpg"

	user, err := models.UserService.DoRegister(db, params)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"status": 1, "msg": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"status": 0, "msg": "注册成功", "action": "/"}

		authkey := common.Md5Encode(c.getClientIp() + "|" + user.Password)
		c.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey, 7*86400)
	}
	c.ServeJSON()
	c.StopRun()
}

func (c *UserController) Login() {
	c.TplName = "user/login.html"
}

// 登录do
func (c *UserController) LoginDo() {
	if !cpt.VerifyReq(c.Ctx.Request) {
		c.Data["json"] = map[string]interface{}{"status": 2, "msg": "验证码错误,刷新重新输入！"}
		c.ServeJSON()
		c.StopRun()
	}

	email := c.GetString("email")
	password := c.GetString("pass")

	user, err := models.UserService.DoLogin(db, email, password)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"status": 1, "msg": err.Error()}
	} else {
		authkey := common.Md5Encode(c.getClientIp() + "|" + user.Password)
		c.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey, 7*86400)
		c.Data["json"] = map[string]interface{}{"status": 0, "msg": "登录成功", "action": "/"}
	}
	c.ServeJSON()
	c.StopRun()
}

func (c *UserController) Logout() {
	c.Ctx.SetCookie("auth", "")
	c.redirect(c.URLFor("MainController.Home"))
}

func (c *UserController) Forget() {
	c.TplName = "user/forget.html"
}

func (c *UserController) Set() {
	//用户信息
	if c.userId < 0 {
		c.redirect(c.URLFor("UserController.Login"))
	}
	user := models.UserService.Get(db, c.userId)
	c.Data["user"] = user
	c.TplName = "user/set.html"
}

func (c *UserController) SetDo() {
	if c.userId < 0 {
		c.Data["json"] = map[string]interface{}{"status": 1, "msg": "请先登录！"}
	} else {
		columns := make(map[string]interface{})

		email := c.GetString("email")
		nickname := c.GetString("nickname")
		description := c.GetString("description")
		avatar := c.GetString("avatar")

		if len(email) > 0 {
			columns["email"] = email
		}
		if len(nickname) > 0 {
			columns["nickname"] = nickname
		}
		if len(description) > 0 {
			columns["description"] = description
		}
		if len(avatar) > 0 {
			columns["avatar"] = avatar
		}
		c.Data["json"] = map[string]interface{}{"status": 1, "msg": "参数错误"}
		if len(columns) > 0 {
			wheres := make(map[string]interface{})
			wheres["id"] = c.userId
			_ = models.UserService.Updates(db, wheres, columns)
			c.Data["json"] = map[string]interface{}{"status": 0, "msg": "ok"}
		}
	}

	c.ServeJSON()
	c.StopRun()
}

func (c *UserController) Message() {
	if c.userId < 0 {
		c.redirect(c.URLFor("UserController.Login"))
	}
	wheres := make(map[string]interface{})
	wheres["user_id"] = c.userId
	messages := models.MessageService.Find(db, wheres, 0)
	c.Data["messages"] = messages
	c.TplName = "user/message.html"
}

func (c *UserController) UploadAvatar() {
	if c.userId < 0 {
		c.Data["json"] = map[string]interface{}{"status": 1, "msg": "请先登录！"}
	} else {
		filepath := "static/upload/" + time.Now().Format("20060102") + "/"
		_, err := os.Stat(filepath)
		if err != nil {
			_ = os.MkdirAll(filepath, 0777)
		}
		rt := c.UploadFile("up_file", filepath)

		c.Data["json"] = rt
	}

	c.ServeJSON()
	c.StopRun()
}

func (c *UserController) MessageNums() {
	wheres := make(map[string]interface{})
	wheres["user_id"] = c.userId
	wheres["status"] = 0

	count := models.MessageService.Count(db, wheres)
	c.Data["json"] = map[string]interface{}{"status": 0, "count": count}
	c.ServeJSON()
	c.StopRun()
}

func (c *UserController) MessageRead() {
	wheres := make(map[string]interface{})
	wheres["user_id"] = c.userId
	wheres["status"] = 0

	columns := make(map[string]interface{})
	columns["status"] = 1

	_ = models.MessageService.Updates(db, wheres, columns)

	c.Data["json"] = map[string]interface{}{"status": 0}
	c.ServeJSON()
	c.StopRun()
}

// 签到
func (c *UserController) SignIn() {
	if c.userId < 0 {
		c.Data["json"] = map[string]interface{}{"status": 2}
		c.ServeJSON()
		c.StopRun()
	}
	wheres := make(map[string]interface{})
	wheres["user_id"] = c.userId

	sign := models.SignInService.Get(db, wheres)
	var days int
	day := time.Now().Format("20060102")

	if sign != nil && sign.Day == day {
		c.Data["json"] = map[string]interface{}{"status": 0, "data": ""}
		c.ServeJSON()
		c.StopRun()
	} else {
		days = sign.Days + 1
	}

	t := models.SignIn{}
	t.UserId = c.userId
	t.Experience = sign.Experience + 5
	t.Days = days
	t.Day = day
	t.CreatedAt = time.Now()

	err := models.SignInService.Create(db, &t)
	rtData := make(map[string]interface{})
	if err == nil {
		rtData["days"] = days
		rtData["experience"] = t.Experience
		rtData["signed"] = true
	}
	c.Data["json"] = map[string]interface{}{"status": 0, "data": rtData}
	c.ServeJSON()
	c.StopRun()
}

func (c *UserController) SignStatus() {
	if c.userId < 0 {
		c.Data["json"] = map[string]interface{}{"status": 2}
		c.ServeJSON()
		c.StopRun()
	}

	day := time.Now().Format("20060102")

	wheres := make(map[string]interface{})
	wheres["user_id"] = c.userId
	//最后一次签到记录
	sign := models.SignInService.Get(db, wheres)

	rtData := make(map[string]interface{})
	rtData["signed"] = false
	rtData["experience"] = 0
	rtData["days"] = 0
	//今天已签到
	if sign != nil && day == sign.Day {
		rtData["signed"] = true
		rtData["experience"] = sign.Experience
		rtData["days"] = sign.Days
	}

	c.Data["json"] = map[string]interface{}{"status": 0, "data": rtData}
	c.ServeJSON()
	c.StopRun()
}

func (c *UserController) TopSignIn() {

	day := time.Now().Format("20060102")

	wheres := make(map[string]interface{})
	wheres["user_id"] = c.userId
	//最后一次签到记录
	sign := models.SignInService.Get(db, wheres)

	rtData := make(map[string]interface{})
	rtData["signed"] = false
	rtData["experience"] = 0
	rtData["days"] = 0
	//今天已签到
	if sign != nil && day == sign.Day {
		rtData["signed"] = true
		rtData["experience"] = sign.Experience
		rtData["days"] = sign.Days
	}

	c.Data["json"] = map[string]interface{}{"status": 0, "data": rtData}
	c.ServeJSON()
	c.StopRun()
}

// 收藏查找
func (c *UserController) CollectionFind() {
	//收藏判断顺便更新帖子阅读量 阅读数+1
	topicId, _ := c.GetInt64("id")
	models.TopicService.CountAdd(db, topicId, "view_count")

	wheres := make(map[string]interface{})
	wheres["user_id"] = c.userId
	wheres["topic_id"] = topicId
	favorite := models.CollectionService.Get(db, wheres)
	collection := false
	if favorite.Id > 0 {
		collection = true
	}

	c.Data["json"] = map[string]interface{}{"status": 0, "data": map[string]interface{}{"collection": collection}}
	c.ServeJSON()
	c.StopRun()
}

func (c *UserController) CollectionRemove() {
	if c.userId < 0 {
		c.Data["json"] = map[string]interface{}{"status": 2, "msg": "请先登录"}
	} else {
		topicId, _ := c.GetInt64("topic_id")
		models.CollectionService.Delete(db, c.userId, topicId)
		c.Data["json"] = map[string]interface{}{"status": 0, "data": map[string]interface{}{"collection": false}}
	}
	c.ServeJSON()
	c.StopRun()
}

// 收藏
func (c *UserController) CollectionAdd() {
	if c.userId < 0 {
		c.Data["json"] = map[string]interface{}{"status": 2, "msg": "请先登录"}
	} else {
		// 收藏判断顺便更新帖子阅读量
		topicId, _ := c.GetInt64("topic_id")

		// 判断帖子是否存在
		topic := models.TopicService.Get(db, topicId)
		if topic.Title == "" {
			c.Data["json"] = map[string]interface{}{"status": 2, "msg": "帖子已经被删除"}
		} else {
			t := models.Favorite{}
			t.UserId = c.userId
			t.TopicId = topicId
			t.Title = topic.Title
			t.CreatedAt = time.Now()

			models.CollectionService.Create(db, &t)
			c.Data["json"] = map[string]interface{}{"status": 0, "data": map[string]interface{}{"collection": true}}
		}
	}
	c.ServeJSON()
	c.StopRun()
}

func (c *UserController) Activate() {
	c.TplName = "user/activate.html"
}
