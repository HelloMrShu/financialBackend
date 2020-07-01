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
	Level 		int
	Type 		string
	Bid_rate 	float64
	Sale_week_rate 	float64
	Sale_month_rate 	float64
	Sale_year_rate 	float64
	Sector_id 	int
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

func FundSave(strCond map[string]string,
	intCond map[string]int,
	floatCond map[string]float64) bool {
	
	timestamp := int32(time.Now().Unix())
	fund := Fund{
		Name: strCond["name"],
		Intro: strCond["intro"],
		Type: strCond["type"],
		Level: intCond["level"],
		Sector_id: intCond["sector_id"],
		Bid_rate: floatCond["bid_rate"],
		Sale_week_rate: floatCond["sale_week_rate"],
		Sale_month_rate: floatCond["sale_month_rate"],
		Sale_year_rate: floatCond["sale_year_rate"],
		Created: timestamp}

	orm := orm.NewOrm()
	_, err := orm.Insert(&fund)
	if err != nil {
		return false
	}
	return true
}

func FundDelete(id int32) bool {
	
	fund := Fund{Id:id}
	orm := orm.NewOrm()
	_, err := orm.Delete(&fund)
	if err != nil {
		return false
	}
	return true
}