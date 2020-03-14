package components

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/config"
	"encoding/json"
	"fmt"
)

func InitLogger()  {
	BConfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil{
		fmt.Println("config init error:", err)
	}
	maxlines, lerr := BConfig.Int64("log::maxlines")
	if lerr != nil {
		maxlines = 1000
	}
	log := logs.NewLogger(maxlines)

	logConf := make(map[string]interface{})
	logConf["filename"] = BConfig.String("log::log_path")
	level,_ := BConfig.Int("log::log_level")
	logConf["level"] = level
	 
    confStr, err := json.Marshal(logConf)
    if err != nil {
        fmt.Println("marshal failed,err:", err)
        return
    }
	log.SetLogger(logs.AdapterFile, string(confStr))
	log.EnableFuncCallDepth(true)
	log.Async()
	log.Debug("logger init success")
}