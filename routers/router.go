package routers

import (
	"github.com/astaxie/beego"
	"yehelaoren/controllers"
)

func init() {

	beego.ErrorController(&controllers.ErrorController{})

	beego.Router("/", &controllers.MainController{}, "*:Home")
	beego.Router("/add.html", &controllers.MainController{}, "*:Add")
	beego.Router("/column/:categoryId([\\d]+).html", &controllers.MainController{}, "*:Column")
	beego.Router("/column/detail/:id.html", &controllers.MainController{}, "*:Detail")
	beego.Router("/column/addDo", &controllers.MainController{}, "*:AddDo")
	beego.Router("/jie/reply/", &controllers.MainController{}, "*:ReplyDo")
	beego.Router("/api/zan/", &controllers.MainController{}, "*:ZanDo")
	beego.Router("/top/reply/", &controllers.MainController{}, "*:TopReply")

	beego.Router("/user/home/:id.html", &controllers.UserController{}, "*:Home")
	beego.Router("/user/index.html", &controllers.UserController{}, "*:Index")
	beego.Router("/user/post", &controllers.UserController{}, "*:Post")
	beego.Router("/user/login.html", &controllers.UserController{}, "*:Login")
	beego.Router("/user/loginDo", &controllers.UserController{}, "*:LoginDo")
	beego.Router("/user/reg.html", &controllers.UserController{}, "*:Reg")
	beego.Router("/user/regDo", &controllers.UserController{}, "*:RegDo")
	beego.Router("/user/logout", &controllers.UserController{}, "*:Logout")
	beego.Router("/user/set.html", &controllers.UserController{}, "*:Set")
	beego.Router("/user/setDo", &controllers.UserController{}, "*:SetDo")
	beego.Router("/user/message", &controllers.UserController{}, "*:Message")
	beego.Router("/user/upload/", &controllers.UserController{}, "*:UploadAvatar")
	beego.Router("/message/nums/", &controllers.UserController{}, "*:MessageNums")
	beego.Router("/message/read/", &controllers.UserController{}, "*:MessageRead")
	beego.Router("/sign/in/", &controllers.UserController{}, "*:SignIn")
	beego.Router("/sign/status/", &controllers.UserController{}, "*:SignStatus")
	beego.Router("/top/signin/", &controllers.UserController{}, "*:TopSignIn")
	beego.Router("/collection/find/", &controllers.UserController{}, "*:CollectionFind")
	beego.Router("/collection/add/", &controllers.UserController{}, "*:CollectionAdd")
	beego.Router("/collection/remove/", &controllers.UserController{}, "*:CollectionRemove")

	beego.Router("/liuyao/index", &controllers.LiuyaoController{}, "*:Index")
	beego.Router("/liuyao/format", &controllers.LiuyaoController{}, "*:Format")
	beego.Router("/liuyao/rand_format", &controllers.LiuyaoController{}, "*:RandFormat")
	beego.Router("/liuyao/rand_index", &controllers.LiuyaoController{}, "*:RandIndex")
}
