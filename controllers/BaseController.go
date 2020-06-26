package controllers

import (
	"github.com/astaxie/beego"
	"math"
)

type BaseController struct {
	beego.Controller
}

//分页Response结构体
type ApiJson struct {
	Msg  string 	    `json:"msg"`
	Code int		    `json:"code"`
	Data interface{}	`json:"data"`		//Data字段需要设置为interface类型以便接收任意数据
	Pagination Paginate `json:"pagination"`
	
}

type Paginate struct {
	Current int 		`json:"current"`
	PageSize int 		`json:"pageSize"`
	Total int 			`json:"total"`
	Pages int 			`json:"pages"` 
}

//普通Response结构体
type NormalJson struct {
	Msg  string 	    `json:"msg"`
	Code int		    `json:"code"`
	Data interface{}	`json:"data"`
}

func (c *BaseController) ApiResponse(code int,msg string,data interface{},total int) {
	var apiJson ApiJson
	apiJson.Msg = msg
	apiJson.Code = code
	apiJson.Data = data

	page,_ := c.GetInt("page", 1)
	pageSize,_ := c.GetInt("page_size", 10)

	var paginate Paginate
	paginate.Current = page
	paginate.PageSize = pageSize
	paginate.Pages = int(math.Ceil(float64(total) / float64(pageSize)))
	paginate.Total = total
	apiJson.Pagination = paginate

	c.Data["json"] = &apiJson	//将结构体数组根据tag解析为json
	c.ServeJSON()			//对json进行序列化输出
	c.StopRun()				//终止执行逻辑
}

func (c *BaseController) Response(code int,msg string,data interface{}) {
	var normalJson NormalJson
	normalJson.Msg = msg
	normalJson.Code = code
	normalJson.Data = data

	c.Data["json"] = &normalJson	//将结构体数组根据tag解析为json
	c.ServeJSON()			//对json进行序列化输出
	c.StopRun()				//终止执行逻辑
}

