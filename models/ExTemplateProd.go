package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ExTemplateProd struct {
	Id   		int32 `orm:"pk"`
	Platform 	string
	Bidmode 	string
	Name 		string
	Media  		string
	Special_type string
	Content 	string
	Is_delete	int8
}

func init() {
	orm.RegisterModel(new(ExTemplateProd))
}

func ExProdList(
	page int,
	page_size int,
	strCond map[string]string) []orm.Params {
	
	offset := (page - 1) * page_size
	o := orm.NewOrm()
	o.Using("exchange")
	qs := o.QueryTable("ex_template_prod")

	if strCond["name"] != "" {
		qs = qs.Filter("name__contains", strCond["name"])
	}
	if strCond["special_type"] != "" {
		qs = qs.Filter("special_type__contains", strCond["special_type"])
	}
	var rules []orm.Params
	qs.Limit(page_size, offset).Values(&rules)
	return rules
}

func ExProdCount(strCond map[string]string) int {
	o := orm.NewOrm()
	o.Using("exchange")
	qs := o.QueryTable("ex_template_prod")
	if strCond["name"] != "" {
		qs = qs.Filter("name__contains", strCond["name"])
	}
	if strCond["special_type"] != "" {
		qs = qs.Filter("special_type__contains", strCond["special_type"])
	}
	c,_ := qs.Count()
	return int(c)
}