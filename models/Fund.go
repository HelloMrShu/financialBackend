package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Fund struct {
	Id   		int32 `orm:"pk"`
	Name 		string
	Intro 		string
	Level 		int32
	Type 		string
	Bid_rate 	float64
	Sale_week_rate 	float64
	Sale_month_rate 	float64
	Sale_year_rate 	float64
	Sector_id 	int32
	Updated 	int32
	Created 	int32
}

func init() {
	orm.RegisterModel(new(Fund))
}

func FundList(page int, page_size int) ([]orm.Params, int) {
	
	offset := (page - 1) * page_size
	o := orm.NewOrm()
	qs := o.QueryTable("fund")

	var funds []orm.Params
	qs.OrderBy("-id").Limit(page_size, offset).Values(&funds)

	total,_ := qs.Count()
	return funds, int(total)
}

func FundSave(name string, intro string) bool {
	
	timestamp := int32(time.Now().Unix())
	sector := Sector{Name:name, Intro:intro, Created:timestamp}

	orm := orm.NewOrm()
	_, err := orm.Insert(&sector)
	if err != nil {
		return false
	}
	return true
}

func FundDelete(id int32) bool {
	
	sector := Sector{Id:id}
	orm := orm.NewOrm()
	_, err := orm.Delete(&sector)
	if err != nil {
		return false
	}
	return true
}