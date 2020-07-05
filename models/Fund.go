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

func FundSave(id int32, 
	strCond map[string]string,
	intCond map[string]int,
	floatCond map[string]float64) int64 {
	
	timestamp := int32(time.Now().Unix())

	o := orm.NewOrm()
	fund := Fund{Id: id}
	rerr := o.Read(&fund)

	var err interface{} = nil
	ret := int64(0)
	
	if rerr == nil {
		fund.Id = id
		fund.Name = strCond["name"]
		fund.Intro = strCond["intro"]
		fund.Type = strCond["type"]
		fund.Level = intCond["level"]
		fund.Sector_id = intCond["sector_id"]
		fund.Bid_rate = floatCond["bid_rate"]
		fund.Sale_week_rate = floatCond["sale_week_rate"]
		fund.Sale_month_rate = floatCond["sale_month_rate"]
		fund.Sale_year_rate = floatCond["sale_year_rate"]
		fund.Updated = timestamp
		ret, err = o.Update(&fund)

	} else {
		fund.Name = strCond["name"]
		fund.Intro = strCond["intro"]
		fund.Type = strCond["type"]
		fund.Level = intCond["level"]
		fund.Sector_id = intCond["sector_id"]
		fund.Bid_rate = floatCond["bid_rate"]
		fund.Sale_week_rate = floatCond["sale_week_rate"]
		fund.Sale_month_rate = floatCond["sale_month_rate"]
		fund.Sale_year_rate = floatCond["sale_year_rate"]
		fund.Created = timestamp
		ret, err = o.Insert(&fund)
	}

	if err != nil {
		return 0;
	}
	return ret
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