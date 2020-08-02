package controllers

import (
	"funds/models"
)

type NewsController struct {
	BaseController
}

func (c *NewsController) NewsList() {
	page,_ := c.GetInt("page", 1)
	pageSize,_ := c.GetInt("page_size", 10)
	id,_ := c.GetInt("id", 0)

	news, totals := models.NewsList(page, pageSize, id)
	for key, new := range(news) {
		ts := new["Created"].(int64)
		news[key]["Created"] = ts
	}
	c.ApiResponse(200, "success", news, totals)
}

func (c *NewsController) NewsSave() {
	title := c.GetString("title")
	content := c.GetString("content")
	id,_ := c.GetInt32("id", 0)

	result := models.NewsSave(id, title, content)	
	c.Response(200, "success", result)
}

func (c *NewsController) NewsDelete() {
	id,_ := c.GetInt32("id")

	result := models.NewsDelete(id)
	c.Response(200, "success", result)
}