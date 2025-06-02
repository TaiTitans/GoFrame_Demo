package logs

import (
	"github.com/gogf/gf/v2/os/glog"
)

var Logger = glog.New()

func Init() {
	Logger.SetPath("./logs/log_files")
	Logger.SetStdoutPrint(true)
	Logger.SetLevelStr("all")
	Logger.SetFile("app.log")
}
