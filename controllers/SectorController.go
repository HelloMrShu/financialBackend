package controllers

import (
	"funds/models"
)

type SectorController struct {
	BaseController
}

func (c *SectorController) SectorList() {
	page,_ := c.GetInt("page", 1)
	pageSize,_ := c.GetInt("page_size", 10)
	sectors, totals := models.SectorList(page, pageSize)
	c.ApiResponse(200, "success", sectors, totals)
}

func (c *SectorController) SectorSave() {
	name := c.GetString("name")
	intro := c.GetString("intro")

	result := models.SectorSave(name, intro)	
	c.Response(200, "success", result)
}

func (c *SectorController) SectorDelete() {
	id,_ := c.GetInt32("id")

	result := models.SectorDelete(id)
	c.Response(200, "success", result)
}