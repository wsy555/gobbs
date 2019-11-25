package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"strconv"
	"strings"
	"yehelaoren/common"
	"yehelaoren/models"
)

type BaseController struct {
	beego.Controller
	userId   int64
	nickName string
	avatar   string
}

var db = models.GetDB()

func (c *BaseController) Prepare() {
	c.auth()
	c.Layout = "layout.html"
}

func (c *BaseController) auth() {
	arr := strings.Split(c.Ctx.GetCookie("auth"), "|")
	c.userId = -1
	c.nickName = "游客"
	c.avatar = "/static/images/default.png"
	if len(arr) == 2 {
		idStr, password := arr[0], arr[1]
		userId, _ := strconv.ParseInt(idStr, 10, 64)
		if userId > 0 {
			user := models.UserService.Get(db, userId)

			if user != nil && password == common.Md5Encode((c.getClientIp()+"|"+user.Password)) {
				c.userId = user.Id

				if user.Nickname == "" {
					c.nickName = user.Email
				} else {
					c.nickName = user.Nickname
				}
				c.avatar = user.Avatar
			}
		}
	}
	c.Data["loginUserId"] = c.userId
	c.Data["loginNickName"] = c.nickName
	c.Data["loginAvatar"] = c.avatar
}

//获取用户IP地址
func (c *BaseController) getClientIp() string {
	s := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

// 上传图片
func (c *BaseController) UploadFile(filename string, filepath string) (r map[string]interface{}) {
	f, h, err := c.GetFile(filename)

	out := make(map[string]interface{})
	if err != nil {
		out["msg"] = "文件读取错误"
	}
	var fileSuffix, newFile string

	fileSuffix = path.Ext(h.Filename) // 获取文件后缀
	newFile = common.GetRandomString(8) + fileSuffix
	err = c.SaveToFile(filename, filepath+newFile)
	if err != nil {
		out["msg"] = "文件保存错误"
	} else {
		out["status"] = 0
		out["url"] = "/" + filepath + newFile
		out["title"] = newFile
		out["original"] = h.Filename
		out["size"] = h.Size
		out["msg"] = "ok"
	}
	defer f.Close()
	return out
}
