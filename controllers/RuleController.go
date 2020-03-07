package controllers

import (
	"template/models"

	"github.com/astaxie/beego"
	"template/utils"
	"log"
	"strconv"
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

	log.Printf("type is %T", c.Data["smedia"])
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

func findMediaId(name string) int {
	mediaMap := map[string]int{
		"搜狐网" : 1,
		"手机搜狐网" : 2,
		"搜狐新闻客户端" : 4,
		"搜狐视频" : 8,
		"移动版视频" : 11,
		"搜狐资讯版" : 25,
	}

	id, ok := mediaMap[name]
	if ok {
		return id
	} else {
		return 0
	}
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
