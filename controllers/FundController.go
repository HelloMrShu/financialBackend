package controllers

import (
	"funds/models"
	"time"
)

type FundController struct {
	BaseController
}

func (c *FundController) FundList() {
	page,_ := c.GetInt("page", 1)
	pageSize,_ := c.GetInt("page_size", 10)

	sectors,_ := models.SectorList(1, 100)
	
	idToNameMap := make(map[int64]string)
	for _, item := range sectors {
		id := item["Id"].(int64)
		name := item["Name"].(string)
		idToNameMap[id] = name 
	}
	
	funds, total := models.FundList(page, pageSize)

	for key, fund := range funds {
		ts := fund["Created"].(int64)
		funds[key]["Created"] = time.Unix(ts, 0).Format("2006-01-02 15:04:05")
		
		level := fund["Level"].(int64)
		levelDisplay := ""
		for i := 0; i < int(level); i++ {
			levelDisplay += "★";
		}
		funds[key]["Level_display"] = levelDisplay

		sectorId := fund["Sector_id"].(int64)
		funds[key]["Sector_name"] = idToNameMap[sectorId]
	}

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