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