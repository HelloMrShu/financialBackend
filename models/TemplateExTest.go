package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TemplateExTest struct {
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
	orm.RegisterModel(new(TemplateExTest))
}

func (m *TemplateExTest) Read() error {
	if err := orm.NewOrm().Read(m); err != nil {
		return err
	}
	return nil
}

func ExTestList(
	page int,
	page_size int,
	strCond map[string]string) []orm.Params {
	
	offset := (page - 1) * page_size
	o := orm.NewOrm()
	o.Using("extest")
	qs := o.QueryTable("template_ex_test")

	if strCond["name"] != "" {
		qs = qs.Filter("name__contains", strCond["name"])
	}

	var rules []orm.Params
	qs.Limit(page_size, offset).Values(&rules)
	return rules
}

func ExCount(strCond map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable("template_ex_test")
	if strCond["name"] != "" {
		qs = qs.Filter("name__contains", strCond["name"])
	}
	c,_ := qs.Count()
	return int(c)
}