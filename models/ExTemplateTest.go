package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ExTemplateTest struct {
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
	orm.RegisterModel(new(ExTemplateTest))
}

func ExTestList(
	page int,
	page_size int,
	strCond map[string]string) []orm.Params {
	
	offset := (page - 1) * page_size
	o := orm.NewOrm()
	o.Using("test1")
	qs := o.QueryTable("ex_template_test")

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

func ExTestCount(strCond map[string]string) int {
	o := orm.NewOrm()
	o.Using("test1")
	qs := o.QueryTable("ex_template_test")
	if strCond["name"] != "" {
		qs = qs.Filter("name__contains", strCond["name"])
	}
	if strCond["special_type"] != "" {
		qs = qs.Filter("special_type__contains", strCond["special_type"])
	}
	c,_ := qs.Count()
	return int(c)
}