package controllers

import (
	"funds/models"
	// "time"
)

type FundController struct {
	BaseController
}

func (c *FundController) FundList() {
	page,_ := c.GetInt("page", 1)
	pageSize,_ := c.GetInt("page_size", 10)
	
	funds, total := models.FundList(page, pageSize)
	c.ApiResponse(200, "success", funds, total)
}

func (c *FundController) FundSave() {

	strCond := make(map[string]string)
	intCond := make(map[string]int8)
	level, _ := c.GetInt8("level", 0)

	strCond["name"] = c.GetString("name")
	strCond["intro"] = c.GetString("intro")
	strCond["type"] = c.GetString("type")
	intCond["level"] = level
	
	status := models.FundSave(strCond, intCond)
	c.Response(200, "success", status)
}

func (c *FundController) FundDelete() {

	id, _ := c.GetInt32("id", 0)

	status := false
	if id != 0 {
		status = models.FundDelete(id)
	}
	c.Response(200, "success", status)
}