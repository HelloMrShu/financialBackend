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
	intCond := make(map[string]int)
	floatCond := make(map[string]float64)

	level, _ := c.GetInt("level", 0)
	bid_rate, _ := c.GetFloat("bid_rate")
	sale_week_rate, _ := c.GetFloat("bid_rate")
	sale_month_rate, _ := c.GetFloat("bid_rate")
	sale_year_rate, _ := c.GetFloat("bid_rate")
	sector_id, _ := c.GetInt("sector_id")

	strCond["name"] = c.GetString("name")
	strCond["intro"] = c.GetString("intro")
	strCond["type"] = c.GetString("type")

	intCond["level"] = level
	intCond["sector_id"] = sector_id

	floatCond["bid_rate"] = bid_rate
	floatCond["sale_week_rate"] = sale_week_rate
	floatCond["sale_month_rate"] = sale_month_rate
	floatCond["sale_year_rate"] = sale_year_rate
	
	status := models.FundSave(strCond, intCond, floatCond)
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