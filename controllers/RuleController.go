package controllers

import (
	"template/models"
	"github.com/astaxie/beego"
	"template/utils"
	"strconv"
	// "encoding/json"
	// "bytes"
	// "log"
)

type RuleController struct {
	beego.Controller
}

func (c *RuleController) AeTest() {
	page, perr := c.GetInt("page")
	if perr != nil {
		page = 1
	}
	page_size, serr := c.GetInt("page_size")
	if serr != nil {
		page_size = 10
	}

	var stringCond = make(map[string]string)
	var intCond = make(map[string]int)

	name := c.GetString("sname")
	stringCond["name"] = name

	mediaId,merr := c.GetInt("smedia")
	if merr == nil {
		intCond["media"] = mediaId
	}

	rules := models.TestList(page, page_size, stringCond, intCond)

	for index, val := range rules {
		media := val["Media"].(string)
		rules[index]["Media"] = convertMedia(media)
	}

	c.Data["rules"] = rules
	c.Data["ap"] = string("ae_test")
	c.Layout = "components/layout.tpl"

	total := models.Count(stringCond, intCond)
	c.Data["paginator"] = utils.Set(page, page_size, total)
	c.Data["sname"] = name
	c.Data["smedia"] = strconv.Itoa(mediaId)

	mediaMap := map[string]string{
		"1":  "搜狐网",
		"2":  "手机搜狐网",
		"4":  "搜狐新闻客户端",
		"8":  "搜狐视频",
		"11": "移动版视频",
		"25": "搜狐资讯版",
	}

	c.Data["mediaList"] = mediaMap
}

func convertMedia(media string) string {
	mediaMap := map[string]string{
		"1":  "搜狐网",
		"2":  "手机搜狐网",
		"4":  "搜狐新闻客户端",
		"8":  "搜狐视频",
		"11": "移动版视频",
		"25": "搜狐资讯版",
	}
	name, ok := mediaMap[media]
	if ok {
		return name
	} else {
		return ""
	}
}

func (c *RuleController) ExTest() {
	page, perr := c.GetInt("page")
	if perr != nil {
		page = 1
	}
	page_size, serr := c.GetInt("page_size")
	if serr != nil {
		page_size = 10
	}

	var stringCond = make(map[string]string)
	name := c.GetString("sname")
	stringCond["name"] = name

	rules := models.ExTestList(page, page_size, stringCond)

	// for index, val := range rules {
	// 	by, _ := json.MarshalIndent(val["Content"], "", "\t")
	// 	rules[index]["Content"] = string(by)
		
	// 	var out bytes.Buffer
	// 	b, _ := json.Marshal(val["Content"])
	// 	err := json.Indent(&out, b, "", "\t")

		
	// 	log.Println(out, err)

	// }

	c.Data["rules"] = rules
	c.Data["ap"] = string("ex_test")
	c.Layout = "components/layout.tpl"

	total := models.ExCount(stringCond)
	c.Data["paginator"] = utils.Set(page, page_size, total)
	c.Data["sname"] = name
}