package controllers

import (
	"template/models"

	"github.com/astaxie/beego"
	"template/utils"
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

	rules := models.TestList(page, page_size)

	for index, val := range rules {
		media := val["Media"].(string)
		rules[index]["Media"] = convertMedia(media)
	}
	c.Data["rules"] = rules
	c.Data["ap"] = string("ae_test")
	c.Layout = "index.tpl"

	total := models.Count()
	c.Data["paginator"] = utils.Set(page, page_size, total)
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
