package controllers

type OtherController struct {
	BaseController
}

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "other/404.html"
}

func (c *OtherController) Case() {
	c.TplName = "case/case.html"
}

func (c *OtherController) Tips() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "other/tips.html"
}

func (c *OtherController) Notice() {
	c.TplName = "other/notice.html"
}
