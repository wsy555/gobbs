package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"yehelaoren/models"
	_ "yehelaoren/routers"
)

func initDB() {
	db, err := models.OpenDB()
	if err != nil {
		panic(err.Error())
	}

	//数据库没有表自动创建 第一次执行之后可以删除
	if !db.HasTable("topic") {
		db.AutoMigrate(models.Models...)
		//初始化分类
		var Category = [...]string{"提问", "分享", "建议", "公告"}
		sql := "INSERT INTO `category` (`id`,`name`,`description`,`status`,`created_at`) VALUES "
		// 循环Category数组,组合sql语句
		for key, value := range Category {
			if len(Category)-1 == key {
				//最后一条数据 以分号结尾
				sql += fmt.Sprintf("(%d,'%s','%s','%d','%s');", 0, value, value, 1, time.Now().Format("2006-01-02 15:04:05"))
			} else {
				sql += fmt.Sprintf("(%d,'%s','%s','%d','%s'),", 0, value, value, 1, time.Now().Format("2006-01-02 15:04:05"))
			}
		}
		db.Exec(sql)
	}
}

// 评论是否点赞
func GetLike(likeListMap map[int64]bool, commentId int64) string {
	fmt.Println(likeListMap, commentId)
	if likeListMap[commentId] == true {
		return "zanok"
	}
	return ""
}

func main() {
	beego.AddFuncMap("getLike", GetLike)
	beego.Debug()
	beego.Run()
}
