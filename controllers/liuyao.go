/**********************************************
** @Des: 文章
** @Author: wangsy
** @Date:   2017-12-09 14:17:37
***********************************************/
package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"
	"yehelaoren/models"
)

type LiuyaoController struct {
	BaseController
}

type Yao struct {
	name   string
	detail map[int]string
	fushen map[int]string
}

// 乾宫八卦俱属金
var qianweitian = Yao{
	name: "乾为天",
	detail: map[int]string{
		6: "父母戌土、世",
		5: "兄弟申金、",
		4: "官鬼午火、",
		3: "父母辰土、应",
		2: "妻财寅木、",
		1: "子孙子水、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var tianfenggou = Yao{
	name: "天风姤",
	detail: map[int]string{
		6: "父母戌土、",
		5: "兄弟申金、",
		4: "官鬼午火、应",
		3: "兄弟酉金、",
		2: "子孙亥水、",
		1: "父母丑土、、世",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "妻财寅木",
		1: "",
	},
}
var tianshandun = Yao{
	name: "天山遁",
	detail: map[int]string{
		6: "父母戌土、",
		5: "兄弟申金、应",
		4: "官鬼午火、",
		3: "兄弟申金、",
		2: "官鬼午火、、世",
		1: "父母辰土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "妻财寅木",
		1: "子孙子水",
	},
}
var tiandipi = Yao{
	name: "天地否",
	detail: map[int]string{
		6: "父母戌土、应",
		5: "兄弟申金、",
		4: "官鬼午火、",
		3: "妻財卯木、、世",
		2: "官鬼巳火、、",
		1: "父母未土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "子孙子水",
	},
}
var fengdiguan = Yao{
	name: "风地观",
	detail: map[int]string{
		6: "妻財卯木、",
		5: "官鬼巳火、",
		4: "父母未土、、世",
		3: "妻財卯木、、",
		2: "官鬼巳火、、",
		1: "父母未土、、应",
	},
	fushen: map[int]string{
		6: "",
		5: "兄弟申金",
		4: "",
		3: "",
		2: "",
		1: "子孙子水",
	},
}
var shandibo = Yao{
	name: "山地剥",
	detail: map[int]string{
		6: "妻財寅木、",
		5: "子孫子水、、世",
		4: "父母戌土、、",
		3: "妻財卯木、、",
		2: "官鬼巳火、、应",
		1: "父母未土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "兄弟申金",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var huodijin = Yao{
	name: "火地晋",
	detail: map[int]string{
		6: "官鬼巳火、",
		5: "父母未土、、",
		4: "兄弟酉金、世",
		3: "妻財卯木、、",
		2: "官鬼巳火、、",
		1: "父母未土、、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "子孙子水",
	},
}
var huotiandayou = Yao{
	name: "火天大有",
	detail: map[int]string{
		6: "官鬼巳火、应",
		5: "父母未土、、",
		4: "兄弟酉金、",
		3: "父母辰土、世",
		2: "妻財寅木、",
		1: "子孫子水、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}

// 坎宫八卦俱属水
var kanweishui = Yao{
	name: "坎为水",
	detail: map[int]string{
		6: "兄弟子水、、世",
		5: "官鬼戌土、",
		4: "父母申金、、",
		3: "妻財午火、、应",
		2: "官鬼辰土、",
		1: "子孫寅木、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var shuizejie = Yao{
	name: "水泽节",
	detail: map[int]string{
		6: "兄弟子水、、",
		5: "官鬼戌土、",
		4: "父母申金、、",
		3: "官鬼丑土、、",
		2: "子孫卯木、",
		1: "妻財巳火、世",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var shuileitun = Yao{
	name: "水雷屯",
	detail: map[int]string{
		6: "兄弟子水、、",
		5: "官鬼戌土、应",
		4: "父母申金、、",
		3: "官鬼辰土、、",
		2: "子孫寅木、、世",
		1: "兄弟子水、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "妻财午火",
		2: "",
		1: "",
	},
}
var shuihuojiji = Yao{
	name: "水火既济",
	detail: map[int]string{
		6: "兄弟子水、、应",
		5: "官鬼戌土、",
		4: "父母申金、、",
		3: "兄弟亥水、世",
		2: "官鬼丑土、、",
		1: "子孫卯木、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "妻财午火",
		2: "",
		1: "",
	},
}
var zehuoge = Yao{
	name: "泽火革",
	detail: map[int]string{
		6: "官鬼未土、、",
		5: "父母酉金、",
		4: "兄弟亥水、世",
		3: "兄弟亥水、",
		2: "官鬼丑土、、",
		1: "子孫卯木、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "妻财午火",
		2: "",
		1: "",
	},
}
var leihuofeng = Yao{
	name: "雷火丰",
	detail: map[int]string{
		6: "官鬼戌土、、",
		5: "父母申金、、世",
		4: "妻財午火、",
		3: "兄弟亥水、",
		2: "官鬼丑土、、应",
		1: "子孫卯木、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var dihuomignyi = Yao{
	name: "地火明夷",
	detail: map[int]string{
		6: "父母酉金、、",
		5: "兄弟亥水、、",
		4: "官鬼丑土、、世",
		3: "兄弟亥水、",
		2: "官鬼丑上、、",
		1: "子孫卯木、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "妻财午火",
		2: "",
		1: "",
	},
}
var dishuishi = Yao{
	name: "地水师",
	detail: map[int]string{
		6: "父母酉金、、应",
		5: "兄弟亥水、、",
		4: "官鬼丑士、、",
		3: "妻財午火、、世",
		2: "官鬼辰土、",
		1: "子孫寅木、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}

// 艮宫八卦
var genweishan = Yao{
	name: "艮为山",
	detail: map[int]string{
		6: "官鬼寅木、世",
		5: "妻財子水、、",
		4: "兄弟戌土、、",
		3: "子孫申金、应",
		2: "父母午火、、",
		1: "兄弟辰土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var shanhuobi = Yao{
	name: "山火贲",
	detail: map[int]string{
		6: "官鬼寅木、",
		5: "妻財子水、、",
		4: "兄弟戌土、、应",
		3: "妻財亥水、",
		2: "兄弟丑土、、",
		1: "官鬼卯木、世",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "子孙申金",
		2: "父母午火",
		1: "",
	},
}
var shantiandaxu = Yao{
	name: "山天大蓄",
	detail: map[int]string{
		6: "官鬼寅木、",
		5: "妻財子水、、应",
		4: "兄弟戌土、、",
		3: "兄弟辰土、",
		2: "官鬼寅木、世",
		1: "妻財子水、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "子孙申金",
		2: "父母午火",
		1: "",
	},
}
var shanzesun = Yao{
	name: "山泽损",
	detail: map[int]string{
		6: "官鬼寅木、应",
		5: "妻財子水、、",
		4: "兄弟戌土、、",
		3: "兄弟丑土、、世",
		2: "官鬼卯木、",
		1: "父母巳火、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "子孙申金",
		2: "",
		1: "",
	},
}
var huozekui = Yao{
	name: "火泽睽",
	detail: map[int]string{
		6: "父母巳火、",
		5: "兄弟未土、、",
		4: "子孫酉金、世",
		3: "兄弟丑土、、",
		2: "官鬼卯木、",
		1: "父母巳火、应",
	},
	fushen: map[int]string{
		6: "",
		5: "妻财子水",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var tianzelv = Yao{
	name: "天泽履",
	detail: map[int]string{
		6: "兄弟戌土、",
		5: "子孫申金、世",
		4: "父母午火、",
		3: "兄弟丑土、、",
		2: "官鬼卯木、应",
		1: "父母巳火、",
	},
	fushen: map[int]string{
		6: "",
		5: "妻财子水",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var fengzezhongfu = Yao{
	name: "风泽中孚",
	detail: map[int]string{
		6: "官鬼卯木、",
		5: "父母巳火、",
		4: "兄弟未土、、世",
		3: "兄弟丑土、、",
		2: "官鬼卯木、",
		1: "父母巳火、应",
	},
	fushen: map[int]string{
		6: "",
		5: "妻财子水",
		4: "",
		3: "子孙申金",
		2: "",
		1: "",
	},
}
var fengshanjian = Yao{
	name: "风山渐",
	detail: map[int]string{
		6: "官鬼卯木、应",
		5: "父母巳火、",
		4: "兄弟未土、、",
		3: "子孫申金、世",
		2: "父母午火、、",
		1: "兄弟辰土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "妻财子水",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}

// 震宫八卦
var zhenweilei = Yao{
	name: "震为雷",
	detail: map[int]string{
		6: "妻財戌士、、世",
		5: "官鬼申金、、",
		4: "子孫午火、",
		3: "妻財辰土、、应",
		2: "兄弟寅木、、",
		1: "父母子水、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var leidiyu = Yao{
	name: "雷地豫",
	detail: map[int]string{
		6: "妻財戌土、、",
		5: "官鬼申金、、",
		4: "子孫午火、应",
		3: "兄弟卯木、、",
		2: "子孫巳火、、",
		1: "妻財未土、、世",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "父母子水",
	},
}
var leishuijie = Yao{
	name: "雷水解",
	detail: map[int]string{
		6: "妻財戌土、、",
		5: "官鬼申金、、应",
		4: "子孫午火、",
		3: "子孫午火、、",
		2: "妻財辰士、世",
		1: "兄弟寅木、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "父母子水",
	},
}
var leifengheng = Yao{
	name: "雷风恒",
	detail: map[int]string{
		6: "妻財戌土、、应",
		5: "官鬼申金、、",
		4: "子孫午火、",
		3: "官鬼酉金、世",
		2: "父母亥水、",
		1: "妻財丑土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "兄弟寅木",
		1: "",
	},
}
var difengsheng = Yao{
	name: "地风升",
	detail: map[int]string{
		6: "官鬼酉金、、",
		5: "父母亥水、、",
		4: "妻財丑土、、世",
		3: "官鬼酉金、",
		2: "父母亥水、",
		1: "妻財丑土、、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "子孙午火",
		3: "",
		2: "兄弟寅木",
		1: "",
	},
}
var shuifengjing = Yao{
	name: "水风井",
	detail: map[int]string{
		6: "父母子水、、",
		5: "妻財戌土、世",
		4: "官鬼申金、、",
		3: "官鬼酉金、",
		2: "父母亥水、应",
		1: "妻財丑土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "子孙午火",
		3: "",
		2: "兄弟寅木",
		1: "",
	},
}
var zefengdaguo = Yao{
	name: "泽风大过",
	detail: map[int]string{
		6: "妻財未土、、",
		5: "官鬼酉金、",
		4: "父母亥水、世",
		3: "官鬼酉金、",
		2: "父母亥水、",
		1: "妻財丑土、、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "子孙午火",
		3: "",
		2: "兄弟寅木",
		1: "",
	},
}
var zeleisui = Yao{
	name: "泽雷随",
	detail: map[int]string{
		6: "妻財未土、、应",
		5: "官鬼酉金、",
		4: "父母亥水、",
		3: "妻財辰士、、世",
		2: "兄弟寅木、、",
		1: "父母子水、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "子孙午火",
		3: "",
		2: "",
		1: "",
	},
}

// 巽宫八卦
var xunweifeng = Yao{
	name: "巽为风",
	detail: map[int]string{
		6: "兄弟卯木、世",
		5: "子孫巳火、",
		4: "妻財未土、、",
		3: "官鬼酉金、应",
		2: "父母亥水、",
		1: "妻財丑土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var fengtianxiaoxu = Yao{
	name: "风天小畜",
	detail: map[int]string{
		6: "兄弟卯木、",
		5: "子孫巳火、",
		4: "妻財未土、、应",
		3: "妻財辰土、",
		2: "兄弟寅木、",
		1: "父母子水、世",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "官鬼酉金",
		2: "",
		1: "",
	},
}
var fenghuojiaren = Yao{
	name: "风火家人",
	detail: map[int]string{
		6: "兄弟卯木、",
		5: "子孫巳火、应",
		4: "妻財未土、、",
		3: "父母亥水、",
		2: "妻財丑土、、世",
		1: "兄弟卯木、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "官鬼酉金",
		2: "",
		1: "",
	},
}
var fengleiyi = Yao{
	name: "风雷益",
	detail: map[int]string{
		6: "兄弟卯木、应",
		5: "子孫巳火、",
		4: "妻財未土、、",
		3: "妻財辰土、、世",
		2: "兄弟寅木、、",
		1: "父母子水、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "官鬼酉金",
		2: "",
		1: "",
	},
}
var tianleiwuwang = Yao{
	name: "天雷无妄",
	detail: map[int]string{
		6: "妻財戌土、",
		5: "官鬼申金、",
		4: "子孫午火、世",
		3: "妻財辰土、、",
		2: "兄弟寅木、、",
		1: "父母子水、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var huoleishike = Yao{
	name: "火雷噬嗑",
	detail: map[int]string{
		6: "子孫巳火、",
		5: "妻財未土、、世",
		4: "官鬼酉金、",
		3: "妻財辰土、、",
		2: "兄弟寅木、、应",
		1: "父母子水、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var shanleiyi = Yao{
	name: "山雷颐",
	detail: map[int]string{
		6: "兄弟寅木、",
		5: "父母子水、、",
		4: "妻財戌土、、世",
		3: "妻財辰土、、",
		2: "兄弟寅木、、",
		1: "父母子水、应",
	},
	fushen: map[int]string{
		6: "",
		5: "子孙巳火",
		4: "",
		3: "官鬼酉金",
		2: "",
		1: "",
	},
}
var shanfenggu = Yao{
	name: "山风蛊",
	detail: map[int]string{
		6: "兄弟寅木、应",
		5: "父母子水、、",
		4: "妻財戌土、、",
		3: "官鬼酉金、世",
		2: "父母亥水、",
		1: "妻財丑土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "子孙巳火",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}

// 离宫八卦
var liweihuo = Yao{
	name: "离为火",
	detail: map[int]string{
		6: "兄弟巳火、世",
		5: "子孫未土、、",
		4: "妻財酉金、",
		3: "官鬼亥水、应",
		2: "子孫丑土、、",
		1: "父母卯木、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var huoshanlv = Yao{
	name: "火山旅",
	detail: map[int]string{
		6: "兄弟巳火、",
		5: "子孫未土、、",
		4: "妻財酉金、应",
		3: "妻財申金、",
		2: "兄弟午火、、",
		1: "子孫辰土、、世",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "官鬼亥水",
		2: "",
		1: "父母卯木",
	},
}
var huofengding = Yao{
	name: "火风鼎",
	detail: map[int]string{
		6: "兄弟巳火、",
		5: "子孫未土、、",
		4: "妻財酉金、应",
		3: "妻財申金、",
		2: "兄弟午火、、",
		1: "子孫辰土、、世",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "官鬼亥水",
		2: "",
		1: "父母卯木",
	},
}
var huoshuiweiji = Yao{
	name: "火水未济",
	detail: map[int]string{
		6: "兄弟巳火、应",
		5: "子孫未土、、",
		4: "妻財酉金、",
		3: "兄弟午火、、世",
		2: "子孫辰土、",
		1: "父母寅木、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "官鬼亥水",
		2: "",
		1: "",
	},
}
var shanshuimeng = Yao{
	name: "山水蒙",
	detail: map[int]string{
		6: "父母寅木、",
		5: "官鬼子水、、",
		4: "子孫戌土、、世",
		3: "兄弟午火、、",
		2: "子孫辰土、",
		1: "父母寅木、、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "妻财酉金",
		3: "",
		2: "",
		1: "",
	},
}
var fengshuihuan = Yao{
	name: "风水涣",
	detail: map[int]string{
		6: "父母卯木、",
		5: "兄弟巳火、世",
		4: "子孫未土、、",
		3: "兄弟午火、、",
		2: "子孫辰土、应",
		1: "父母寅木、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "妻财酉金",
		3: "官鬼亥水",
		2: "",
		1: "",
	},
}
var tianshuisong = Yao{
	name: "天水讼",
	detail: map[int]string{
		6: "子孫戌士、",
		5: "妻財申金、",
		4: "兄弟午火、世",
		3: "兄弟午火、、",
		2: "子孫辰土、",
		1: "父母寅木、、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "官鬼亥水",
		2: "",
		1: "",
	},
}
var tianhuotongren = Yao{
	name: "天火同人",
	detail: map[int]string{
		6: "子孫戌土、应",
		5: "妻財申金、",
		4: "兄弟午火、",
		3: "官鬼亥水、世",
		2: "子孫丑土、、",
		1: "父母卯木、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}

// 坤宫八卦
var kunweidi = Yao{
	name: "坤为地",
	detail: map[int]string{
		6: "子孫酉金、、世",
		5: "妻財亥水、、",
		4: "兄弟丑土、、",
		3: "官鬼卯木、、应",
		2: "父母巳火、、",
		1: "兄弟未土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var dileifu = Yao{
	name: "地雷复",
	detail: map[int]string{
		6: "子孫酉金、、",
		5: "妻財亥水、、",
		4: "兄弟丑土、、应",
		3: "兄弟辰土、、",
		2: "官鬼寅木、、",
		1: "妻財子水、世",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "父母巳火",
		1: "",
	},
}
var dizelin = Yao{
	name: "地泽临",
	detail: map[int]string{
		6: "子孫酉金、、",
		5: "妻財亥水、、应",
		4: "兄弟丑土、、",
		3: "兄弟丑土、、",
		2: "官鬼卯木、世",
		1: "父母巳火、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var ditiantai = Yao{
	name: "地天泰",
	detail: map[int]string{
		6: "子孫酉金、、应",
		5: "妻財亥水、、",
		4: "兄弟丑土、、",
		3: "兄弟辰士、世",
		2: "官鬼寅木、",
		1: "妻財子水、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "父母巳火",
		1: "",
	},
}
var leitiandazhuang = Yao{
	name: "雷天大壮",
	detail: map[int]string{
		6: "兄弟戌土、、",
		5: "子孫申金、、",
		4: "父母午火、世",
		3: "兄弟辰土、",
		2: "官鬼寅木、",
		1: "妻財子水、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var zetianguai = Yao{
	name: "择天夬",
	detail: map[int]string{
		6: "兄弟未土、、",
		5: "子孫酉金、世",
		4: "妻財亥水、",
		3: "兄弟辰土、",
		2: "官鬼寅木、应",
		1: "妻財子水、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "父母巳火",
		1: "",
	},
}
var shuitianxu = Yao{
	name: "水天需",
	detail: map[int]string{
		6: "妻財子水、、",
		5: "兄弟戌土、",
		4: "子孫申金、、世",
		3: "兄弟辰土、",
		2: "官鬼寅木、",
		1: "妻財子水、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "父母巳火",
		1: "",
	},
}
var shuidibi = Yao{
	name: "水地比",
	detail: map[int]string{
		6: "妻財子水、、应",
		5: "兄弟戌土、",
		4: "子孫申金、、",
		3: "官鬼卯木、、世",
		2: "父母巳火、、",
		1: "兄弟未土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}

// 兑宫八卦
var duiweize = Yao{
	name: "兑为泽",
	detail: map[int]string{
		6: "父母未土、、世",
		5: "兄弟酉金、",
		4: "子孫亥水、",
		3: "父母丑土、、应",
		2: "妻財卯木、",
		1: "官鬼巳火、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var zeshuikun = Yao{
	name: "泽水困",
	detail: map[int]string{
		6: "父母未土、、",
		5: "兄弟酉金、",
		4: "子孫亥水、应",
		3: "官鬼午火、、",
		2: "父母辰土、",
		1: "妻財卯木、、世",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var zedicui = Yao{
	name: "泽地萃",
	detail: map[int]string{
		6: "父母未土、、",
		5: "兄弟酉金、应",
		4: "子孫亥水、",
		3: "妻財卯木、、",
		2: "官鬼巳火、、世",
		1: "父母未土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "",
		1: "",
	},
}
var zeshanxian = Yao{
	name: "泽山咸",
	detail: map[int]string{
		6: "父母未土、、应",
		5: "兄弟酉金、",
		4: "子孫亥水、",
		3: "兄弟申金、世",
		2: "官鬼午火、、",
		1: "父母辰土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "妻财卯木",
		1: "",
	},
}
var shuishanjian = Yao{
	name: "水山蹇",
	detail: map[int]string{
		6: "子孫子水、、",
		5: "父母戌土、",
		4: "兄弟申金、、世",
		3: "兄弟申金、",
		2: "官鬼午火、、",
		1: "父母辰土、、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "妻财卯木",
		1: "",
	},
}
var dishanqian = Yao{
	name: "地山谦",
	detail: map[int]string{
		6: "兄弟酉金、、",
		5: "子孫亥水、、世",
		4: "父母丑土、、",
		3: "兄弟申金、",
		2: "官鬼午火、、应",
		1: "父母辰土、、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "",
		3: "",
		2: "妻财卯木",
		1: "",
	},
}
var leishanxiaoguo = Yao{
	name: "雷山小过",
	detail: map[int]string{
		6: "父母戌土、、",
		5: "兄弟申金、、",
		4: "官鬼午火、世",
		3: "兄弟申金、",
		2: "官鬼午火、、",
		1: "父母辰土、、应",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "子孙亥水",
		3: "",
		2: "妻财卯木",
		1: "",
	},
}
var leizeguimei = Yao{
	name: "雷泽归妹",
	detail: map[int]string{
		6: "父母戌土、、应",
		5: "兄弟申金、、",
		4: "官鬼午火、",
		3: "父母丑土、、世",
		2: "妻財卯木、",
		1: "官鬼巳火、",
	},
	fushen: map[int]string{
		6: "",
		5: "",
		4: "子孙亥水",
		3: "",
		2: "",
		1: "",
	},
}

// 农历请求地址
const LunarApi = "https://www.sojson.com/open/api/lunar/json.shtml?date="

var BaGua = map[int]interface{}{
	111: "乾",
	211: "巽",
	121: "离",
	112: "兑",
	222: "坤",
	122: "震",
	212: "坎",
	221: "艮",
}
var tianganNumMap = map[string]int{
	"甲": 1,
	"乙": 2,
	"丙": 3,
	"丁": 4,
	"戊": 5,
	"己": 6,
	"庚": 7,
	"辛": 8,
	"壬": 9,
	"癸": 10,
}
var dizhiNumMap = map[string]int{
	"子": 1,
	"丑": 2,
	"寅": 3,
	"卯": 4,
	"辰": 5,
	"巳": 6,
	"午": 7,
	"未": 8,
	"申": 9,
	"酉": 10,
	"戌": 11,
	"亥": 12,
}
var kongwangMap = map[int]string{
	1:  "子丑",
	2:  "子丑",
	3:  "寅卯",
	4:  "寅卯",
	5:  "辰巳",
	6:  "辰巳",
	7:  "午未",
	8:  "午未",
	9:  "申酉",
	10: "申酉",
	11: "戌亥",
	12: "戌亥",
}
var LiuShenStr = "青龙朱雀勾陈腾蛇白虎玄武"
var LiuShenSort = map[string]int{
	"青龙": 1,
	"朱雀": 2,
	"勾陈": 3,
	"腾蛇": 4,
	"白虎": 5,
	"玄武": 6,
}
var LiuShen = map[string]string{
	"甲": "青龙",
	"乙": "青龙",
	"丙": "朱雀",
	"丁": "朱雀",
	"戊": "勾陈",
	"己": "腾蛇",
	"庚": "白虎",
	"辛": "白虎",
	"壬": "玄武",
	"癸": "玄武",
}
var GuiRen = map[string]string{
	"甲": "丑,未",
	"戊": "丑,未",
	"乙": "子,申",
	"己": "子,申",
	"丙": "亥,酉",
	"丁": "亥,酉",
	"壬": "卯,巳",
	"癸": "卯,巳",
	"庚": "寅,午",
	"辛": "寅,午",
}
var MaXin = map[string]string{
	"申": "寅",
	"子": "寅",
	"辰": "寅",
	"寅": "巳",
	"午": "巳",
	"戌": "巳",
	"巳": "亥",
	"酉": "亥",
	"丑": "亥",
	"亥": "巳",
	"卯": "巳",
	"未": "巳",
}
var TaoHua = map[string]string{
	"寅": "卯",
	"午": "卯",
	"戌": "卯",
	"申": "酉",
	"子": "酉",
	"辰": "酉",
	"巳": "午",
	"酉": "午",
	"丑": "午",
	"亥": "子",
	"卯": "子",
	"未": "子",
}
var RiLu = map[string]string{
	"甲": "寅",
	"乙": "卯",
	"戊": "巳",
	"己": "午",
	"丙": "巳",
	"丁": "午",
	"壬": "亥",
	"癸": "子",
	"庚": "申",
	"辛": "酉",
}
var YaoWuXingMap = map[string]string{
	"子": "水",
	"亥": "水",
	"辰": "土",
	"戌": "土",
	"丑": "土",
	"巳": "火",
	"午": "火",
	"申": "金",
	"酉": "金",
	"寅": "木",
	"卯": "木",
}
var LiuQinMap = map[string]string{
	"金木": "妻财",
	"金水": "子孙",
	"金金": "兄弟",
	"金火": "官鬼",
	"金土": "父母",

	"木土": "妻财",
	"木火": "子孙",
	"木木": "兄弟",
	"木金": "官鬼",
	"木水": "父母",

	"水火": "妻财",
	"水木": "子孙",
	"水水": "兄弟",
	"水土": "官鬼",
	"水金": "父母",

	"火金": "妻财",
	"火土": "子孙",
	"火火": "兄弟",
	"火水": "官鬼",
	"火木": "父母",

	"土水": "妻财",
	"土金": "子孙",
	"土土": "兄弟",
	"土木": "官鬼",
	"土火": "父母",
}
var Gua64 = map[string]Yao{
	// 乾宫八卦
	"111111": qianweitian,
	"111112": tianfenggou,
	"111122": tianshandun,
	"111222": tiandipi,
	"112222": fengdiguan,
	"122222": shandibo,
	"121222": huodijin,
	"121111": huotiandayou,
	// 坎宫八卦
	"212212": kanweishui,
	"212211": shuizejie,
	"212221": shuileitun,
	"212121": shuihuojiji,
	"211121": zehuoge,
	"221121": leihuofeng,
	"222121": dihuomignyi,
	"222212": dishuishi,
	// 艮宫八卦
	"122122": genweishan,
	"122121": shanhuobi,
	"122111": shantiandaxu,
	"122211": shanzesun,
	"121211": huozekui,
	"111211": tianzelv,
	"112211": fengzezhongfu,
	"112122": fengshanjian,
	// 震宫八卦
	"221221": zhenweilei,
	"221222": leidiyu,
	"221212": leishuijie,
	"221112": leifengheng,
	"222112": difengsheng,
	"212112": shuifengjing,
	"211112": zefengdaguo,
	"211221": zeleisui,
	//巽宫八卦
	"112112": xunweifeng,
	"112111": fengtianxiaoxu,
	"112121": fenghuojiaren,
	"112221": fengleiyi,
	"111221": tianleiwuwang,
	"121221": huoleishike,
	"122221": shanleiyi,
	"122112": shanfenggu,
	// 离宫八卦
	"121121": liweihuo,
	"121122": huoshanlv,
	"121112": huofengding,
	"121212": huoshuiweiji,
	"122212": shanshuimeng,
	"112212": fengshuihuan,
	"111212": tianshuisong,
	"111121": tianhuotongren,
	// 坤宫八卦
	"222222": kunweidi,
	"222221": dileifu,
	"222211": dizelin,
	"222111": ditiantai,
	"221111": leitiandazhuang,
	"211111": zetianguai,
	"212111": shuitianxu,
	"212222": shuidibi,
	// 对宫八卦
	"211211": duiweize,
	"211212": zeshuikun,
	"211222": zedicui,
	"211122": zeshanxian,
	"212122": shuishanjian,
	"222122": dishanqian,
	"221122": leishanxiaoguo,
	"221211": leizeguimei,
}

var GuaWuXinMap = map[string]string{
	// 乾宫八卦
	"111111": "金",
	"111112": "金",
	"111122": "金",
	"111222": "金",
	"112222": "金",
	"122222": "金",
	"121222": "金",
	"121111": "金",
	// 坎宫八卦
	"212212": "水",
	"212211": "水",
	"212221": "水",
	"212121": "水",
	"211121": "水",
	"221121": "水",
	"222121": "水",
	"222212": "水",
	// 艮宫八卦
	"122122": "土",
	"122121": "土",
	"122111": "土",
	"122211": "土",
	"121211": "土",
	"111211": "土",
	"112211": "土",
	"112122": "土",
	// 震宫八卦
	"221221": "木",
	"221222": "木",
	"221212": "木",
	"221112": "木",
	"222112": "木",
	"212112": "木",
	"211112": "木",
	"211221": "木",
	// 巽宫八卦
	"112112": "木",
	"112111": "木",
	"112121": "木",
	"112221": "木",
	"111221": "木",
	"121221": "木",
	"122221": "木",
	"122112": "木",
	// 离宫八卦
	"121121":  "火",
	"121122":  "火",
	"121112":  "火",
	"121212":  "火",
	"122212":  "火",
	"112212":  "火",
	"111212":  "火",
	"1111121": "火",
	// 坤宫八卦
	"222222": "土",
	"222221": "土",
	"222211": "土",
	"222111": "土",
	"221111": "土",
	"211111": "土",
	"212111": "土",
	"212222": "土",
	// 兑宫八卦
	"211211": "金",
	"211212": "金",
	"211222": "金",
	"211122": "金",
	"212122": "金",
	"222122": "金",
	"221122": "金",
	"221211": "金",
}

func init() {
}

func (ly *LiuyaoController) Index() {

	out := make(map[string]interface{})

	out["class_id"] = 0
	out["title"] = "六爻排盘"

	ly.Data["data"] = out
	ly.TplName = "liuyao/index.html"
}
func (ly *LiuyaoController) RandIndex() {

	out := make(map[string]interface{})

	out["class_id"] = 0
	out["title"] = "六爻自动起卦"

	ly.Data["data"] = out
	ly.TplName = "liuyao/randindex.html"
}

func (ly *LiuyaoController) Format() {
	title := ly.GetString("title")
	sex, _ := ly.GetInt("sex")
	sexStr := "女"
	if sex == 1 {
		sexStr = "男"
	}
	startTime := ly.GetString("start_time")
	yao1 := ly.GetString("yao1")
	yao2 := ly.GetString("yao2")
	yao3 := ly.GetString("yao3")
	yao4 := ly.GetString("yao4")
	yao5 := ly.GetString("yao5")
	yao6 := ly.GetString("yao6")

	var userGua = yao6 + yao5 + yao4 + yao3 + yao2 + yao1
	var error string
	if utf8.RuneCountInString(userGua) < 6 {
		error = "错误，卦爻错误"
	}

	if startTime == "" {
		now := time.Now()
		year, month, day := now.Date()
		startTime = fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	}
	var resBodyData = make(map[string]interface{})
	//查看库里的农历
	lunarData := models.LunarService.GetRowByDay(db, startTime)

	if lunarData.Day == "" {
		resBodyData = getLunar(startTime)
	} else {
		resBodyData["hyear"] = lunarData.LunarYear
		resBodyData["cyclicalMonth"] = lunarData.LunarMonth
		resBodyData["cyclicalDay"] = lunarData.LunarDay
		resBodyData["suit"] = lunarData.Suit
		resBodyData["taboo"] = lunarData.Taboo
	}

	// 动爻
	userGua2 := strings.Replace(userGua, "3", "1", -1)
	userGua2 = strings.Replace(userGua2, "4", "2", -1)

	userGuaBian := strings.Replace(userGua, "4", "1", -1)
	userGuaBian = strings.Replace(userGuaBian, "3", "2", -1)

	userGuaFormat := Gua64[userGua2]
	userGuaFormat2 := Gua64[userGuaBian]

	// 变卦六亲根据主卦计算
	guaWuXing := GuaWuXinMap[userGua2]

	bianYaoArr := strings.Split(userGuaFormat2.detail[1], "")
	bianYaoWuXing := bianYaoArr[3]
	bianYaoLiuqin := LiuQinMap[guaWuXing+bianYaoWuXing]
	bianyao1 := bianYaoLiuqin + strings.Join(bianYaoArr[2:], "")

	bianYaoArr2 := strings.Split(userGuaFormat2.detail[2], "")
	bianYaoWuXing2 := bianYaoArr2[3]
	bianYaoLiuqin2 := LiuQinMap[guaWuXing+bianYaoWuXing2]
	bianyao2 := bianYaoLiuqin2 + strings.Join(bianYaoArr2[2:], "")

	bianYaoArr3 := strings.Split(userGuaFormat2.detail[3], "")
	bianYaoWuXing3 := bianYaoArr3[3]
	bianYaoLiuqin3 := LiuQinMap[guaWuXing+bianYaoWuXing3]
	bianyao3 := bianYaoLiuqin3 + strings.Join(bianYaoArr3[2:], "")

	bianYaoArr4 := strings.Split(userGuaFormat2.detail[4], "")
	bianYaoWuXing4 := bianYaoArr4[3]
	bianYaoLiuqin4 := LiuQinMap[guaWuXing+bianYaoWuXing4]
	bianyao4 := bianYaoLiuqin4 + strings.Join(bianYaoArr4[2:], "")

	bianYaoArr5 := strings.Split(userGuaFormat2.detail[5], "")
	bianYaoWuXing5 := bianYaoArr5[3]
	bianYaoLiuqin5 := LiuQinMap[guaWuXing+bianYaoWuXing5]
	bianyao5 := bianYaoLiuqin5 + strings.Join(bianYaoArr5[2:], "")

	bianYaoArr6 := strings.Split(userGuaFormat2.detail[6], "")
	bianYaoWuXing6 := bianYaoArr6[3]
	bianYaoLiuqin6 := LiuQinMap[guaWuXing+bianYaoWuXing6]
	bianyao6 := bianYaoLiuqin6 + strings.Join(bianYaoArr6[2:], "")

	guaArr := strings.Split(userGua, "")

	userGuaYaoOrigial := make(map[int]string)

	// 主卦动爻替换
	for i, v := range guaArr {
		index := 6 - i
		if v == "3" {
			yaoLaoYang := userGuaFormat.detail[index]
			yaoLaoYang = strings.Replace(yaoLaoYang, "、", " O->", 1)
			userGuaYaoOrigial[index] = yaoLaoYang
			continue
		}
		if v == "4" {
			yaoLaoYin := userGuaFormat.detail[index]
			yaoLaoYin = strings.Replace(yaoLaoYin, "、、", " X->", 1)
			index := 6 - i
			userGuaYaoOrigial[index] = yaoLaoYin
			continue
		}
		userGuaYaoOrigial[index] = userGuaFormat.detail[index]
	}

	// 装六神
	todayLunar := resBodyData["cyclicalDay"].(string)
	todayLunarArr := strings.Split(todayLunar, "")

	todayTianGan := todayLunarArr[0]
	todayDiZhi := todayLunarArr[1]

	todayTianGanNum := tianganNumMap[todayTianGan]
	todayDiZhiNum := dizhiNumMap[todayDiZhi]
	// 计算空亡
	kongwangNum := 0
	if todayDiZhiNum <= todayTianGanNum {
		kongwangNum = (todayDiZhiNum + 12) - todayTianGanNum
	} else {
		kongwangNum = todayDiZhiNum - todayTianGanNum
	}

	dayLiuShen := LiuShen[todayTianGan]

	LiuShenIndex := LiuShenSort[dayLiuShen]

	var LiuShenMap = map[int]string{}

	LiuShenMap[0] = dayLiuShen

	var i = 0
	for {
		i++
		LiuShenIndexNew := 0
		LiuShenIndexNew = LiuShenIndex + i
		if LiuShenIndexNew <= 6 {
			start := (LiuShenIndexNew - 1) * 6
			end := LiuShenIndexNew * 6
			LiuShenMap[i] = LiuShenStr[start:end]
		} else {
			LiuShenIndexNew = LiuShenIndexNew - 6
			start := (LiuShenIndexNew - 1) * 6
			end := LiuShenIndexNew * 6
			LiuShenMap[i] = LiuShenStr[start:end]
		}
		if len(LiuShenMap) == 6 {
			break
		}
	}

	out := make(map[string]interface{})

	out["class_id"] = 10
	out["kongwang"] = kongwangMap[kongwangNum]

	out["liushen1"] = LiuShenMap[0]
	out["liushen2"] = LiuShenMap[1]
	out["liushen3"] = LiuShenMap[2]
	out["liushen4"] = LiuShenMap[3]
	out["liushen5"] = LiuShenMap[4]
	out["liushen6"] = LiuShenMap[5]
	out["title"] = title
	out["sex"] = sexStr
	out["start_time"] = startTime
	out["lunar"] = resBodyData
	out["shensha"] = "驿马-" + MaXin[todayDiZhi] +
		"|		桃花-" + TaoHua[todayDiZhi] +
		"|		日禄-" + RiLu[todayTianGan] +
		"|		贵人-" + GuiRen[todayTianGan]
	out["gua_name"] = userGuaFormat.name
	out["gua_bian_name"] = userGuaFormat2.name
	out["yao1"] = userGuaYaoOrigial[1]
	out["yao2"] = userGuaYaoOrigial[2]
	out["yao3"] = userGuaYaoOrigial[3]
	out["yao4"] = userGuaYaoOrigial[4]
	out["yao5"] = userGuaYaoOrigial[5]
	out["yao6"] = userGuaYaoOrigial[6]
	out["bian1"] = bianyao1
	out["bian2"] = bianyao2
	out["bian3"] = bianyao3
	out["bian4"] = bianyao4
	out["bian5"] = bianyao5
	out["bian6"] = bianyao6
	out["fushen1"] = userGuaFormat.fushen[1]
	out["fushen2"] = userGuaFormat.fushen[2]
	out["fushen3"] = userGuaFormat.fushen[3]
	out["fushen4"] = userGuaFormat.fushen[4]
	out["fushen5"] = userGuaFormat.fushen[5]
	out["fushen6"] = userGuaFormat.fushen[6]
	out["error"] = error

	ly.Data["data"] = out
	ly.TplName = "liuyao/format.html"
}

func (ly *LiuyaoController) RandFormat() {
	title := ly.GetString("title")
	startTime := ly.GetString("start_time")

	sex, _ := ly.GetInt("sex")
	sexStr := "女"
	if sex == 1 {
		sexStr = "男"
	}

	rand.Seed(time.Now().UnixNano())
	var yingbiStr1, yingbiStr2, yingbiStr3, yingbiStr4, yingbiStr5, yingbiStr6 string
	var yingbi1, yingbi2, yingbi3, yingbi4, yingbi5, yingbi6 string
	for i := 0; i < 3; i++ {
		rand := rand.Intn(100)
		if rand%2 == 0 {
			yingbi1 = fmt.Sprintf("%d", 2)
		} else {
			yingbi1 = fmt.Sprintf("%d", 1)
		}
		yingbiStr1 += fmt.Sprintf("%s", yingbi1)
	}
	for i := 0; i < 3; i++ {
		rand := rand.Intn(100)
		if rand%2 == 0 {
			yingbi2 = fmt.Sprintf("%d", 2)
		} else {
			yingbi2 = fmt.Sprintf("%d", 1)
		}
		yingbiStr2 += fmt.Sprintf("%s", yingbi2)
	}
	for i := 0; i < 3; i++ {
		rand := rand.Intn(100)
		if rand%2 == 0 {
			yingbi3 = fmt.Sprintf("%d", 2)
		} else {
			yingbi3 = fmt.Sprintf("%d", 1)
		}
		yingbiStr3 += fmt.Sprintf("%s", yingbi3)
	}
	for i := 0; i < 3; i++ {
		rand := rand.Intn(100)
		if rand%2 == 0 {
			yingbi4 = fmt.Sprintf("%d", 2)
		} else {
			yingbi4 = fmt.Sprintf("%d", 1)
		}
		yingbiStr4 += fmt.Sprintf("%s", yingbi4)
	}
	for i := 0; i < 3; i++ {
		rand := rand.Intn(100)
		if rand%2 == 0 {
			yingbi5 = fmt.Sprintf("%d", 2)
		} else {
			yingbi5 = fmt.Sprintf("%d", 1)
		}
		yingbiStr5 += fmt.Sprintf("%s", yingbi5)
	}
	for i := 0; i < 3; i++ {
		rand := rand.Intn(100)
		if rand%2 == 0 {
			yingbi6 = fmt.Sprintf("%d", 2)
		} else {
			yingbi6 = fmt.Sprintf("%d", 1)
		}
		yingbiStr6 += fmt.Sprintf("%s", yingbi6)
	}
	var yaoMap = map[string]string{
		"111": "3",
		"222": "4",
		"112": "2",
		"121": "2",
		"211": "2",
		"221": "1",
		"212": "1",
		"122": "1",
	}
	yao1 := yaoMap[yingbiStr1]
	yao2 := yaoMap[yingbiStr2]
	yao3 := yaoMap[yingbiStr3]
	yao4 := yaoMap[yingbiStr4]
	yao5 := yaoMap[yingbiStr5]
	yao6 := yaoMap[yingbiStr6]

	var userGua = yao6 + yao5 + yao4 + yao3 + yao2 + yao1
	var error string
	if utf8.RuneCountInString(userGua) < 6 {
		error = "错误，卦爻错误"
	}

	if startTime == "" {
		now := time.Now()
		year, month, day := now.Date()
		startTime = fmt.Sprintf("%d-%02d-%02d", year, month, day)
	}
	var resBodyData = make(map[string]interface{})
	//查看库里的农历
	lunarData := models.LunarService.GetRowByDay(db, startTime)
	fmt.Println("lunarData", lunarData.Day)
	if lunarData.Day == "" {
		resBodyData = getLunar(startTime)
	} else {
		resBodyData["hyear"] = lunarData.LunarYear
		resBodyData["cyclicalMonth"] = lunarData.LunarMonth
		resBodyData["cyclicalDay"] = lunarData.LunarDay
		resBodyData["suit"] = lunarData.Suit
		resBodyData["taboo"] = lunarData.Taboo
	}

	//动爻
	userGua2 := strings.Replace(userGua, "3", "1", -1)
	userGua2 = strings.Replace(userGua2, "4", "2", -1)

	userGuaBian := strings.Replace(userGua, "4", "1", -1)
	userGuaBian = strings.Replace(userGuaBian, "3", "2", -1)
	fmt.Println(userGua)
	fmt.Println(userGua2)
	userGuaFormat := Gua64[userGua2]
	userGuaFormat2 := Gua64[userGuaBian]

	//变卦六亲根据主卦计算
	guaWuXing := GuaWuXinMap[userGua2]

	bianYaoArr := strings.Split(userGuaFormat2.detail[1], "")
	bianYaoWuXing := bianYaoArr[3]
	bianYaoLiuqin := LiuQinMap[guaWuXing+bianYaoWuXing]
	bianyao1 := bianYaoLiuqin + strings.Join(bianYaoArr[2:], "")

	bianYaoArr2 := strings.Split(userGuaFormat2.detail[2], "")
	bianYaoWuXing2 := bianYaoArr2[3]
	bianYaoLiuqin2 := LiuQinMap[guaWuXing+bianYaoWuXing2]
	bianyao2 := bianYaoLiuqin2 + strings.Join(bianYaoArr2[2:], "")

	bianYaoArr3 := strings.Split(userGuaFormat2.detail[3], "")
	bianYaoWuXing3 := bianYaoArr3[3]
	bianYaoLiuqin3 := LiuQinMap[guaWuXing+bianYaoWuXing3]
	bianyao3 := bianYaoLiuqin3 + strings.Join(bianYaoArr3[2:], "")

	bianYaoArr4 := strings.Split(userGuaFormat2.detail[4], "")
	bianYaoWuXing4 := bianYaoArr4[3]
	bianYaoLiuqin4 := LiuQinMap[guaWuXing+bianYaoWuXing4]
	bianyao4 := bianYaoLiuqin4 + strings.Join(bianYaoArr4[2:], "")

	bianYaoArr5 := strings.Split(userGuaFormat2.detail[5], "")
	bianYaoWuXing5 := bianYaoArr5[3]
	bianYaoLiuqin5 := LiuQinMap[guaWuXing+bianYaoWuXing5]
	bianyao5 := bianYaoLiuqin5 + strings.Join(bianYaoArr5[2:], "")

	bianYaoArr6 := strings.Split(userGuaFormat2.detail[6], "")
	bianYaoWuXing6 := bianYaoArr6[3]
	bianYaoLiuqin6 := LiuQinMap[guaWuXing+bianYaoWuXing6]
	bianyao6 := bianYaoLiuqin6 + strings.Join(bianYaoArr6[2:], "")

	guaArr := strings.Split(userGua, "")

	userGuaYaoOrigial := make(map[int]string)

	//主卦动爻替换
	for i, v := range guaArr {
		index := 6 - i
		if v == "3" {
			yaoLaoYang := userGuaFormat.detail[index]
			yaoLaoYang = strings.Replace(yaoLaoYang, "、", " O->", 1)
			userGuaYaoOrigial[index] = yaoLaoYang
			continue
		}
		if v == "4" {
			yaoLaoYin := userGuaFormat.detail[index]
			yaoLaoYin = strings.Replace(yaoLaoYin, "、、", " X->", 1)
			index := 6 - i
			userGuaYaoOrigial[index] = yaoLaoYin
			continue
		}
		userGuaYaoOrigial[index] = userGuaFormat.detail[index]
	}
	//装六神
	todayLunar := resBodyData["cyclicalDay"].(string)
	todayLunarArr := strings.Split(todayLunar, "")

	todayTianGan := todayLunarArr[0]
	todayDiZhi := todayLunarArr[1]

	todayTianGanNum := tianganNumMap[todayTianGan]
	todayDiZhiNum := dizhiNumMap[todayDiZhi]
	//计算空亡
	kongwangNum := 0
	if todayDiZhiNum <= todayTianGanNum {
		kongwangNum = (todayDiZhiNum + 12) - todayTianGanNum
	} else {
		kongwangNum = todayDiZhiNum - todayTianGanNum
	}

	dayLiuShen := LiuShen[todayTianGan]

	LiuShenIndex := LiuShenSort[dayLiuShen]

	var LiuShenMap = map[int]string{}

	LiuShenMap[0] = dayLiuShen

	var i = 0
	for {
		i++
		LiuShenIndexNew := 0
		LiuShenIndexNew = LiuShenIndex + i
		if LiuShenIndexNew <= 6 {
			start := (LiuShenIndexNew - 1) * 6
			end := LiuShenIndexNew * 6
			LiuShenMap[i] = LiuShenStr[start:end]
		} else {
			LiuShenIndexNew = LiuShenIndexNew - 6
			start := (LiuShenIndexNew - 1) * 6
			end := LiuShenIndexNew * 6
			LiuShenMap[i] = LiuShenStr[start:end]
		}
		if len(LiuShenMap) == 6 {
			break
		}
	}

	out := make(map[string]interface{})

	out["class_id"] = 10
	out["kongwang"] = kongwangMap[kongwangNum]

	out["liushen1"] = LiuShenMap[0]
	out["liushen2"] = LiuShenMap[1]
	out["liushen3"] = LiuShenMap[2]
	out["liushen4"] = LiuShenMap[3]
	out["liushen5"] = LiuShenMap[4]
	out["liushen6"] = LiuShenMap[5]
	out["title"] = title
	out["sex"] = sexStr
	out["shensha"] = "驿马-" + MaXin[todayDiZhi] +
		"|		桃花-" + TaoHua[todayDiZhi] +
		"|		日禄-" + RiLu[todayTianGan] +
		"|		贵人-" + GuiRen[todayTianGan]
	out["start_time"] = startTime
	out["lunar"] = resBodyData
	out["gua_name"] = userGuaFormat.name
	out["gua_bian_name"] = userGuaFormat2.name
	out["yao1"] = userGuaYaoOrigial[1]
	out["yao2"] = userGuaYaoOrigial[2]
	out["yao3"] = userGuaYaoOrigial[3]
	out["yao4"] = userGuaYaoOrigial[4]
	out["yao5"] = userGuaYaoOrigial[5]
	out["yao6"] = userGuaYaoOrigial[6]
	out["bian1"] = bianyao1
	out["bian2"] = bianyao2
	out["bian3"] = bianyao3
	out["bian4"] = bianyao4
	out["bian5"] = bianyao5
	out["bian6"] = bianyao6
	out["fushen1"] = userGuaFormat.fushen[1]
	out["fushen2"] = userGuaFormat.fushen[2]
	out["fushen3"] = userGuaFormat.fushen[3]
	out["fushen4"] = userGuaFormat.fushen[4]
	out["fushen5"] = userGuaFormat.fushen[5]
	out["fushen6"] = userGuaFormat.fushen[6]
	out["error"] = error

	ly.Data["data"] = out
	ly.TplName = "liuyao/randformat.html"
}

func getLunar(startTime string) map[string]interface{} {
	client := &http.Client{}
	request, err := http.NewRequest("GET", LunarApi+startTime, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")

	var resBodyData map[string]interface{}

	if err != nil {
		panic(err)
	}

	response, _ := client.Do(request)
	defer response.Body.Close()

	var resBody map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&resBody); err != nil {
		panic(err)
	}
	status := resBody["status"]
	if status == float64(200) {
		resBodyData = resBody["data"].(map[string]interface{})
		//日历保存
		year := resBodyData["year"].(float64)
		month := resBodyData["month"].(float64)
		day := resBodyData["day"].(float64)

		stringDay := fmt.Sprintf("%1.0f%02.0f%02.0f", year, month, day)
		LunarYear := resBodyData["cyclicalYear"]
		LunarMonth := resBodyData["cyclicalMonth"]
		LunarcDay := resBodyData["cyclicalDay"]
		Suit := resBodyData["suit"]
		Taboo := resBodyData["taboo"]
		Jieqi := ""

		Lunar := models.LunarCalendar{}
		Lunar.Day = stringDay
		Lunar.LunarYear = LunarYear.(string)
		Lunar.LunarMonth = LunarMonth.(string)
		Lunar.LunarDay = LunarcDay.(string)
		Lunar.Suit = Suit.(string)
		Lunar.Taboo = Taboo.(string)
		Lunar.Jieqi = Jieqi
		Lunar.CreatedAt = time.Now()

		if err := models.LunarService.Create(db, &Lunar); err != nil {

		}
	} else {
		panic(resBody)
	}
	return resBodyData
}
