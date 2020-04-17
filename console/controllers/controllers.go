package controllers

import (
	"net/http"

	"smsgate/console/controllers/app"
	"smsgate/console/controllers/blacklist"
	"smsgate/console/controllers/sms"
	"smsgate/console/controllers/send"
	"smsgate/console/controllers/stat"
	"smsgate/console/controllers/tpl"
	"smsgate/console/controllers/user"
	"smsgate/console/controllers/tools"
)

//Controllers 进行路由配置
var Controllers = map[string]http.HandlerFunc{
	"/":                 stat.Index,
	"/sms/index":        sms.Index,
	"/sms/reply":        sms.Reply,
	"/app/index":        app.Index,
	"/app/add":          app.Add,
	"/app/edit":         app.Edit,
	"/app/update":       app.Update,
	"/app/delete":       app.Delete,
	"/tpl/index":        tpl.Index,
	"/tpl/add":          tpl.Add,
	"/tpl/edit":         tpl.Edit,
	"/tpl/update":       tpl.Update,
	"/tpl/delete":       tpl.Delete,
	"/blacklist/index":  blacklist.Index,
	"/blacklist/add":    blacklist.Add,
	"/blacklist/edit":   blacklist.Edit,
	"/blacklist/delete": blacklist.Delete,
	"/user/username":    user.Username,
	"/send/index":       send.Index,
	"/send/send":        send.Send,
	// 统计
	"/stat/index":       stat.Index,
	"/stat/daily":       stat.Daily,
	"/stat/channel":     stat.Channel,
	"/stat/cost":        stat.Cost,
	"/stat/app":         stat.App,
	"/stat/status":      stat.Status,
	"/stat/number":      stat.Number,
	"/stat/tpl":         stat.Tpl,
	"/stat/mobile":      stat.Mobile,
	// 工具
	"/tools/status":     tools.Status,
	"/tools/queue":      tools.Queue,
}
